package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

// ConfirmCommand asks user for confirmation
func ConfirmCommand(command string) error {
	green := color.New(color.FgGreen, color.Bold)
	fmt.Print(green.Sprint("? "))

	fmt.Print(command)
	fmt.Print(" [Y/n] ")
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	input := s.Text()
	sudo := strings.HasSuffix(input, "!")
	if strings.HasPrefix(input, "Y") || strings.HasPrefix(input, "y") || s.Text() == "" {

		cmds := strings.Split(strings.TrimSpace(command), "|")
		if len(cmds) >= 2 {
			return pipeCommands(cmds, sudo)
		}
		return runCommand(command, sudo)

	}
	return nil
}
