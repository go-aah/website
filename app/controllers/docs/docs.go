package docs

import "github.com/go-aah/website/app/controllers"

// Docs struct documentation application controller
type Docs struct {
	controllers.App
}

// Index method is documentation application home page
func (d *Docs) Index() {
	d.AddViewArg("IsDocumentation", true)
}
