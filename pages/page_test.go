package pages

import (
	"testing"
)

func TestParsePage(t *testing.T) {
	testPage := "# cd \n\n" +
		"> Change the current working directory.\n\n" +
		"- Go to the given directory:\n\n" +
		"`cd {{path/to/directory}}`\n\n" +
		"- Go to home directory of current user:\n\n" +
		"`cd`\n\n" +
		"- Go up to the parent of the current directory:\n\n" +
		"`cd ..`\n\n" +
		"- Go to the previously chosen directory:\n\n" +
		"`cd -`\n"

	page := ParsePage(testPage)
	if page == nil {
		t.Fatal("could not generate page")
	}
	if len(page.Tips) != 4 {
		t.Fatal("could not generate tips as expected")
	}
}

func TestQueryCommon(t *testing.T) {
	// var tests = []struct {
	// 	input    string
	// 	expected *Page
	// }{
	// 	{"curl", &Page{Name: "curl"}},
	// 	{"wget", &Page{Name: "wget"}},
	// 	{"man", &Page{Name: "man"}},
	// }
	// for _, test := range tests {
	// 	if p, err := queryCommon(test.input); p.Name != test.expected.Name && err != nil {
	// 		t.Errorf("Test Failed: {%s} inputted, {%s} expected, recieved: {%s}", test.input, test.expected.Name, p.Name)
	// 	}
	// }
}

func TestQueryOS(t *testing.T) {
	// var tests = []struct {
	// 	osname   string
	// 	input    string
	// 	expected *Page
	// }{
	// 	{"osx", "brew", &Page{Name: "brew"}},
	// 	{"linux", "apt", &Page{Name: "apt"}},
	// }
	// for _, test := range tests {
	// 	if config.OSName() == test.osname {
	// 		if p, err := queryOS(test.input); p.Name != test.expected.Name && err != nil {
	// 			t.Errorf("Test Failed: {%s} inputted, {%s} expected, recieved: {%s}", test.input, test.expected.Name, p.Name)
	// 		}
	// 	}
	// }
}
