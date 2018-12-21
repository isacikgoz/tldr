package pages

import ()

type Tip struct {
	Desc string
	Cmd  *Command
}

func (t *Tip) String() string {
	s := t.Desc
	return s
}

func (t *Tip) Display() string {
	s := "- " + blue.Sprint(t.Desc) + "\n" + t.Cmd.Display()
	return s
}
