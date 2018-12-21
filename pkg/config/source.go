package config

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/src-d/go-git.v4"
)

const (
	giturl = "https://github.com/tldr-pages/tldr.git"
)

var (
	dir       = DataDir() + "/tldr"
	SourceDir = dir
)

func Clear() error {
	os.RemoveAll(SourceDir)

	fmt.Printf("%s\n", initialMessage())
	_, err := git.PlainClone(dir, false, &git.CloneOptions{
		URL:               giturl,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
		Progress:          os.Stdout,
	})
	if err == nil {
		fmt.Printf("Successfully cloned into: %s\n", dir)
	}
	return err
}

func PullSource() error {

	fmt.Printf("%s\n", initialMessage())
	r, err := git.PlainOpen(dir)
	if err != nil {
		return err
	}
	w, err := r.Worktree()
	if err != nil {
		return err
	}
	err = w.Pull(&git.PullOptions{
		RemoteName: "origin",
		Progress:   os.Stdout,
	})
	if err != nil {
		fmt.Printf(" %s\n", err.Error())
	} else {
		fmt.Printf("Successfully cloned into: %s\n", dir)
	}
	return err
}

// returns OS dependent data dir. see XDG Base Directory Specification:
// https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html
func DataDir() (d string) {
	switch osname := OSName(); osname {
	case "windows":
		d = os.Getenv("APPDATA")
	case "osx":
		d = os.Getenv("HOME") + "/Library/Application Support"
	case "linux":
		d = os.Getenv("HOME") + "/.local/share"
	case "solaris":
		d = os.Getenv("HOME") + "/.local/share"
	default:
		fmt.Println("Operating system couldn't be recognized")
	}
	return d
}

func staled() bool {
	file, _ := os.Open(SourceDir)
	fstat, _ := file.Stat()
	// now := time.Now()

	diff := time.Now().Sub(fstat.ModTime())
	if diff > 24*7*2*time.Hour {
		return true
	}
	return false
}
