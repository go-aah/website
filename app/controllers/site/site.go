package controllers

import (
	"aahframework.org/aah/render"

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

// Overview method display aah framework overview's.
func (s *Site) Overview() {
	data := render.Data{
		"Markdown":   string(markdown.Get(s.Req.Path)),
		"CodeBlock":  true,
		"IsOverview": true,
	}
	s.Reply().HTML(data)
}

// GetInvolved method display aah framework community and contribution info.
func (s *Site) GetInvolved() {
	s.AddViewArg("IsGetInvoled", true)
}

// ContributeCode method display the instruction for how to contribute to code.
func (s *Site) ContributeCode() {
	data := render.Data{
		"Markdown":         string(markdown.Get(s.Req.Path)),
		"IsContributeCode": true,
	}
	s.Reply().HTML(data)
}

// Security method display aah framework instructions to report
// security issues privately and the disclosing to public.
func (s *Site) Security() {
	data := render.Data{
		"Markdown":   string(markdown.Get(s.Req.Path)),
		"IsSecurity": true,
	}
	s.Reply().HTML(data)
}

// Credits method display aah framework wesite credit info.
func (s *Site) Credits() {
	data := render.Data{
		"Markdown":  string(markdown.Get(s.Req.Path)),
		"CodeBlock": true,
	}
	s.Reply().HTML(data)
}
