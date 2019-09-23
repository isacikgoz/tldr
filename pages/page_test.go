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
