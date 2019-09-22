package pages

import (
	"fmt"
	"strings"
)

// Page is the representation of a tldr page itself
type Page struct {
	Name string
	Desc string
	Tips []*Tip
}

// Tip is the list item of a tldr page
type Tip struct {
	Desc string
	Cmd  *Command
}

// Command is the representation of a tip's command suggestion
type Command struct {
	Command string
	Args    []string
}

// ParsePage parses from bare markdown string. Rather than parsing markdown itself
// initial implementation approach is stripping from a single string
func ParsePage(s string) *Page {
	l := strings.Split(s, "\n")

	n := l[0][2:]
	var d string
	var c int
	for ln := 2; ln < len(l); ln++ {
		line := l[ln]
		if len(line) > 0 && line[0] == '>' {
			d = d + line[2:] + "\n"
		} else {
			c = ln
			break
		}
	}

	tips := make([]*Tip, 0)
	for ln := c; ln < len(l); {
		line := l[ln]
		if len(line) > 0 && line[0] == '-' {
			// remove last rune then first two runes
			d := line[:len(line)-1][2:]
			c := l[ln+2]
			var cmd *Command
			if len(c) > 0 && c[0] == '`' {
				cmd = &Command{
					Command: c[:len(c)-1][1:],
					Args:    stripCommandArgs(c),
				}
				ln = ln + 2
			} else {
				break
			}
			tips = append(tips, &Tip{
				Desc: d,
				Cmd:  cmd,
			})
		}
		ln++
	}

	p := &Page{
		Name: n,
		Desc: d,
		Tips: tips,
	}
	return p
}

func (p *Page) String() string {
	return fmt.Sprintf("%s\n%s", p.Name, p.Desc)
}

func (t *Tip) String() string {
	return t.Desc
}

func (c *Command) String() string {
	return c.Command
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
