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
	old       = ".old"
)

// Clear removes the existing tldr directory. TODO: maybe require user to confirm
func Clear() error {
	os.Rename(SourceDir, SourceDir+old)

	PrintLogo()
	_, err := git.PlainClone(dir, false, &git.CloneOptions{
		URL:               giturl,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
		Progress:          os.Stdout,
	})
	if err == nil {
		fmt.Printf("Successfully cloned into: %s\n", dir)
		os.RemoveAll(SourceDir + old)
	}
	if err != nil {
		os.RemoveAll(SourceDir)
		os.Rename(SourceDir+old, SourceDir)
	}
	return err
}

// Pulls the github.com/tldr-pages/tldr repository
func PullSource() error {

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
		fmt.Printf("%s\n", err.Error())
	} else {
		fmt.Printf("Successfully pulled into: %s\n", dir)
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

// staled checks if the source folder is older than two weeks
func staled() (bool, error) {
	file, err := os.Open(SourceDir)
	if err != nil {
		return false, err
	}
	fstat, err := file.Stat()
	if err != nil {
		return false, err
	}

	diff := time.Now().Sub(fstat.ModTime())
	// Two week update time, seems fair.
	if diff > 24*7*2*time.Hour {
		return true, nil
	}
	return false, nil
}
