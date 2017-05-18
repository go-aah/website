package models

import "path/filepath"

type (
	// GithubPushEvent github push event payload structure.
	GithubPushEvent struct {
		Ref        string     `json:"ref"`
		Repository Repository `json:"repository"`
		Commits    []Commit   `json:"commits"`
	}

	// Commit file list.
	Commit struct {
		ID       string   `json:"id"`
		Message  string   `json:"message"`
		Removed  []string `json:"removed"`
		Modified []string `json:"modified"`
	}

	// Repository repo info.
	Repository struct {
		Name     string `json:"name"`
		FullName string `json:"full_name"`
	}
)

// BranchName method returns the branch name of the push event.
func (pe *GithubPushEvent) BranchName() string {
	return filepath.Base(pe.Ref)
}
