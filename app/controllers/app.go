package controllers

import "aahframework.org/aah.v0"

// App struct application controller
type App struct {
	*aah.Context
}

// Before method is called for all the application requests
// before the controllers action gets called.
func (a App) Before() {
	releases, _ := aah.AppConfig().StringList("docs.releases")
	a.AddViewArg("Releases", releases)
}
