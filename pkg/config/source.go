package config

import (
	"fmt"
	"os"

	"gopkg.in/src-d/go-git.v4"
)

const (
	giturl    = "https://github.com/tldr-pages/tldr.git"
	directory = "/home/isacikgoz/.tldr"
)

func CloneSource() error {
	_, err := git.PlainClone(directory, false, &git.CloneOptions{
		URL:               giturl,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
		Progress:          os.Stdout,
	})
	if err == nil {
		fmt.Printf("Successfully cloned into: %s\n", directory)
	}
	return err
}

func PullSource() error {
	r, err := git.PlainOpen(directory)
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
		fmt.Printf("Successfully cloned into: %s\n", directory)
	}
	return err
}
