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

// Display returns colored and indented text for rendering output
func (t *Tip) Display() string {
	s := " " + blue.Sprint(t.Desc) + "\n" + t.Cmd.Display()
	return s
}
