package main

import (
	"os"

	"github.com/jessevdk/go-flags"
)

// Define the command line options
type GorgonOpts struct {
	// Specify a config file
	Config string `short:"c" long:"config" description:"Config file" default:"~/.config/gorgon/gorgon.json"`
	// Persist cmdline options to the config file
	Persist bool `short:"P" long:"persist" description:"Persist cmdline options to the config file"`
	// Specify the project name
	Project string `short:"p" long:"project" description:"Project name" required:"true" env:"GORGON_PROJ_NAME"`
	// Specify the project owner
	Owner string `short:"o" long:"owner" description:"Project owner" required:"true" env:"GORGON_PROJ_OWNER"`
	// Specify the projct id
	ID int `short:"i" long:"id" description:"Project ID" required:"true" env:"GORGON_PROJ_ID"`
	// Specify the username if different than the proejct owner
	Username string `short:"u" long:"username" description:"Username" env:"GORGON_USER"`
	// Specify the number of results to accept, default 4096
	Results int `short:"r" long:"results" description:"Number of results to accept" default:"4096"`
	// Specify upstream sync, default is down
	Upstream bool `short:"U" long:"upstream" description:"Upstream sync"`
	// Enable verbose mode
	Verbose bool `short:"v" long:"verbose" description:"Show verbose debug information"`
	// Enable debug mode
	Debug bool `short:"d" long:"debug" description:"Show debug information"`
	// Disable output
	Quiet bool `short:"q" long:"quiet" description:"Disable output"`
}

func parseArgs() *GorgonOpts {
	opts := GorgonOpts{}

	// Parse the command line options
	_, err := flags.Parse(&opts)

	// Check for errors, go-flags will print the error message
	if err != nil {
		os.Exit(1)
	}

	return &opts
}
