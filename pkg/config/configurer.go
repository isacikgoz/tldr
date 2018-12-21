package config

import (
	"os"
	"runtime"

	log "github.com/sirupsen/logrus"
)

func StartUp() error {
	ok, _ := exists(SourceDir)
	if !ok {
		return Clear()
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
		log.Warn("Operating system couldn't be recognized")
	}
	return n
}

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
