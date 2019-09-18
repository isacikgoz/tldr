package cli

import (
	"fmt"

	"github.com/isacikgoz/gitin/prompt"
	"github.com/isacikgoz/tldr/pages"
)

// DefaultPrompt creates a prompt for tldr app
func DefaultPrompt(command string, opts *prompt.Options) (*prompt.Prompt, error) {

	page, err := pages.Read([]string{command})
	if err != nil {
		return nil, fmt.Errorf("could not read page: %v", err)
	}

	list, err := prompt.NewList(page.Tips, opts.LineSize)
	if err != nil {
		return nil, fmt.Errorf("could not create list: %v", err)
	}

	p := prompt.Create(command, opts, list)

	return p, nil
}
