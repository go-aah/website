package controllers

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"path"
	"path/filepath"
	"strings"

	"aahframework.org/aah.v0-unstable"
	"aahframework.org/ahttp.v0"
	"aahframework.org/essentials.v0"
	"aahframework.org/log.v0"

	"github.com/go-aah/website/app/markdown"
	"github.com/go-aah/website/app/models"
	"github.com/go-aah/website/app/util"
)

var (
	releases      []string
	docBasePath   string
	editURLPrefix string
)

// DocController struct documentation domain controller
type DocController struct {
	AppController
}

// Before method doc brfore interceptor
func (d *DocController) Before() {
	// Execute nested interceptor
	d.AppController.Before()

	d.AddViewArg("ShowVersionDocs", true).
		AddViewArg("ShowInsightSideNav", true).
		AddViewArg("CodeBlock", true).
		AddViewArg("ShowVersionNo", true)
}

// Index method is documentation application home page
func (d *DocController) Index() {
	d.Reply().Redirect(d.ReverseURL("docs.version_home", releases[0]))
}

// VersionHome method Displays the documentation in selected language. Default
// is English.
func (d *DocController) VersionHome(version string) {
	if !ess.IsSliceContainsString(releases, version) {
		switch ess.StripExt(version) {
		case "favicon":
			d.Reply().ContentType("image/x-icon").File(filepath.Join("img", version))
		case "robots":
			d.Reply().ContentType(ahttp.ContentTypePlainText.String()).File("docs_" + version)
		case "sitemap":
			d.Reply().ContentType(ahttp.ContentTypeXML.String()).File("docs_" + version)
		case "godoc":
			d.GoDoc()
		case "tutorials":
			d.Tutorials()
		default:
			queryStr := d.Req.Params.Query.Encode()
			targetURL := d.ReverseURL("docs.show_doc", releases[0], version)
			if !ess.IsStrEmpty(queryStr) {
				targetURL = targetURL + "?" + queryStr
			}
			d.Reply().Redirect(targetURL)
		}
		return
	}

	data := aah.Data{
		"IsVersionHome":      true,
		"ShowVersionDocs":    false,
		"ShowInsightSideNav": false,
		"CurrentVersion":     version,
	}
	d.Reply().HTMLl("docs.html", data)
}

// ShowDoc method displays requested documentation page based language and version.
func (d *DocController) ShowDoc(version, content string) {
	isTutorial := false
	if version == "tutorial" {
		if content == "/i18n.html" {
			d.Reply().Redirect("/tutorial/i18n-url-query-param.html")
			return
		}

		version = releases[0] // take the latest version
		isTutorial = true
		d.ViewArgs()["ShowVersionNo"] = false
	}

	d.AddViewArg("CurrentVersion", version)
	branchName := util.GetBranchName(version)
	if branchName == "master" {
		d.AddViewArg("LatestRelease", true)
	}

	switch ess.StripExt(util.TrimPrefixSlash(content)) {
	case "release-notes":
		d.ReleaseNotes(version)
		return
	}

	// if it's add prefix
	if isTutorial {
		content = filepath.Join("tutorial", content)
	}

	docPath := path.Clean(path.Join(version, content))
	mdPath := util.FilePath(docPath, docBasePath)
	article, found := markdown.Get(mdPath)
	if !found {
		d.NotFound()
		return
	}
	data := aah.Data{"Article": article, "DocFile": ess.StripExt(content) + ".md"}
	d.Reply().HTMLl("docs.html", data)
}

// GoDoc method display aah framework godoc links
func (d *DocController) GoDoc() {
	data := aah.Data{
		"IsGoDoc": true,
	}
	d.Reply().HTMLlf("docs.html", "godoc.html", data)
}

// Tutorials method display aah framework tutorials github links or guide.
func (d *DocController) Tutorials() {
	d.Reply().HTMLlf("docs.html", "tutorials.html", aah.Data{
		"ShowVersionNo": false,
	})
}

// ReleaseNotes method display aah framework release notes, changelogs, migration notes.
func (d *DocController) ReleaseNotes(version string) {
	changelogMdPath := util.FilePath(path.Join(version, "changelog.md"), docBasePath)
	whatsNewMdPath := util.FilePath(path.Join(version, "whats-new.md"), docBasePath)
	migrationGuideMdPath := util.FilePath(path.Join(version, "migration-guide.md"), docBasePath)

	changelog, _ := markdown.Get(changelogMdPath)
	whatsNew, _ := markdown.Get(whatsNewMdPath)
	migrationGuide, _ := markdown.Get(migrationGuideMdPath)

	data := aah.Data{
		"IsReleaseNotes": true,
		"Changelog":      changelog,
		"WhatsNew":       whatsNew,
		"MigrationGuide": migrationGuide,
	}
	d.Reply().HTMLlf("docs.html", "release-notes.html", data)
}

// BeforeRefreshDoc method is interceptor.
func (d *DocController) BeforeRefreshDoc() {
	githubEvent := strings.TrimSpace(d.Req.Header.Get("X-Github-Event"))
	githubDeliveryID := strings.TrimSpace(d.Req.Header.Get("X-Github-Delivery"))
	if githubEvent != "push" || ess.IsStrEmpty(githubDeliveryID) {
		log.Warnf("Github event: %s, DeliveryID: %s", githubEvent, githubDeliveryID)
		d.Reply().BadRequest().JSON(aah.Data{"message": "bad request"})
		d.Abort()
		return
	}

	hubSignature := strings.TrimSpace(d.Req.Header.Get("X-Hub-Signature"))
	log.Infof("Github Signature: %s", hubSignature)

	body, err := ioutil.ReadAll(d.Req.Unwrap().Body)
	if err != nil {
		log.Errorf("Body read error: %s", hubSignature)
		d.Reply().BadRequest().JSON(aah.Data{"message": "bad request"})
		d.Abort()
		return
	}
	d.Req.Unwrap().Body = ioutil.NopCloser(bytes.NewReader(body))

	if ess.IsStrEmpty(hubSignature) || !util.IsValidHubSignature(hubSignature, body) {
		log.Warnf("Github Invalied Signature: %s", hubSignature)
		d.Reply().BadRequest().JSON(aah.Data{"message": "bad request"})
		d.Abort()
		return
	}
	log.Infof("Event: %s, DeliveryID: %s", githubEvent, githubDeliveryID)
}

// RefreshDoc method to refresh documentation from github
func (d *DocController) RefreshDoc(pushEvent *models.GithubPushEvent) {
	go util.RefreshDocContent(pushEvent)

	log.Info("Github event received, docs are being refereshed")
	d.Reply().JSON(aah.Data{
		"message": "Docs are being refreshed",
	})
}

// NotFound method handles not found URLs.
func (d *DocController) NotFound() {
	log.Warnf("Page not found: %s", d.Req.Path)
	d.Reply().HTMLlf("docs.html", "notfound.html", aah.Data{
		"IsNotFound": true,
	})
}

func docsContentRefresh(e *aah.Event) {
	cfg := aah.AppConfig()
	editURLPrefix = cfg.StringDefault("docs.edit_url_prefix", "")
	releases, _ = cfg.StringList("docs.releases")
	docBasePath = filepath.Join(cfg.StringDefault("docs.dir", ""), "aah-documentation")

	_ = ess.MkDirAll(docBasePath, 0755)
	util.GitRefresh(releases)

	if cfg.BoolDefault("markdown.cache", false) {
		go markdown.LoadCache(filepath.Join(docBasePath, releases[0]))
		go markdown.LoadCache(util.ContentBasePath())
	}
}

func init() {
	aah.AddTemplateFunc(template.FuncMap{
		"docurlc": func(viewArgs map[string]interface{}, key string) template.HTML {
			params := viewArgs[aah.KeyViewArgRequestParams].(*ahttp.Params)
			version := params.PathValue("version")
			if !ess.IsSliceContainsString(releases, version) {
				version = releases[0]
			}

			return template.HTML(fmt.Sprintf("/%s/%s",
				version,
				aah.AppConfig().StringDefault(key, "")))
		},
		"absrequrl": func(viewArgs map[string]interface{}) template.URL {
			return template.URL(fmt.Sprintf("%s://%s%s", viewArgs["Scheme"], viewArgs["Host"], viewArgs["RequestPath"]))
		},
		"docediturl": func(docFile string) template.URL {
			var pattern string
			if strings.HasPrefix(docFile, "/") {
				pattern = "%s%s"
			} else {
				pattern = "%s/%s"
			}
			return template.URL(fmt.Sprintf(pattern, editURLPrefix, docFile))
		},
	})

	aah.OnStart(docsContentRefresh)
}
