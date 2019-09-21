package main

import (
	"fmt"
	"os"

	"github.com/isacikgoz/gitin/prompt"
	"github.com/isacikgoz/tldr/cli"
	env "github.com/kelseyhightower/envconfig"
)

func main() {
	var o prompt.Options
	err := env.Process("gitin", &o)
	exitIfError(err)

	p, err := cli.NewDefaultPrompt(os.Args[1], &o)
	exitIfError(err)

	exitIfError(p.Run())

	item := p.Selection()
	if item == nil {
		return
	}
	command := cli.SuggestCommand(item)

	cli.ConfirmCommand(command)

}

func exitIfError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
