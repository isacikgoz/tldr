package config

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/fatih/color"
	env "github.com/kelseyhightower/envconfig"
)

type envConf struct {
	OS string
}

// StartUp
func StartUp(clear, update bool) error {
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
	}
	return nil
}

// OSName is the running program's operating system
func OSName() (n string) {
	var conf envConf
	var osname string
	err := env.Process("tldr", &conf)
	if err != nil {
		//log.Warn(err.Error())
	}
	if len(conf.OS) > 0 {
		osname = strings.ToLower(conf.OS)
	} else {
		osname = runtime.GOOS
	}
	switch osname {
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
