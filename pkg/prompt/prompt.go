package prompt

import (
	"os"
	"os/exec"
	"strings"

	prompt "github.com/c-bata/go-prompt"
	"github.com/c-bata/go-prompt/completer"
	"github.com/isacikgoz/tldr/pkg/pages"
	"gopkg.in/AlecAivazis/survey.v1"
	"gopkg.in/AlecAivazis/survey.v1/core"
)

type Prompt struct {
	Page      *pages.Page
	Questions []*survey.Question
}

func New(p *pages.Page) *Prompt {
	return &Prompt{
		Page: p,
	}
}
func (p *Prompt) RenderPage() error {

	options := make([]string, 0)

	for _, t := range p.Page.Tips {
		options = append(options, t.Display())
	}
	core.ErrorTemplate = ""
	// the questions to ask
	p.Questions = []*survey.Question{
		{
			Name: "Tip",
			Prompt: &survey.Select{
				Message: p.Page.Display(),
				Options: options,
			},
			Validate: survey.Required,
		},
	}
	return nil
}

func (p *Prompt) Selection() (t *pages.Tip, err error) {
	answer := struct {
		Tip string
	}{}

	// ask the question
	err = survey.Ask(p.Questions, &answer)

	if err != nil {
		return nil, err
	}

	var st *pages.Tip
	for _, t := range p.Page.Tips {
		if t.Display() == answer.Tip {
			st = t
		}
	}
	return st, err
}

func (p *Prompt) GenerateCommand(t *pages.Tip) (string, error) {
	// fmt.Print(st.Cmd.Display())
	answers := make([]string, 0)
	for _, arg := range t.Cmd.Args {
		cs, _ := suggestCompleterFunc(arg)
		answers = append(answers, prompt.Input(
			"$"+" "+arg[:len(arg)-2][2:]+" -> ",
			cs.Complete,
			prompt.OptionPrefixTextColor(prompt.Cyan),
			prompt.OptionSuggestionBGColor(prompt.DarkGray),
			prompt.OptionAddKeyBind(prompt.KeyBind{
				Key: prompt.Escape,
				Fn: func(buf *prompt.Buffer) {
					return
				}}),
			prompt.OptionCompletionWordSeparator(completer.FilePathCompletionSeparator),
		))
	}

	fs := t.Cmd.String()
	for i, arg := range t.Cmd.Args {
		fs = strings.Replace(fs, arg, answers[i], 1)
	}
	return fs, nil
}

func (p *Prompt) Run(command string) error {
	run := false
	confirm := &survey.Confirm{
		Message: command,
		Default: true,
	}
	survey.AskOne(confirm, &run, nil)

	if run {
		args := strings.Fields(command)
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			return err
		}
	}
	return nil
}
