package config

import (
	"gopkg.in/src-d/go-git.v4"

	"os"
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
	return w.Pull(&git.PullOptions{RemoteName: "origin"})
}
