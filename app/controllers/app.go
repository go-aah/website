package controllers

import "aahframework.org/aah"

// App struct application controller
type App struct {
	*aah.Controller
}

// Before method is interceptor method for all requests
// before controllers gets called.
func (a App) Before() {
	a.AddViewArg("AppVersion", aah.AppBuildInfo().Version)
}
