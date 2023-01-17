package config

import (
	"testing"

	git "github.com/go-git/go-git/v5"
)

func TestClear(t *testing.T) {
	if err := Clear(); err != nil {
		t.Errorf("Test Failed: %s", err.Error())
	}
}

func TestPullSource(t *testing.T) {
	if err := PullSource(); err != nil && err != git.NoErrAlreadyUpToDate {
		t.Errorf("Test Failed: %s", err.Error())
	}
}

func TestDataDir(t *testing.T) {
	if len(DataDir()) <= 0 {
		t.Errorf("Test Failed.")
	}

}

func TestStaled(t *testing.T) {
	if _, err := staled(); err != nil {
		t.Errorf("Test Failed: %s", err.Error())
	}
}
