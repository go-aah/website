package controllers

import (
	"aahframe.work"
)

// AppController struct application controller
type AppController struct {
	*aah.Context
}

// Before method is called for all the application requests
// before the controllers action gets called.
func (c *AppController) Before() {
	releases, _ := aah.App().Config().StringList("docs.releases")
	c.AddViewArg("Releases", releases)
}

// HealthCheck method is used ping..pong health check
func (c *AppController) HealthCheck() {
	c.Reply().Text("pong")
}
