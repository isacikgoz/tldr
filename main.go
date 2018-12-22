package main

import (
	"fmt"

	"github.com/isacikgoz/tldr/pkg/config"
	"github.com/isacikgoz/tldr/pkg/pages"
	"github.com/isacikgoz/tldr/pkg/prompt"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	clear  = kingpin.Flag("clear-cache", "clear local repository then clone github.com/tldr-pages/tldr").Short('c').Bool()
	update = kingpin.Flag("update", "pulls the latest commits from github.com/tldr-pages/tldr").Short('u').Bool()
	static = kingpin.Flag("static", "static mode.").Short('s').Default("false").Bool()

	page = kingpin.Arg("command", "Name of the command.").Strings()
)

func main() {

	kingpin.Version("tldr++ version 0.1.1 (pre-release)")
	kingpin.Parse()

	config.StartUp(*clear, *update)

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
	if err = prompter.RenderPage(*static); err != nil {
		fmt.Printf("%s", err.Error())
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
