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
		t := arg[2:]
		t = t[:len(t)-2]
		s = strings.Replace(s, arg, cyan.Sprint(t), 1)
	}
	return "   " + s + "\n"
}
