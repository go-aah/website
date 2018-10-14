package controllers

import (
	"net/http"

	"aahframe.work"

	"aahframework.org/website/app/markdown"
	"aahframework.org/website/app/util"
)

// SiteController struct root domain controller
type SiteController struct {
	AppController
}

// Index method is application home page
func (c *SiteController) Index() {
	c.Reply().Ok()
}

// GetInvolved method display aah framework community and contribution info.
func (c *SiteController) GetInvolved() {
	c.Reply().HTML(aah.Data{
		"CodeBlock":     true,
		"IsGetInvolved": true,
	})
}

// Content method display the content based on request path.
func (c *SiteController) Content() {
	mdPath := util.FilePath(c.Req.Path, util.ContentBasePath())
	data := aah.Data{"CodeBlock": true}

	if article, found := markdown.Get(mdPath); found {
		data["Article"] = article
	} else {
		c.Log().Warnf("Page not found: %s", c.Req.Path)
		c.Reply().Error(&aah.Error{
			Code:    http.StatusNotFound,
			Message: "Not Found",
		})
		return
	}

	var viewFile string
	switch util.CreateKey(c.Req.Path) {
	case "contribute-to-code":
		c.AddViewArg("IsContributeCode", true)
		viewFile = "contribute-code.html"
	case "security-vulnerabilities":
		viewFile = "security-vulnerabilities.html"
	}

	c.Reply().HTMLf(viewFile, data)
}

// Team method display aah framework team info.
func (c *SiteController) Team() {
	c.Reply().HTML(aah.Data{
		"IsTeam": true,
	})
}

// Privacy method display aahframework.org websit privacy information.
func (c *SiteController) Privacy() {
	c.Reply().Ok()
}

// WhyAah method to show info about aah framework.
func (c *SiteController) WhyAah() {
	c.Reply().HTML(aah.Data{
		"IsWhyAah": true,
	})
}

// Support method to display support aah page.
func (c *SiteController) Support() {
	c.Reply().HTML(aah.Data{
		"IsSupport": true,
	})
}

// Features method to show aah features.
func (c *SiteController) Features() {
	c.Reply().HTML(aah.Data{
		"IsFeatures": true,
	})
}

// Security method to show aah security info.
func (c *SiteController) Security() {
	c.Reply().HTML(aah.Data{
		"IsSecurity": true,
	})
}
