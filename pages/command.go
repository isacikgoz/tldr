package pages

import (
	"strings"
)

// Command is the representation of a tip's command suggestion
type Command struct {
	Command string
	Args    []string
}

func (c *Command) String() string {
	s := c.Command
	return s
}

// Display returns colored and indented text for rendering output
func (c *Command) Display() string {
	s := c.Command
	for _, arg := range c.Args {
		s = strings.Replace(s, "{{"+arg+"}}", cyan.Sprint(arg), 1)
	}
	return "   " + s + "\n"
}

func stripCommandArgs(in string) []string {
	out := make([]string, 0)
	ir := []rune(in)
	argStart := 0
	argEnd := 0
	for ix := 0; ix < len(ir); {
		if ir[ix] == '{' && ir[ix+1] == '{' {
			argStart = ix + 2
			ix = ix + 2
		} else if ir[ix] == '}' && ir[ix+1] == '}' {
			argEnd = ix
			out = append(out, in[:argEnd][argStart:])
			ix = ix + 2
		} else {
			ix++
		}
	}
	return out
}
