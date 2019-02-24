package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"

	"github.com/fatih/color"
)

var settingFile = os.TempDir() + "tldr"

// StartUp
func StartUp(clear, update bool, osname string) error {
	ok, _ := exists(SourceDir)
	// is staled, first check for internet connectivity, we don't want to
	// existing source if so
	if st, _ := staled(); st {
		yellow := color.New(color.FgYellow)
		fmt.Println(yellow.Sprint("TLDR repository is older than 2 weeks, consider updating it with -u option."))
	}

	if clear || !ok {
		err := Clear()
		if err != nil {

		}
		os.Exit(0)
	} else if update {
		err := PullSource()
		if err != nil {

		}
		os.Exit(0)
	} else if osname != "" {
		switch osname {
		case "windows", "osx", "linux", "sunos":
			if err := ioutil.WriteFile(settingFile, []byte(osname), 0644); err != nil {
				fmt.Printf("Fail save manual os: %v", err)
			}
		default:
			fmt.Printf("%s is invalid osname, choose one from [windows|linux|osx|sunos].", osname)
		}
		os.Exit(0)
	}
	return nil
}

func PageOSName() (n string) {
	if _, err := os.Stat(settingFile); err == nil {
		if content, err := ioutil.ReadFile(settingFile); err == nil && len(content) > 0 {
			return string(content)
		}
	}
	return OSName()
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

func PrintLogo() {
	fmt.Printf("%s\n", colorLogo())
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
