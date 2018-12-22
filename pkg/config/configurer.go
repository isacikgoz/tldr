package config

import (
	"fmt"
	"os"
	"runtime"
)

// StartUp
func StartUp(clear, update bool) error {
	ok, _ := exists(SourceDir)
	// is staled, first check for internet connectivity, we don't want to
	// existing source if so
	if st, _ := staled(); st || !ok {
		if err := Clear(); err != nil {
			return err
		}
	}
	if clear {
		err := Clear()
		if err != nil {

		}
		os.Exit(0)
	} else if update {
		err := PullSource()
		if err != nil {

		}
		os.Exit(0)
	}
	return nil
}

// OSName is the running program's operating system
func OSName() (n string) {
	switch osname := runtime.GOOS; osname {
	case "windows":
		n = osname
	case "darwin":
		n = "osx"
	case "linux":
		n = osname
	case "solaris":
		n = "sunos"
	default:
		fmt.Println("Operating system couldn't be recognized")
		os.Exit(1)
	}
	return n
}

// exists checks if the file exists
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
