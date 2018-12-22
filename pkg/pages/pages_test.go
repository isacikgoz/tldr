package pages

import (
	// "strings"
	"testing"

	"github.com/isacikgoz/tldr/pkg/config"
)

func TestRead(t *testing.T) {
	// var tests = []struct {
	// 	input    []string
	// 	expected *Page
	// }{
	// 	{[]string{"git", "add"}, &Page{Name: "git add"}},
	// 	{[]string{"git"}, &Page{Name: "git"}},
	// }
	// for _, test := range tests {
	// 	if p, err := Read(test.input); p != nil && p.Name != test.expected.Name && err != nil {
	// 		t.Errorf("Test Failed: {%s} inputted, {%s} expected, recieved: {%s}", strings.Join(test.input, ""), test.expected.Name, p.Name)
	// 	}
	// }
}

func TestQueryCommon(t *testing.T) {
	var tests = []struct {
		input    string
		expected *Page
	}{
		{"curl", &Page{Name: "curl"}},
		{"wget", &Page{Name: "wget"}},
		{"man", &Page{Name: "man"}},
	}
	for _, test := range tests {
		if p, err := queryCommon(test.input); p.Name != test.expected.Name && err != nil {
			t.Errorf("Test Failed: {%s} inputted, {%s} expected, recieved: {%s}", test.input, test.expected.Name, p.Name)
		}
	}
}

func TestQueryOS(t *testing.T) {
	var tests = []struct {
		osname   string
		input    string
		expected *Page
	}{
		{"osx", "brew", &Page{Name: "brew"}},
		{"linux", "apt", &Page{Name: "apt"}},
	}
	for _, test := range tests {
		if config.OSName() == test.osname {
			if p, err := queryOS(test.input); p.Name != test.expected.Name && err != nil {
				t.Errorf("Test Failed: {%s} inputted, {%s} expected, recieved: {%s}", test.input, test.expected.Name, p.Name)
			}
		}
	}
}
