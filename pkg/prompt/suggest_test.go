package prompt

import (
	"fmt"
	"testing"

	prompt "github.com/c-bata/go-prompt"
	"github.com/c-bata/go-prompt/completer"
)

func TestFileExtCompleterFunc(t *testing.T) {
	var tests = []struct {
		input    prompt.Document
		expected []prompt.Suggest
	}{
		{
			prompt.Document{
				Text: "File.go",
			},
			[]prompt.Suggest{
				prompt.Suggest{Text: "File_1.go", Description: ""},
				prompt.Suggest{Text: "File_2", Description: ""},
				prompt.Suggest{Text: "File_", Description: ""},
				prompt.Suggest{Text: "File_3", Description: ""},
			},
		},
	}

	for _, test := range tests {
		if output := fileExtCompleterFunc(test.input); len(output) != len(test.expected) {
			for _, d := range output {
				fmt.Println(d.Text)
			}
			t.Errorf("Test Failed: {%s} inputted, {%d} expected, recieved: {%d}", test.input.Text, len(test.expected), len(output))
		}
	}
}

func TestSuggestCompleterFunct(t *testing.T) {
	var tests = []struct {
		input    string
		expected completer.FilePathCompleter
	}{
		{"archived.7z", completer.FilePathCompleter{}},
		{"path/to/file", completer.FilePathCompleter{}},
		// {"8", completer.FilePathCompleter{}},
	}
	for _, test := range tests {
		if _, err := suggestCompleterFunc(test.input); err != nil {
			t.Errorf("Failed, err: %s", err.Error())
		}
	}
}

func TestGetFileExtension(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		{"archived.7z", ".7z"},
		{"path/to/file", ""},
		{"file.tar.gz", ".gz"},
		{"main.d", ".d"},
		{"compressed.zip", ".zip"},
		{"...", ""},
		{"7", ""},
	}
	for _, test := range tests {
		if output := getFileExtension(test.input); output != test.expected {
			t.Errorf("Test Failed: {%s} inputted, {%s} expected, recieved: {%s}", test.input, test.expected, output)
		}
	}
}

func TestRemoveDuplicates(t *testing.T) {
	var tests = []struct {
		input    []prompt.Suggest
		expected []prompt.Suggest
	}{
		{
			[]prompt.Suggest{
				prompt.Suggest{Text: "File_1", Description: ""},
				prompt.Suggest{Text: "File_2", Description: ""},
				prompt.Suggest{Text: "File_2", Description: ""},
				prompt.Suggest{Text: "File_1", Description: ""},
			},
			[]prompt.Suggest{
				prompt.Suggest{Text: "File_1", Description: ""},
				prompt.Suggest{Text: "File_2", Description: ""},
			},
		},
	}

	for _, test := range tests {
		if output := removeDuplicates(test.input); len(output) != len(test.expected) {
			t.Errorf("Test Failed: {%d} inputted, {%d} expected, recieved: {%d}", len(test.input), len(test.expected), len(output))
		}
	}
}
