package main

import (
	"fmt"

	"github.com/isacikgoz/tldr/pkg/config"
	"github.com/isacikgoz/tldr/pkg/pages"
	"github.com/isacikgoz/tldr/pkg/prompt"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	clear  = kingpin.Flag("clear-cache", "Clear local repository then clone github.com/tldr-pages/tldr").Short('c').Bool()
	update = kingpin.Flag("update", "Pulls the latest commits from github.com/tldr-pages/tldr").Short('u').Bool()
	static = kingpin.Flag("static", "Static mode, application behaves like a conventional tldr client.").Short('s').Default("false").Bool()
	random = kingpin.Flag("random", "Random page for testing purposes.").Short('r').Default("false").Bool()

	page = kingpin.Arg("command", "Name of the command. (e.g. tldr grep)").Strings()
)

func main() {

	kingpin.Version("tldr++ version 0.5.0")
	kingpin.Parse()

	config.StartUp(*clear, *update)

	var p *pages.Page
	var err error
	if *random {
		p, err = pages.QueryRandom()
	} else if len(*page) == 0 && !*static {
		p, err = pages.ReadAll()
	} else if len(*page) == 0 && *static {
		config.PrintLogo()
		kingpin.Usage()
		return
	} else {
		p, err = pages.Read(*page)
	}
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}

	prompter := prompt.New(p)
	if err = prompter.RenderPage(*static); err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}

	t, err := prompter.Selection()
	if err != nil {
		fmt.Printf("%s", err.Error())
		return
	}

	if t == nil {
		return
	}

	cmd, err := prompter.GenerateCommand(t)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}

	if err = prompter.Run(cmd); err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}
}
