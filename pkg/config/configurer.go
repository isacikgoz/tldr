package config

import (
	"fmt"
	"runtime"

	log "github.com/sirupsen/logrus"
)

func StartUp() error {
	fmt.Printf("%s\n", initialMessage())
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
		log.Warn("Operating system couldn't be recognized")
	}
	return n
}
