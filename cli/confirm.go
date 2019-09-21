package cli

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

// ConfirmCommand asks user for confirmation
func ConfirmCommand(command string) error {
	green := color.New(color.FgGreen, color.Bold)
	fmt.Print(green.Sprint("? "))

	fmt.Print(command)
	fmt.Print(" (Y/n) ")
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	if s.Text() == "Y" || s.Text() == "y" || s.Text() == "" {
		args := strings.Fields(command)
		if s.Text() == "Y!" || s.Text() == "y!" {
			args = []string{"sudo"}
			args = append(strings.Fields(command))
		}
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			return err
		}
	}
	return nil
}
