package util

import (
	"encoding/json"
	"fmt"
	"html/template"
	"path"
	"path/filepath"
	"strings"

	"github.com/go-aah/website/app/markdown"
	"github.com/go-aah/website/app/models"

	"aahframework.org/aah.v0"
	"aahframework.org/ahttp.v0"
	"aahframework.org/essentials.v0"
	"aahframework.org/log.v0"
)

var (
	releases      []string
	editURLPrefix string
)

// BranchName method returns the confirmed branch name
func BranchName(version string) string {
	releases, _ := aah.AppConfig().StringList("docs.releases")
	if version == releases[0] {
		return "master"
	}
	return version
}

// DocBaseDir method returns the aah documentation based directory.
func DocBaseDir() string {
	return filepath.Join(aah.AppConfig().StringDefault("docs.dir", ""), "aah-documentation")
}

// DocVersionBaseDir method returns the documentation dir path for given
// language and version.
func DocVersionBaseDir(version string) string {
	return filepath.Join(DocBaseDir(), version)
}

// RefreshDocContent method clone's the GitHub branch doc version wise into
// local and if already exits it takes a update from GitHub.
// It clears cache too.
func RefreshDocContent(pushEvent *models.GithubPushEvent) {
	version := pushEvent.BranchName()
	if version == "master" {
		version = releases[0]
	}

	if !ess.IsSliceContainsString(releases, version) {
		log.Warnf("Branch Name [%s] not found", version)
		return
	}

	GitRefresh(releases...)

	log.Infof("BranchName: %s", version)
	docVersionBaseDir := DocVersionBaseDir(version)
	for _, commit := range pushEvent.Commits {
		log.Infof("CommitID: %s, Message: %s", commit.ID, commit.Message)
		log.Infof("Modified: %s, Removed: %s", commit.Modified, commit.Removed)

		for _, f := range commit.Modified {
			if strings.HasSuffix(f, "LICENSE") || strings.HasSuffix(f, "README.md") {
				continue
			}
			mdPath := FilePath(f, docVersionBaseDir)
			markdown.RefreshCacheByFile(mdPath)
		}

		for _, f := range commit.Removed {
			mdPath := FilePath(f, docVersionBaseDir)
			markdown.RemoveCacheByFile(mdPath)
		}
	}
}

// GitRefresh method clone's the GitHub branch doc version wise into
// local and if already exits it takes a update from GitHub.
func GitRefresh(releases ...string) {
	for _, version := range releases {
		docDirPath := DocVersionBaseDir(version)
		branchName := BranchName(version)
		log.Infof("Git refresh: %s => %s", version, docDirPath)
		if err := GitCloneAndPull(docDirPath, branchName); err != nil {
			log.Error(err)
		}
	}
}

// ContentBasePath method returns the Markdown files base path.
func ContentBasePath() string {
	return filepath.Join(aah.AppVirtualBaseDir(), "content")
}

// FilePath method returns markdown file path from given path.
// it bacially remove any extension and adds ".md"
func FilePath(reqPath, prefix string) string {
	reqPath = strings.ToLower(TrimPrefixSlash(reqPath))
	reqPath = ess.StripExt(reqPath) + ".md"
	return path.Clean(path.Join(prefix, reqPath))
}

// TrimPrefixSlash method trims the prefix slash from the given path
func TrimPrefixSlash(str string) string {
	return strings.TrimPrefix(str, "/")
}

// CreateKey method creates markdown file name from request path.
func CreateKey(rpath string) string {
	key := ess.StripExt(TrimPrefixSlash(rpath))
	return strings.Replace(key, "/", "-", -1)
}

// PullGithubDocsAndLoadCache method pulls github docs and populate documentation
// in the cache
func PullGithubDocsAndLoadCache(e *aah.Event) {
	cfg := aah.AppConfig()
	editURLPrefix = cfg.StringDefault("docs.edit_url_prefix", "")
	releases, _ = cfg.StringList("docs.releases")

	docBasePath := DocBaseDir()

	if aah.AppProfile() == "prod" {
		ess.DeleteFiles(docBasePath)
	}

	_ = ess.MkDirAll(docBasePath, 0755)
	GitRefresh(releases[0])
	go GitRefresh(releases[1:]...)

	if cfg.BoolDefault("markdown.cache", false) {
		go markdown.LoadCache(filepath.Join(docBasePath, releases[0]))
		go markdown.LoadCache(ContentBasePath())
	}
}

// ReadJSON method read the JSON file for given path and unmarshals into given object.
func ReadJSON(ctx *aah.Context, jsonPath string, v interface{}) {
	f, err := aah.AppVFS().Open(jsonPath)
	if err != nil {
		ctx.Log().Errorf("%s: %v", jsonPath, err)
	}
	defer ess.CloseQuietly(f)

	if err = json.NewDecoder(f).Decode(v); err != nil {
		ctx.Log().Errorf("%s: %v", jsonPath, err)
	}
}

// TmplDocURLc method compose documentation navi URL based on version
func TmplDocURLc(viewArgs map[string]interface{}, key string) template.HTML {
	req := viewArgs[aah.KeyViewArgRequest].(*ahttp.Request)
	version := req.PathValue("version")
	if !ess.IsSliceContainsString(releases, version) {
		version = releases[0]
	}

	fileName := aah.AppConfig().StringDefault(key, "")
	return template.HTML(fmt.Sprintf("/%s/%s", version, fileName))
}

// TmplRDocURL method returns the doc relative url with give prefix.
func TmplRDocURL(rootPrefix template.URL, key string) template.URL {
	return template.URL(fmt.Sprintf("%s%s/%s", string(rootPrefix), releases[0], aah.AppConfig().StringDefault(key, "")))
}

// TmplDocEditURL method compose github documentation edit URL
func TmplDocEditURL(docFile string) template.URL {
	var pattern string
	if strings.HasPrefix(docFile, "/") {
		pattern = "%s%s"
	} else {
		pattern = "%s/%s"
	}
	return template.URL(fmt.Sprintf(pattern, editURLPrefix, docFile))
}
