package config

import (
	"runtime"
	"testing"
)

func TestStartUp(t *testing.T) {
	var tests = []struct {
		input_1 bool
		input_2 bool
		input_3 string
	}{
		{false, false, ""},
	}
	for _, test := range tests {
		if err := StartUp(test.input_1, test.input_2, test.input_3); err != nil {
			t.Errorf("Test Failed: {%t, %t} inputted, recieved: {%s}", test.input_1,
				test.input_2, err.Error())
		}
	}
}

func TestOSName(t *testing.T) {
	if osName := OSName(); osName != runtime.GOOS {
		if osName == "osx" {
			return
		}
		t.Errorf("Test Failed: {%s} expected, recieved: {%s}", osName, runtime.GOOS)
	}
}

func TestExists(t *testing.T) {
	var tests = []struct {
		input    string
		expected bool
	}{
		{"./configurer.go", true},
		{"/path/to/file", false},
	}
	for _, test := range tests {
		if output, err := exists(test.input); output != test.expected || err != nil {
			t.Errorf("Test Failed: {%s} inputted, {%t} expected, recieved: {%t}", test.input, test.expected, output)
		}
	}
}
