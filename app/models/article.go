package models

import "aahframework.org/essentials.v0"

// Article hold parsed data of one Markdown file.
type Article struct {
	Title    string
	Desc     string
	Keywords string
	Content  string
	File     string
}

// IsContent method returns the true if content is available otherwise false.
func (a *Article) IsContent() bool {
	return !ess.IsStrEmpty(a.Content)
}
