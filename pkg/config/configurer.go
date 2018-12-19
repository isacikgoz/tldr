package config

import (
	"fmt"
)

func StartUp() error {
	fmt.Printf("%s\n", initialMessage())
	return nil
}
