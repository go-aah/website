// aah framework v0.2 - https://aahframework.org
// FILE: aah.go
// GENERATED CODE - DO NOT EDIT

package main

import (
	"flag"
	"fmt"
	"reflect"

	"aahframework.org/aah"
	"aahframework.org/config"
	"aahframework.org/essentials"
	"aahframework.org/log"
	docs "github.com/go-aah/website/app/controllers/docs"
	root "github.com/go-aah/website/app/controllers/root"
)

var (
	AppBinaryName = "aahframework"
	AppVersion = "5fd15af"
	AppBuildDate = "2017-02-23T16:21:23-08:00"
	_ = reflect.Invalid
)

func main() {
	// Defining flags
	version := flag.Bool("version", false, "Display application version and build date.")
	configPath := flag.String("config", "", "Absolute path of external config file.")
	flag.Parse()

	// display application information
	if *version {
		fmt.Printf("%-12s: %s\n", "Binary Name", AppBinaryName)
		fmt.Printf("%-12s: %s\n", "Version", AppVersion)
		fmt.Printf("%-12s: %s\n", "Build Date", AppBuildDate)
		return
	}

	aah.Init("github.com/go-aah/website")

	// Loading externally supplied config file
	if !ess.IsStrEmpty(*configPath) {
		externalConfig, err := config.LoadFile(*configPath)
		if err != nil {
			log.Fatalf("Unable to load external config: %s", *configPath)
		}

		aah.MergeAppConfig(externalConfig)
	}

	// Adding all the controllers which refers 'aah.Controller' directly
	// or indirectly from app/controllers/** 
	aah.AddController((*docs.DocsApp)(nil),
	  []*aah.MethodInfo{
	    &aah.MethodInfo{
	      Name: "Index",
	      Parameters: []*aah.ParameterInfo{ 
	      },
	    },
	    
	  })
	
	aah.AddController((*root.App)(nil),
	  []*aah.MethodInfo{
	    &aah.MethodInfo{
	      Name: "Index",
	      Parameters: []*aah.ParameterInfo{ 
	      },
	    },
	    &aah.MethodInfo{
	      Name: "Overview",
	      Parameters: []*aah.ParameterInfo{ 
	      },
	    },
	    &aah.MethodInfo{
	      Name: "Credits",
	      Parameters: []*aah.ParameterInfo{ 
	      },
	    },
	    
	  })
	

  aah.Start()
}
