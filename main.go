package main

import (
	"fmt"
	"time"

	"github.com/isacikgoz/tldr/pkg/config"
	"github.com/isacikgoz/tldr/pkg/pages"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	clone  = kingpin.Flag("clone", "clones the repository from github.com/tldr-pages/tldr").Short('c').Bool()
	update = kingpin.Flag("update", "pulls the latest commits from github.com/tldr-pages/tldr").Short('u').Bool()

	page = kingpin.Arg("command", "Name of the command.").Required().String()
)

func main() {
	start := time.Now()
	kingpin.Version("tldr++ version 0.0.1 (pre-release)")
	// parse the command line flag and options
	kingpin.Parse()
	if *clone {
		config.StartUp()
		err := config.CloneSource()
		if err != nil {

		}
	}
	if *update {
		config.StartUp()
		err := config.PullSource()
		if err != nil {

		}
	}
	s, err := pages.Read(*page)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}
	fmt.Printf("%s\n", s.String())
	elapsed := time.Since(start)
	fmt.Printf("Query finished in: %s\n", elapsed)
}
