package controllers

import (
	"aahframework.org/aah.v0-unstable"

	"github.com/go-aah/website/app/controllers"
	"github.com/go-aah/website/app/markdown"
)

// Site struct application controller
type Site struct {
	controllers.App
}

// Index method is application home page
func (s *Site) Index() {
}

// GetInvolved method display aah framework community and contribution info.
func (s *Site) GetInvolved() {
	s.AddViewArg("IsGetInvoled", true)
}

// ContributeCode method display the instruction for how to contribute to code.
func (s *Site) ContributeCode() {
	mdPath := markdown.FilePath(s.Req.Path, markdown.ContentBasePath())
	data := aah.Data{
		"Markdown":         string(markdown.Get(mdPath)),
		"IsContributeCode": true,
		"CodeBlock":        true,
	}
	s.Reply().HTML(data)
}

// Security method display aah framework instructions to report
// security issues privately and the disclosing to public.
func (s *Site) Security() {
	mdPath := markdown.FilePath(s.Req.Path, markdown.ContentBasePath())
	data := aah.Data{
		"Markdown":   string(markdown.Get(mdPath)),
		"IsSecurity": true,
	}
	s.Reply().HTML(data)
}

// Credits method display aah framework wesite credit info.
func (s *Site) Credits() {
	mdPath := markdown.FilePath(s.Req.Path, markdown.ContentBasePath())
	data := aah.Data{
		"Markdown":  string(markdown.Get(mdPath)),
		"CodeBlock": true,
	}
	s.Reply().HTML(data)
}
