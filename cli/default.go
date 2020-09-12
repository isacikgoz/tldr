package cli

import (
	"context"
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
func NewDefaultPrompt(pgs []string, opts *prompt.Options, static, random bool) (*DefaultPrompt, error) {
	var page *pages.Page
	var err error
	if random {
		page, err = pages.QueryRandom()
	} else if len(pgs) == 0 {
		page, err = pages.ReadAll()
	} else {
		page, err = pages.Read(pgs)
	}
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
	if static {
		if err := printStatic(page.Tips); err != nil {
			return nil, err
		}
		return d, nil
	}
	d.prompt = p
	return d, nil
}

// Run starts the prompt within
func (d *DefaultPrompt) Run(ctx context.Context) error {
	if d.prompt == nil {
		return nil
	}
	return d.prompt.Run(ctx)
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
