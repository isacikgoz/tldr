package pages

import (
	// "fmt"
	"regexp"
	"strings"
)

var (
	rarg = regexp.MustCompile(`{{.[^}}]+}}`)
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
			d = line[2:] + "\n"
		} else {
			c = ln
			break
		}
	}
	tips := make([]*Tip, 0)
	for ln := c; ln < len(l); {
		line := l[ln]
		if len(line) > 0 && line[0] == '-' {
			d := line[2:]
			c := l[ln+2]
			var cmd *Command
			if len(c) > 0 && c[0] == '`' {
				cmd = &Command{
					Command: c,
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
	for _, t := range p.Tips {
		s = s + "\n" + t.Desc + "\n"
		for _, args := range t.Cmd.Args {
			for _, arg := range args {
				s = s + string(arg)
			}
		}
	}
	return s
}
