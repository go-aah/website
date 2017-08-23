package controllers

import (
	"net/http"

	"aahframework.org/aah.v0-unstable"
	"aahframework.org/ahttp.v0"
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
		"CodeBlock":     true,
		"IsGetInvolved": true,
	})
}

// Content method display the content based on request path.
func (s *SiteController) Content() {
	mdPath := util.FilePath(s.Req.Path, util.ContentBasePath())
	data := aah.Data{"CodeBlock": true}

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
func (s *SiteController) Team() {
	s.Reply().HTML(aah.Data{
		"CodeBlock": true,
		"IsTeam":    true,
	})
}

// Privacy method display aahframework.org websit privacy information.
func (s *SiteController) Privacy() {
	s.Reply().Ok()
}

func onPreReplyEvent(e *aah.Event) {
	ctx := e.Data.(*aah.Context)
	if ctx.IsStaticRoute() {
		ctx.Res.Header().Set(ahttp.HeaderAccessControlAllowOrigin, "*")
	}
}

func init() {
	aah.OnPreReply(onPreReplyEvent)
}
