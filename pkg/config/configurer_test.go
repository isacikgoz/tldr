package config

import (
	"os"
	"runtime"
	"testing"
)

func TestStartUp(t *testing.T) {
	var tests = []struct {
		input_1 bool
		input_2 bool
	}{
		{false, false},
	}
	for _, test := range tests {
		if err := StartUp(test.input_1, test.input_2); err != nil {
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

func TestPageOSName(t *testing.T) {
	if err := os.Setenv("PAGEOS", "linux"); err != nil {
		t.Error("Test Failed: fail to mock PAGEOS")
	}
	if osName := PageOSName(); osName != "linux" {
		t.Errorf("Test Failed: {%s} expected when PAGEOS is set, received: {%s}", "linux", osName)
	}
}
