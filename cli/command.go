package cli

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func runCommand(c string, sudo bool) error {
	args := strings.Fields(c)
	if sudo {
		args = []string{"sudo"}
		args = append(args, strings.Fields(c)...)
	}
	cmd := exec.Command(args[0], args[1:]...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Fprintf(os.Stdout, "%s", out)
	return nil
}

func pipeCommands(commands []string, sudo bool) error {
	cmds := make([]*exec.Cmd, 0)
	for _, c := range commands {
		args := strings.Fields(strings.TrimSpace(c))
		if sudo {
			args = append([]string{"sudo"}, args...)
		}
		cmds = append(cmds, exec.Command(args[0], args[1:]...))
	}

	return execute(cmds...)
}

func execute(stack ...*exec.Cmd) error {
	var b bytes.Buffer
	var errBuf bytes.Buffer
	pipes := make([]*io.PipeWriter, len(stack)-1)
	i := 0
	for ; i < len(stack)-1; i++ {
		reader, writer := io.Pipe()
		stack[i].Stdout = writer
		stack[i].Stderr = &errBuf
		stack[i+1].Stdin = reader
		pipes[i] = writer
	}
	stack[i].Stdout = &b
	stack[i].Stderr = &errBuf

	if err := call(stack, pipes); err != nil {
		io.Copy(os.Stderr, &errBuf)
		return err
	}
	_, err := io.Copy(os.Stdout, &b)
	return err
}

func call(stack []*exec.Cmd, pipes []*io.PipeWriter) (err error) {
	if stack[0].Process == nil {
		if err = stack[0].Start(); err != nil {
			return err
		}
	}
	if len(stack) > 1 {
		if err = stack[1].Start(); err != nil {
			return err
		}
		defer func() {
			if err == nil {
				pipes[0].Close()
				err = call(stack[1:], pipes[1:])
			}
		}()
	}
	return stack[0].Wait()
}
