package controllers

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"path"
	"path/filepath"
	"strings"

	"aahframe.work"
	"aahframe.work/ahttp"
	"aahframe.work/essentials"

	"aahframework.org/website/app/markdown"
	"aahframework.org/website/app/models"
	"aahframework.org/website/app/util"
)

var (
	releases    []string
	docBasePath string
)

// DocController struct documentation domain controller
type DocController struct {
	AppController
}

// Before method doc brfore interceptor
func (c *DocController) Before() {
	// Execute nested interceptor
	c.AppController.Before()

	c.AddViewArg("ShowVersionDocs", true).
		AddViewArg("ShowInsightSideNav", true).
		AddViewArg("ShowVersionNo", true)
}

// Index method is documentation application home page
func (c *DocController) Index() {
	c.Reply().Redirect(c.RouteURL("docs.version_home", releases[0]))
}

// VersionHome method Displays the documentation in selected language. Default
// is English.
func (c *DocController) VersionHome(version string) {
	if !ess.IsSliceContainsString(releases, version) {
		switch ess.StripExt(version) {
		case "favicon":
			c.Reply().ContentType("image/x-icon").
				File(filepath.Join("static", "img", version))
		case "robots":
			c.Reply().ContentType(ahttp.ContentTypePlainText.String()).
				File(filepath.Join("static", "docs_"+version))
		case "sitemap":
			c.Reply().ContentType(ahttp.ContentTypeXML.String()).
				File(filepath.Join("static", "docs_"+version))
		case "browserconfig":
			c.Reply().ContentType(ahttp.ContentTypeXML.String()).
				File(filepath.Join("static", version))
		case "site", "manifest":
			c.Reply().ContentType(ahttp.ContentTypeJSON.String()).
				File(filepath.Join("static", version))
		case "godoc":
			c.GoDoc()
		default:
			queryStr := c.Req.URL().Query().Encode()
			targetURL := c.RouteURL("docs.show_doc", releases[0], version)
			if !ess.IsStrEmpty(queryStr) {
				targetURL = targetURL + "?" + queryStr
			}
			c.Reply().Redirect(targetURL)
		}
		return
	}

	data := aah.Data{
		"IsVersionHome":      true,
		"IsGettingStarted":   false,
		"ShowVersionDocs":    false,
		"ShowInsightSideNav": false,
		"CurrentDocVersion":  version,
	}
	c.addDocVersionComparison(version)
	c.Reply().HTMLl("docs.html", data)
}

// ShowDoc method displays requested documentation page based language and version.
func (c *DocController) ShowDoc(version, content string) {
	// handle certian doc path and updates
	switch content {
	case "error-handling.html":
		if util.VersionLtEq(version, "v0.9") {
			c.Reply().Redirect(c.RouteURL("docs.show_doc", version, "/centralized-error-handler.html"))
			return
		}
	case "centralized-error-handler.html":
		if util.VersionGtEq(version, "v0.10") {
			c.Reply().RedirectWithStatus(
				c.RouteURL("docs.show_doc", version, "/error-handling.html"),
				http.StatusMovedPermanently,
			)
			return
		}
	}

	var pathSeg string
	if !util.IsVersionNo(version) {
		pathSeg = version
		version = releases[0]
	}

	c.AddViewArg("CurrentDocVersion", version)
	c.addDocVersionComparison(version)
	branchName := util.BranchName(version)
	if branchName == releases[0] {
		c.AddViewArg("LatestRelease", true)
	}

	docPath := path.Clean(path.Join(version, pathSeg, content))
	mdPath := util.FilePath(docPath, docBasePath)
	article, found := markdown.Get(mdPath)
	if !found {
		if util.VersionLtEq(version, "v0.10") {
			if strings.Contains(content, "auth-schemes") {
				c.Reply().Redirect(c.RouteURL("docs.show_doc", version, "authentication.html"))
			} else if strings.Contains(content, "cryptography") {
				c.Reply().Redirect(c.RouteURL("docs.version_home", "cryptography.html"))
			}
			return
		}

		c.NotFound()
		return
	}

	data := aah.Data{
		"Article":          article,
		"DocFile":          ess.StripExt(content) + ".md",
		"IsShowDoc":        true,
		"IsGettingStarted": strings.HasSuffix(content, "getting-started.html"),
	}

	c.Reply().HTMLl("docs.html", data)
}

// GoDoc method display aah framework godoc links
func (c *DocController) GoDoc() {
	jsonPath := filepath.ToSlash(path.Join(util.ContentBasePath(), "godoc.json"))
	var godoc []*struct {
		Name       string `json:"name"`
		ImportPath string `json:"importPath"`
		Codecov    string `json:"codecov"`
	}
	util.ReadJSON(c.Context, jsonPath, &godoc)

	c.addDocVersionComparison(releases[0])
	c.Reply().HTMLlf("docs.html", "godoc.html", aah.Data{
		"IsGodoc":           true,
		"ShowVersionNo":     false,
		"CurrentDocVersion": releases[0],
		"Godoc":             godoc,
	})
}

// Examples method display aah framework examples github links or guide.
func (c *DocController) Examples() {
	jsonPath := filepath.ToSlash(path.Join(util.ContentBasePath(), "examples.json"))
	var groups []*struct {
		GroupHeading string `json:"groupHeading"`
		Examples     []*struct {
			DisplayName string `json:"displayName"`
			Name        string `json:"name"`
		} `json:"examples"`
	}
	util.ReadJSON(c.Context, jsonPath, &groups)

	c.addDocVersionComparison(releases[0])
	c.Reply().HTMLlf("docs.html", "examples.html", aah.Data{
		"IsExamples":        true,
		"ShowVersionNo":     false,
		"CurrentDocVersion": releases[0],
		"Examples":          groups,
	})
}

// ReleaseNotes method display aah framework release notes, changelogs, migration notes.
func (c *DocController) ReleaseNotes(version string) {
	changelogMdPath := util.FilePath(path.Join(version, "changelog.md"), docBasePath)
	whatsNewMdPath := util.FilePath(path.Join(version, "whats-new.md"), docBasePath)
	migrationGuideMdPath := util.FilePath(path.Join(version, "migration-guide.md"), docBasePath)

	changelog, _ := markdown.Get(changelogMdPath)
	whatsNew, _ := markdown.Get(whatsNewMdPath)
	migrationGuide, _ := markdown.Get(migrationGuideMdPath)

	c.addDocVersionComparison(version)
	data := aah.Data{
		"IsReleaseNotes":    true,
		"CurrentDocVersion": version,
		"Changelog":         changelog,
		"WhatsNew":          whatsNew,
		"MigrationGuide":    migrationGuide,
	}
	c.Reply().HTMLlf("docs.html", "release-notes.html", data)
}

// BeforeRefreshDoc method is interceptor.
func (c *DocController) BeforeRefreshDoc() {
	if !aah.App().IsEnvProfile("prod") {
		return
	}
	githubEvent := strings.TrimSpace(c.Req.Header.Get("X-Github-Event"))
	githubDeliveryID := strings.TrimSpace(c.Req.Header.Get("X-Github-Delivery"))
	if githubEvent != "push" || ess.IsStrEmpty(githubDeliveryID) {
		c.Log().Warnf("Github event: %s, DeliveryID: %s", githubEvent, githubDeliveryID)
		c.Reply().BadRequest().JSON(aah.Data{"message": "bad request"})
		c.Abort()
		return
	}

	hubSignature := strings.TrimSpace(c.Req.Header.Get("X-Hub-Signature"))
	c.Log().Infof("Github Signature: %s", hubSignature)

	body, err := ioutil.ReadAll(c.Req.Unwrap().Body)
	if err != nil {
		c.Log().Errorf("Body read error: %s", hubSignature)
		c.Reply().BadRequest().JSON(aah.Data{"message": "bad request"})
		c.Abort()
		return
	}

	if ess.IsStrEmpty(hubSignature) || !util.IsValidHubSignature(hubSignature, body) {
		c.Log().Warnf("Github Invalied Signature: %s", hubSignature)
		c.Reply().BadRequest().JSON(aah.Data{"message": "bad request"})
		c.Abort()
		return
	}

	c.Req.Unwrap().Body = ioutil.NopCloser(bytes.NewReader(body))

	c.Log().Infof("Event: %s, DeliveryID: %s", githubEvent, githubDeliveryID)
}

// RefreshDoc method to refresh documentation from github
func (c *DocController) RefreshDoc(pushEvent *models.GithubPushEvent) {
	go util.RefreshDocContent(pushEvent)

	c.Log().Info("Github event received, docs are being refereshed")
	c.Reply().JSON(aah.Data{
		"message": "Docs are being refreshed",
	})
}

// NotFound method handles not found URLs.
func (c *DocController) NotFound() {
	c.Log().Warnf("Page not found: %s", c.Req.Path)
	c.Reply().HTMLlf("docs.html", "notfound.html", aah.Data{
		"IsNotFound": true,
	})
}

func (c *DocController) addDocVersionComparison(curVer string) {
	cv := util.VerRep.Replace(curVer)
	for _, ver := range releases {
		keyPart := util.VerKeyRep.Replace(ver)
		c.AddViewArg("Is"+keyPart+"AndGr", util.VersionGtEq(cv, ver))
	}
}

// LoadValuesFromConfig method loads required value from configuration and others
func LoadValuesFromConfig(e *aah.Event) {
	app := aah.App()
	releases, _ = app.Config().StringList("docs.releases")
	docBasePath = "/aah/documentation"
	docPhysicalBasePath := util.DocBaseDir()
	_ = ess.MkDirAll(docPhysicalBasePath, 0755)

	if err := app.VFS().AddMount(docBasePath, docPhysicalBasePath); err != nil {
		app.Log().Fatal("vfs:", err)
	}
}
