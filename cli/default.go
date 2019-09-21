package cli

import (
	"fmt"

	"github.com/isacikgoz/gitin/prompt"
	"github.com/isacikgoz/gitin/term"
	"github.com/isacikgoz/tldr/pages"
)

// DefaultPrompt is the default tldr prompt
type DefaultPrompt struct {
	prompt *prompt.Prompt
	item   interface{}
}

// NewDefaultPrompt creates a prompt for tldr app
func NewDefaultPrompt(command string, opts *prompt.Options) (*DefaultPrompt, error) {
	page, err := pages.Read([]string{command})
	if err != nil {
		return nil, fmt.Errorf("could not read page: %v", err)
	}

	fmt.Printf("%s", page.Desc)
	list, err := prompt.NewList(page.Tips, opts.LineSize)
	if err != nil {
		return nil, fmt.Errorf("could not create list: %v", err)
	}
	d := &DefaultPrompt{}
	p := prompt.Create("", opts, list,
		prompt.WithItemRenderer(renderItem),
		prompt.WithInformation(information),
		prompt.WithSelectionHandler(d.selection),
	)
	p.SetExitMsg(defaultExitMessage(list))
	d.prompt = p
	return d, nil
}

// Run starts the prompt within
func (d *DefaultPrompt) Run() error {
	return d.prompt.Run()
}

// Selection returns the selected item
func (d *DefaultPrompt) Selection() interface{} {
	return d.item
}

// selection implements the prompt.selectionHandlerFunc interface
func (d *DefaultPrompt) selection(item interface{}) error {
	d.item = item
	var cells [][]term.Cell
	cells = append(cells, term.Cprint(""))
	cells = append(cells, renderItem(item, nil, false)...)
	d.prompt.SetExitMsg(cells)
	d.prompt.Stop()
	return nil
}

func defaultExitMessage(l *prompt.List) [][]term.Cell {
	var cells [][]term.Cell
	cells = append(cells, term.Cprint(""))
	items, _ := l.Items()
	for _, item := range items {
		cells = append(cells, renderItem(item, nil, false)...)
	}
	return cells
}
