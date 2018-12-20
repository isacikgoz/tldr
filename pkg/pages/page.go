package pages

import (
	// "fmt"
	"regexp"
	"strings"

	"github.com/fatih/color"
)

var (
	rarg = regexp.MustCompile(`{{.[^}}]+}}`)

	bold  = color.New(color.Bold)
	blue  = color.New(color.FgBlue)
	red   = color.New(color.FgRed)
	cyan  = color.New(color.FgCyan)
	white = color.New(color.FgWhite)
)

type Page struct {
	Name string
	Desc string
	Tips []*Tip
}

type Tip struct {
	Desc string
	Cmd  *Command
}

type Command struct {
	Command string
	Args    []string
}

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
					Args:    rarg.FindAllString(c, -1),
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
	s := p.Name + "\n" + p.Desc
	return s
}

func (t *Tip) String() string {
	s := t.Desc
	return s
}

func (c *Command) String() string {
	s := c.Command
	return s
}

func (p *Page) Display() string {
	s := bold.Sprint(p.Name) + "\n\n" + p.Desc
	return s
}

func (t *Tip) Display() string {
	s := "- " + blue.Sprint(t.Desc) + "\n" + t.Cmd.Display()
	return s
}

func (c *Command) Display() string {
	s := c.Command
	for _, arg := range c.Args {
		t := arg[2:]
		t = t[:len(t)-2]
		s = strings.Replace(s, arg, white.Sprint(bold.Sprint(t)), 1)
	}
	return "    " + s + "\n"
}
