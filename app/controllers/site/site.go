package controllers

import (
	"aahframework.org/aah.v0"
	"aahframework.org/essentials.v0"
	"aahframework.org/log.v0"

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

// Content method display the content based on request path.
func (s *Site) Content() {
	mdPath := markdown.FilePath(s.Req.Path, markdown.ContentBasePath())
	data := aah.Data{
		"Markdown":  string(markdown.Get(mdPath)),
		"CodeBlock": true,
	}

	switch ess.StripExt(s.Req.Path)[1:] {
	// Display the instruction for how to contribute to code.
	case "contribute-to-code":
		s.AddViewArg("IsContributeCode", true)
		s.Reply().HTMLlf("master.html", "contribute-code.html", data)

	// Display aah framework instructions to report
	// security issues privately and the disclosing to public.
	case "security":
		s.Reply().HTMLlf("master.html", "security.html", data)

	// Display aah framework features
	case "features":
		s.AddViewArg("IsFeatures", true)
		s.Reply().HTMLlf("master.html", "features.html", data)

	// Display why aah framework description
	case "why-aah":
		s.AddViewArg("IsWhyAah", true)
		s.Reply().HTMLlf("master.html", "why-aah.html", data)

	// Display aah framework website credit info.
	case "credits":
		s.Reply().HTMLlf("master.html", "credits.html", data)
	default:
		s.NotFound(false)
	}
}

// NotFound method for unavailable pages on the site.
func (s *Site) NotFound(isStatic bool) {
	log.Warnf("Page not found: %s", s.Req.Path)
	s.Reply().HTMLlf("master.html", "notfound.html", aah.Data{})
}
