// aah application initialization - configuration, server extensions, middleware's, etc.
// Customize it per your application needs.

package main

import (
	"html/template"

	"aahframe.work"

	// Registering HTML minifier for web application
	_ "aahframe.work/minify/html"

	"aahframework.org/website/app/controllers"
	"aahframework.org/website/app/markdown"
	"aahframework.org/website/app/util"
)

func init() {
	app := aah.App()

	//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
	// Server Extensions
	// Doc: https://docs.aahframework.org/server-extension.html
	//__________________________________________________________________________
	app.OnStart(controllers.LoadValuesFromConfig)
	app.OnStart(markdown.FetchMarkdownConfig)
	app.OnStart(util.PullGithubDocsAndLoadCache)
	app.OnStart(SubscribeHTTPEvents)

	app.OnPostShutdown(markdown.ClearCache)

	//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
	// Middleware's
	// Doc: https://docs.aahframework.org/middleware.html
	//
	// Executed in the order they are defined. It is recommended; NOT to change
	// the order of pre-defined aah framework middleware's.
	//__________________________________________________________________________
	app.HTTPEngine().Middlewares(
		aah.RouteMiddleware,
		aah.CORSMiddleware,
		aah.BindMiddleware,
		aah.AntiCSRFMiddleware,
		aah.AuthcAuthzMiddleware,
		aah.ActionMiddleware,
	)

	//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
	// Add Custom Template Functions
	// Doc: https://docs.aahframework.org/template-funcs.html
	//__________________________________________________________________________
	app.AddTemplateFunc(template.FuncMap{
		"docurlc":    util.TmplDocURLc,
		"docurl":     util.TmplRDocURL,
		"docediturl": util.TmplDocEditURL,
		"absrequrl":  util.TmplAbsReqURL,
		"vergteq":    util.TmplVerGtEq,
		"dverdis":    util.TmplDVerDis,
	})
}

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// HTTP Events
//
// Subscribing HTTP events on app start.
//__________________________________________________________________________

func SubscribeHTTPEvents(_ *aah.Event) {
	he := aah.App().HTTPEngine()
	he.OnPreReply(util.AllowAllOriginForStaticFiles)
}
