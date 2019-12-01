package commandProcessor

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func execPipeCommands(commands []*userInputCommand, pwd string) {
	var cmds []*systemCommand
	var currentReadPipe *io.PipeReader
	var buffer bytes.Buffer
	for index, command := range commands {
		cmd := exec.Command(command.inputCommand, command.options...)
		cmd.Dir = pwd
		systemCmd := newSystemCommand(cmd)
		if index > 0 {
			systemCmd.cmd.Stdin = currentReadPipe
		} else {
			systemCmd.cmd.Stdin = os.Stdin
		}
		r, w := io.Pipe()
		if index == len(commands)-1 {
			systemCmd.cmd.Stdout = &buffer
			systemCmd.writer = nil
		} else {
			systemCmd.cmd.Stdout = w
			systemCmd.writer = w
		}
		currentReadPipe = r
		cmds = append(cmds, systemCmd)
	}

	for _, systemCmd := range cmds {
		err := systemCmd.cmd.Start()
		if err != nil {
			fmt.Printf("%s\n", err)
		}
	}

	for _, systemCmd := range cmds {
		systemCmd.cmd.Wait()
		if systemCmd.writer != nil {
			systemCmd.writer.Close()
		} else {
			io.Copy(os.Stdout, &buffer)
		}
	}
}

func execCommand(command *userInputCommand, pwd string) {
	cmd := exec.Command(command.inputCommand, command.options...)
	cmd.Dir = pwd
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		fmt.Printf("%s\n", err)
	}
}
