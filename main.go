package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	// "time"

	prompt "github.com/c-bata/go-prompt"
	"github.com/isacikgoz/tldr/pkg/config"
	"github.com/isacikgoz/tldr/pkg/pages"
	"gopkg.in/AlecAivazis/survey.v1"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	clone  = kingpin.Flag("clone", "clones the repository from github.com/tldr-pages/tldr").Short('c').Bool()
	update = kingpin.Flag("update", "pulls the latest commits from github.com/tldr-pages/tldr").Short('u').Bool()

	page = kingpin.Arg("command", "Name of the command.").String()
)

func main() {
	// start := time.Now()
	kingpin.Version("tldr++ version 0.0.1 (pre-release)")
	// parse the command line flag and options
	kingpin.Parse()
	if *clone {
		config.StartUp()
		err := config.CloneSource()
		if err != nil {

		}
		return
	}
	if *update {
		config.StartUp()
		err := config.PullSource()
		if err != nil {

		}
		return
	}
	p, err := pages.Read(*page)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}
	// fmt.Printf("%s\n", p.Display())
	// elapsed := time.Since(start)
	// fmt.Printf("Query finished in: %s\n", elapsed)

	options := make([]string, 0)

	for _, t := range p.Tips {
		options = append(options, t.Display())
	}

	// the questions to ask
	sr := []*survey.Question{
		{
			Name: "Tip",
			Prompt: &survey.Select{
				// Message: p.Display(),
				Options: options,
			},
			Validate: survey.Required,
		},
	}

	answer := struct {
		Tip string
	}{}

	// ask the question
	err = survey.Ask(sr, &answer)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var st *pages.Tip
	for _, t := range p.Tips {
		if t.Display() == answer.Tip {
			st = t
		}
	}
	// fmt.Print(st.Cmd.Display())
	answers := make([]string, 0)
	for _, arg := range st.Cmd.Args {
		answers = append(answers, prompt.Input(
			"$ "+arg[:len(arg)-2][2:]+" -> ",
			completer,
			// prompt.OptionLivePrefix(livePrefix),
			prompt.OptionSuggestionBGColor(prompt.DarkGray),
		))
	}

	fs := st.Cmd.String()
	for i, arg := range st.Cmd.Args {
		fs = strings.Replace(fs, arg, answers[i], 1)
	}

	fmt.Printf("Command: %s\n", fs)
}

func completer(t prompt.Document) []prompt.Suggest {
	files, _ := ioutil.ReadDir("./")
	s := make([]prompt.Suggest, 0)
	for _, f := range files {
		s = append(s, prompt.Suggest{
			Text: f.Name(),
		})
	}
	return s
}
