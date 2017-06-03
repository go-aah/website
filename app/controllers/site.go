package controllers

import (
	"aahframework.org/aah.v0"
	"aahframework.org/log.v0"

	"github.com/go-aah/website/app/markdown"
	"github.com/go-aah/website/app/util"
)

// Site struct application controller
type Site struct {
	App
}

// Index method is application home page
func (s *Site) Index() {
	s.Reply().Ok()
}

// GetInvolved method display aah framework community and contribution info.
func (s *Site) GetInvolved() {
	data := aah.Data{"CodeBlock": true, "IsGetInvolved": true}
	s.Reply().HTML(data)
}

// Content method display the content based on request path.
func (s *Site) Content() {
	mdPath := util.FilePath(s.Req.Path, util.ContentBasePath())
	data := aah.Data{"CodeBlock": true}

	if article, found := markdown.Get(mdPath); found {
		data["Article"] = article
	} else {
		s.NotFound()
		return
	}

	switch util.CreateKey(s.Req.Path) {
	// Display the instruction for how to contribute to code.
	case "contribute-to-code":
		s.AddViewArg("IsContributeCode", true)
		s.Reply().HTMLf("contribute-code.html", data)
		return

		// Display aah framework instructions to report
		// security issues privately and the disclosing to public.
	case "security-vulnerabilities":
		s.Reply().HTMLf("security-vulnerabilities.html", data)
		return

	// Display aah framework features
	case "features":
		s.AddViewArg("IsFeatures", true)
		s.Reply().HTMLf("features.html", data)
		return

	// Display why aah framework description
	case "why-aah":
		s.AddViewArg("IsWhyAah", true)
	}

	s.Reply().HTML(data)
}

// Team method display aah framework team info.
func (s *Site) Team() {
	data := aah.Data{"CodeBlock": true, "IsTeam": true}
	s.Reply().HTML(data)
}

// NotFound method for unavailable pages on the site.
func (s *Site) NotFound() {
	log.Warnf("Page not found: %s", s.Req.Path)
	s.Reply().HTMLl("notfound.html", aah.Data{})
}
