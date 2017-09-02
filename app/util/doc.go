package util

import (
	"path/filepath"
	"strings"

	"github.com/go-aah/website/app/markdown"
	"github.com/go-aah/website/app/models"

	"aahframework.org/aah.v0"
	"aahframework.org/essentials.v0"
	"aahframework.org/log.v0"
)

// GetBranchName method returns the confirmed branch name
func GetBranchName(version string) string {
	releases, _ := aah.AppConfig().StringList("docs.releases")
	if version == releases[0] {
		return "master"
	}
	return version
}

// GetDocDirPath method returns the documentation dir path for given
// language and version.
func GetDocDirPath(version string) string {
	docsDir := filepath.Join(aah.AppConfig().StringDefault("docs.dir", ""), "aah-documentation")
	return filepath.Join(docsDir, version)
}

// RefreshDocContent method clone's the GitHub branch doc version wise into
// local and if already exits it takes a update from GitHub.
// It clears cache too.
func RefreshDocContent(pushEvent *models.GithubPushEvent) {
	releases, _ := aah.AppConfig().StringList("docs.releases")
	GitRefresh(releases)

	version := pushEvent.BranchName()
	if version == "master" {
		version = releases[0]
	}

	log.Infof("BranchName: %s", version)
	docDirPath := GetDocDirPath(version)
	for _, commit := range pushEvent.Commits {
		log.Infof("CommitID: %s, Message: %s", commit.ID, commit.Message)
		log.Infof("Modified: %s, Removed: %s", commit.Modified, commit.Removed)

		for _, f := range commit.Modified {
			if strings.HasSuffix(f, "LICENSE") || strings.HasSuffix(f, "README.md") {
				continue
			}
			mdPath := FilePath(f, docDirPath)
			markdown.RefreshCacheByFile(mdPath)
		}

		for _, f := range commit.Removed {
			mdPath := FilePath(f, docDirPath)
			markdown.RemoveCacheByFile(mdPath)
		}
	}
}

// GitRefresh method clone's the GitHub branch doc version wise into
// local and if already exits it takes a update from GitHub.
func GitRefresh(releases []string) {
	for _, version := range releases {
		docDirPath := GetDocDirPath(version)
		branchName := GetBranchName(version)
		err := GitCloneAndPull(docDirPath, branchName)
		if err != nil {
			log.Error(err)
		}
	}
}

// ContentBasePath method returns the Markdown files base path.
func ContentBasePath() string {
	return filepath.Join(aah.AppBaseDir(), "content")
}

// FilePath method returns markdown file path from given path.
// it bacially remove any extension and adds ".md"
func FilePath(reqPath, prefix string) string {
	reqPath = strings.ToLower(TrimPrefixSlash(reqPath))
	reqPath = ess.StripExt(reqPath) + ".md"
	return filepath.Clean(filepath.Join(prefix, reqPath))
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
