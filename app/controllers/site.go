package controllers

import (
	"net/http"

	"aahframework.org/aah.v0"
	"aahframework.org/log.v0"

	"github.com/go-aah/website/app/markdown"
	"github.com/go-aah/website/app/util"
)

// SiteController struct root domain controller
type SiteController struct {
	AppController
}

// Index method is application home page
func (s *SiteController) Index() {
	s.Reply().Ok()
}

// GetInvolved method display aah framework community and contribution info.
func (s *SiteController) GetInvolved() {
	s.Reply().HTML(aah.Data{
		"IsGetInvolved": true,
	})
}

// Content method display the content based on request path.
func (s *SiteController) Content() {
	mdPath := util.FilePath(s.Req.Path, util.ContentBasePath())
	data := aah.Data{}

	if article, found := markdown.Get(mdPath); found {
		data["Article"] = article
	} else {
		log.Warnf("Page not found: %s", s.Req.Path)
		s.Reply().Error(&aah.Error{
			Code:    http.StatusNotFound,
			Message: "Not Found",
		})
		return
	}

	var viewFile string
	switch util.CreateKey(s.Req.Path) {
	case "contribute-to-code":
		s.AddViewArg("IsContributeCode", true)
		viewFile = "contribute-code.html"
	case "security-vulnerabilities":
		viewFile = "security-vulnerabilities.html"
	}

	s.Reply().HTMLf(viewFile, data)
}

// Team method display aah framework team info.
func (s *SiteController) Team() {
	s.Reply().HTML(aah.Data{
		"IsTeam": true,
	})
}

// Privacy method display aahframework.org websit privacy information.
func (s *SiteController) Privacy() {
	s.Reply().Ok()
}

// WhyAah method to show info about aah framework.
func (s *SiteController) WhyAah() {
	s.Reply().HTML(aah.Data{
		"IsWhyAah": true,
	})
}

// Support method to display support aah page.
func (s *SiteController) Support() {
	s.Reply().HTML(aah.Data{
		"IsSupport": true,
	})
}

// Features method to show aah features.
func (s *SiteController) Features() {
	s.Reply().HTML(aah.Data{
		"IsFeatures": true,
	})
}

// Security method to show aah security info.
func (s *SiteController) Security() {
	s.Reply().HTML(aah.Data{
		"IsSecurity": true,
	})
}
