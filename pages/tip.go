package pages

// Tip is the list item of a tldr page
type Tip struct {
	Desc string
	Cmd  *Command
}

func (t *Tip) String() string {
	s := t.Desc
	return s
}
