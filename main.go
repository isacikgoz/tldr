package main

import (
	"github.com/isacikgoz/tldr/pkg/config"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	clone  = kingpin.Flag("clone", "clones the repository from github.com/tldr-pages/tldr").Short('c').Bool()
	update = kingpin.Flag("update", "pulls the latest commits from github.com/tldr-pages/tldr").Short('u').Bool()
)

func main() {
	kingpin.Version("tldr++ version 0.0.1 (pre-release)")
	// parse the command line flag and options
	kingpin.Parse()
	config.StartUp()
	if *clone {
		err := config.CloneSource()
		if err != nil {

		}
	}
	if *update {
		err := config.PullSource()
		if err != nil {

		}
	}

}
