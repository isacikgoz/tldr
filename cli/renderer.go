package cli

import (
	"github.com/fatih/color"
	"github.com/isacikgoz/gitin/term"
	"github.com/isacikgoz/tldr/pages"
)

func renderItem(item interface{}, matches []int, selected bool) [][]term.Cell {
	var tipText [][]term.Cell

	tip := item.(*pages.Tip)

	var line1 []term.Cell
	if selected {
		line1 = append(line1, term.Cprint("> ", color.FgCyan)...)
	} else {
		line1 = append(line1, term.Cprint("  ", color.FgWhite)...)
	}
	line1 = append(line1, term.Cprint(tip.String(), color.FgHiBlue)...)
	line2 := term.Cprint("  "+tip.Cmd.String(), color.FgHiCyan)

	tipText = append(tipText, line1)
	tipText = append(tipText, line2)
	return tipText
}

func information(item interface{}) [][]term.Cell {
	return [][]term.Cell{}
}
