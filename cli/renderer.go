package cli

import (
	"fmt"
	"strings"

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
		line1 = append(line1, term.Cprint("  ")...)
	}
	line1 = append(line1, term.Cprint(tip.String(), color.FgHiBlue)...)

	line2 := term.Cprint("  ", color.FgWhite)

	s := fmt.Sprint(tip.Cmd)
	start := 0
	index := ""
	if len(tip.Cmd.Args) == 0 {
		line2 = append(line2, term.Cprint(s)...)
	}
	for _, arg := range tip.Cmd.Args {
		s = strings.Replace(s, "{{"+arg+"}}", arg, 1)            // fix the arg
		ts := s[start:strings.Index(s, arg)]                     // w/o arg
		line2 = append(line2, term.Cprint(ts)...)                // append temp
		line2 = append(line2, term.Cprint(arg, color.FgCyan)...) // append arg
		index += ts + arg                                        // to keep the index of where to cut
		start = len(index)
	}

	tipText = append(tipText, line1)
	tipText = append(tipText, line2)
	return tipText
}

func information(item interface{}) [][]term.Cell {
	return [][]term.Cell{}
}
