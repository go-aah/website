package controllers

import (
	"aahframework.org/aah"
	"aahframework.org/aah/render"
	"github.com/go-aah/website/app/markdown"
)

// App struct application controller
type App struct {
	*aah.Controller
}

// Index method is application home page
func (a *App) Index() {
}

// Overview method display aah framework overview's.
func (a *App) Overview() {
	data := render.Data{
		"Markdown":   string(markdown.Get(a.Req.Path)),
		"CodeBlock":  true,
		"IsOverview": true,
	}
	a.Reply().HTML(data)
}

// Credits method display aah framework wesite credit info.
func (a *App) Credits() {
	data := render.Data{
		"Markdown":  string(markdown.Get(a.Req.Path)),
		"CodeBlock": true,
	}
	a.Reply().HTML(data)
}
