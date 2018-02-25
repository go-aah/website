// aah application initialization - configuration, server extensions, middleware's, etc.
// Customize it per your application needs.

package main

import (
	"html/template"

	"aahframework.org/aah.v0"

	// Registering HTML minifier for web application
	_ "github.com/aah-cb/minify"

	"github.com/go-aah/website/app/controllers"
	"github.com/go-aah/website/app/markdown"
	"github.com/go-aah/website/app/util"
)

func init() {

	//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
	// Server Extensions
	// Doc: https://docs.aahframework.org/server-extension.html
	//
	// Recommended: Define a function with meaningful name on a package and
	// register it here. Extensions function name gets logged in the log,
	// its very helpful to have meaningful log information.
	//
	// Such as:
	//    - Dedicated package for config loading
	//    - Dedicated package for datasource connections
	//    - etc
	//__________________________________________________________________________

	// Event: OnInit
	// Published right after the `aah.AppConfig()` is loaded.
	//
	// aah.OnInit(config.LoadRemote)

	// Event: OnStart
	// Published right before the start of aah go Server.
	//
	// aah.OnStart(db.Connect)
	// aah.OnStart(cache.Load)
	aah.OnStart(controllers.LoadValuesFromConfig)
	aah.OnStart(markdown.FetchMarkdownConfig)
	aah.OnStart(util.PullGithubDocsAndLoadCache)

	// Event: OnShutdown
	// Published on receiving OS Signals `SIGINT` or `SIGTERM`.
	//
	// aah.OnShutdown(cache.Flush)
	// aah.OnShutdown(db.Disconnect)
	aah.OnShutdown(markdown.ClearCache)

	// Event: OnPreReply
	aah.OnPreReply(util.AllowAllOriginForStaticFiles)

	//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
	// Middleware's
	// Doc: https://docs.aahframework.org/middleware.html
	//
	// Executed in the order they are defined. It is recommended; NOT to change
	// the order of pre-defined aah framework middleware's.
	//__________________________________________________________________________
	aah.Middlewares(
		aah.RouteMiddleware,
		aah.CORSMiddleware,
		aah.BindMiddleware,
		aah.AntiCSRFMiddleware,
		aah.AuthcAuthzMiddleware,

		//
		// NOTE: Register your Custom middleware's right here
		//

		aah.ActionMiddleware,
	)

	//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
	// Register template methods
	//__________________________________________________________________________
	aah.AddTemplateFunc(template.FuncMap{
		"docurlc":    util.TmplDocURLc,
		"docediturl": util.TmplDocEditURL,
		"absrequrl":  util.TmplAbsReqURL,
		"vergteq":    util.TmplVerGtEq,
		"dverdis":    util.TmplDVerDis,
	})

}
