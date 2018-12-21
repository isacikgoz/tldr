package main

import (
	"fmt"

	"github.com/isacikgoz/tldr/pkg/config"
	"github.com/isacikgoz/tldr/pkg/pages"
	"github.com/isacikgoz/tldr/pkg/prompt"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	clear       = kingpin.Flag("clear-cache", "clear local repository then clone github.com/tldr-pages/tldr").Short('c').Bool()
	update      = kingpin.Flag("update", "pulls the latest commits from github.com/tldr-pages/tldr").Short('u').Bool()
	interactive = kingpin.Flag("interactive", "interactive mode.").Short('i').Default("true").Bool()

	page = kingpin.Arg("command", "Name of the command.").String()
)

func main() {
	// start := time.Now()
	kingpin.Version("tldr++ version 0.0.1 (pre-release)")
	// parse the command line flag and options
	kingpin.Parse()
	config.StartUp()
	if *clear {
		err := config.Clear()
		if err != nil {

		}
		return
	}
	if *update {
		err := config.PullSource()
		if err != nil {

		}
		return
	}
	if len(*page) == 0 {
		kingpin.Usage()
		return
	}
	p, err := pages.Read(*page)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}

	prompter := prompt.New(p)

	if err = prompter.RenderPage(); err != nil {
		fmt.Printf("%s\n", err.Error())
	}

	t, err := prompter.Selection()
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}

	cmd, err := prompter.GenerateCommand(t)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}

	if err = prompter.Run(cmd); err != nil {
		fmt.Printf("%s\n", err.Error())
	}

}
