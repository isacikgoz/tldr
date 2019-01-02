package prompt

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	prompt "github.com/c-bata/go-prompt"
	"github.com/c-bata/go-prompt/completer"
	"github.com/isacikgoz/survey"
	"github.com/isacikgoz/tldr/pkg/pages"
	xterm "golang.org/x/crypto/ssh/terminal"
	"gopkg.in/AlecAivazis/survey.v1/core"
	"gopkg.in/AlecAivazis/survey.v1/terminal"
)

// Prompt struct is responsible for maintaining the life cycle of a tldr man page
type Prompt struct {
	Page      *pages.Page
	Questions []*survey.Question
}

// New creates a new *prompt.Prompt obj. from a tldr Page
func New(p *pages.Page) *Prompt {
	return &Prompt{
		Page: p,
	}
}

// RenderPage prints the tldr man page as is
func (p *Prompt) RenderPage(static bool) error {
	if static {
		fmt.Println(p.Page.Display())
		for _, t := range p.Page.Tips {
			fmt.Println("-" + t.Display())
		}
		return errors.New("")
	}
	options := make([]string, 0)

	for _, t := range p.Page.Tips {
		options = append(options, t.Display())
	}
	core.ErrorTemplate = ""

	var mxRows int
	_, height, err := xterm.GetSize(0)
	if err != nil || len(p.Page.Tips) > 10 {
		mxRows = 7
	} else {
		mxRows = (height - 7) / (3)
		if mxRows < 3 {
			mxRows = 3
		}
	}

	// if len(p.Page.Tips) > 7 {
	// 	mxRows = 7
	// } else {
	// 	mxRows = len(p.Page.Tips)
	// }
	// genereate questions to ask
	p.Questions = []*survey.Question{
		{
			Name: "Tip",
			Prompt: &survey.Select{
				Message:  p.Page.Display(),
				Options:  options,
				VimMode:  true,
				PageSize: mxRows,
				Help:     "Select a command template to fill args.",
			},
			Validate: survey.Required,
		},
	}
	survey.SelectQuestionTemplate = `
{{- color "default+hb"}}{{ .Message }}{{ .FilterMessage }}{{color "reset"}}
{{- if .ShowAnswer}}{{color "cyan"}} {{.Answer}}{{color "reset"}}{{"\n"}}
{{- else}}
  {{- "   "}}(Use{{" "}}{{- color "cyan"}}arrows{{- color "reset"}}` +
		` to move,{{" "}}{{- color "cyan"}}type{{- color "reset"}} to search or{{" "}}` +
		`{{- color "red"}}ctrl+c{{- color "reset"}} to return{{- if and .Help (not .ShowHelp)}}` +
		`. {{""}}{{- color "green"}}{{ HelpInputRune }}{{- color "reset"}} for more{{end}})
		{{- if .ShowHelp }}{{"\n"}}{{- color "green"}}{{ HelpIcon }}{{"  "}}{{ .Help }}{{color "reset"}}{{end}}
  {{- "\n\n"}}
  {{- range $ix, $choice := .PageEntries}}
    {{- if eq $ix $.SelectedIndex}}{{color "blue+b"}}{{ "*" }} {{else}}{{color "default"}}  {{end}}
    {{- $choice}}
    {{- color "reset"}}{{"\n"}}
  {{- end}}
{{- end}}`

	terminal.InterruptErr = errors.New("\x0d")
	return nil
}

// Selection is where user interaction starts, hence we can *pages.Tip to iterate
// user interaction
func (p *Prompt) Selection() (t *pages.Tip, err error) {
	answer := struct {
		Tip string
	}{}
	// bug: https://github.com/AlecAivazis/survey/issues/101
	// make terminal not line wrap
	fmt.Printf("\x1b[?7l")
	// defer restoring line wrap
	defer fmt.Printf("\x1b[?7h")
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

// GenerateCommand generates command form *pages.Tip
func (p *Prompt) GenerateCommand(t *pages.Tip) (string, error) {
	answers := make([]string, 0)
	for _, arg := range t.Cmd.Args {
		// cs, _ := suggestCompleterFunc(arg)
		ext = getFileExtension(arg)
		answers = append(answers, prompt.Input(
			"$"+" "+arg+" -> ",
			fileExtCompleterFunc,
			prompt.OptionPreviewSuggestionTextColor(prompt.Cyan),
			prompt.OptionInputTextColor(prompt.Cyan),
			prompt.OptionAddKeyBind(prompt.KeyBind{
				Key: prompt.ControlC,
				Fn: func(buf *prompt.Buffer) {
					os.Exit(0)
					return
				}}),
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
		// since we have double curlybraces on args
		fs = strings.Replace(fs, "{{"+arg+"}}", answers[i], 1)
	}
	return fs, nil
}

// Run gets final confirmation from user and executes the command with its args
func (p *Prompt) Run(command string) error {
	run := false
	confirm := &survey.Confirm{
		Message: command,
		Default: true,
		Help:    "type y! to run command with",
	}
	survey.ConfirmQuestionTemplate = `
{{- if .ShowHelp }}{{- color "cyan"}}{{ HelpIcon }} {{ .Help }} {{color "red"}}sudo{{color "reset"}}{{"\n"}}{{end}}
{{- color "green+hb"}}{{ QuestionIcon }} {{color "reset"}}
{{- if .ShowSudo }}{{- color "yellow"}}sudo{{color "reset"}} {{end}}
{{- color "default+hb"}}{{ .Message }} {{color "reset"}}
{{- if .Answer}}
  {{- color "cyan"}}{{.Answer}}{{color "reset"}}{{- if .ShowSudo }}{{- color "yellow"}}!{{color "reset"}}{{end}}{{"\n"}}
{{- else }}
  {{- color "white"}}{{if .Default}}(Y/n) {{else}}(y/N) {{end}}{{color "reset"}}
{{- end}}`
	survey.AskOne(confirm, &run, nil)
	if confirm.AddSudo {
		command = "sudo " + command
	}

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
