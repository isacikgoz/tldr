package prompt

import (
	"testing"

	"github.com/isacikgoz/tldr/pkg/pages"
)

func TestNew(t *testing.T) {
	var tests = []struct {
		input    *pages.Page
		expected *Prompt
	}{
		{&pages.Page{
			Name: "Page-1",
			Desc: "Cool page indeed",
		}, &Prompt{}},
	}
	for _, test := range tests {
		output := New(test.input)
		if output.Page != test.input {
			t.Error("Could not create prompt")
		}
	}
}
