package util

import (
	"path/filepath"

	"github.com/go-aah/website/app/markdown"

	"aahframework.org/aah.v0"
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
	docsDir := aah.AppConfig().StringDefault("docs.dir", "")
	return filepath.Join(docsDir, version)
}

// RefreshDocContent method clone's the GitHub branch doc version wise into
// local and if already exits it takes a update from GitHub.
// It clears cache too.
func RefreshDocContent() {
	releases, _ := aah.AppConfig().StringList("docs.releases")
	for _, version := range releases {
		docDirPath := GetDocDirPath(version)
		branchName := GetBranchName(version)
		err := GitCloneAndPull(docDirPath, branchName)
		if err != nil {
			log.Error(err)
		}
	}

	// Clear markdown cache
	markdown.ClearCache()
}
