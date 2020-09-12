package cli

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"strings"
)

func pipeCommands(commands []string, sudo bool) error {
	cmds := make([]*exec.Cmd, len(commands))
	for _, c := range commands {
		args := strings.Fields(c)
		cmds = append(cmds, exec.Command(args[0], args[1:]...))
	}

	return command(cmds...)
}

func runCommand(c string, sudo bool) error {
	args := strings.Fields(c)
	if sudo {
		args = []string{"sudo"}
		args = append(args, strings.Fields(c)...)
	}
	cmd := exec.Command(args[0], args[1:]...)
	return command(cmd)
}

func command(stack ...*exec.Cmd) error {
	var stderr bytes.Buffer
	pipeStack := make([]*io.PipeWriter, len(stack)-1)
	i := 0
	for ; i < len(stack)-1; i++ {
		inPipe, outPipe := io.Pipe()
		stack[i].Stdout = outPipe
		stack[i].Stderr = &stderr
		stack[i+1].Stdin = inPipe
		pipeStack[i] = outPipe
	}
	stack[i].Stdout = os.Stdout
	stack[i].Stderr = &stderr

	return call(stack, pipeStack)
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
