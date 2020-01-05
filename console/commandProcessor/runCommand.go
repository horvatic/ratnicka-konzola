package commandProcessor

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
)

func execPipeCommands(commands []*userInputCommand, pwd string) {
	var cmds []*systemCommand
	var currentReadPipe *io.PipeReader
	var finalOutBuffer bytes.Buffer
	for index, command := range commands {
		errBuffer := new(bytes.Buffer)
		cmd := exec.Command(command.inputCommand, command.options...)
		cmd.Dir = pwd
		cmd.Stderr = errBuffer
		systemCmd := newSystemCommand(cmd)
		systemCmd.errorWriter = errBuffer
		if index > 0 {
			systemCmd.cmd.Stdin = currentReadPipe
		} else {
			systemCmd.cmd.Stdin = os.Stdin
		}
		r, w := io.Pipe()
		if index == len(commands)-1 {
			systemCmd.cmd.Stdout = &finalOutBuffer
			systemCmd.writer = nil
		} else {
			systemCmd.cmd.Stdout = w
			systemCmd.writer = w
		}
		currentReadPipe = r
		cmds = append(cmds, systemCmd)
	}
	runCommandList(cmds, &finalOutBuffer)
}

func runCommandList(cmds []*systemCommand, finalOutBuffer *bytes.Buffer) {
	for _, systemCmd := range cmds {
		err := systemCmd.cmd.Start()
		if err != nil {
			fmt.Printf("%s\n", err)
		}
	}

	for _, systemCmd := range cmds {
		systemCmd.cmd.Wait()
		io.Copy(os.Stdout, systemCmd.errorWriter)
		if systemCmd.writer != nil {
			systemCmd.writer.Close()
		} else {
			io.Copy(os.Stdout, finalOutBuffer)
		}
	}
}

func execStreamCommand(command *userInputCommand, pwd string) {
	resp, err := http.Post("http://127.0.0.1:8080/", "text/plain", bytes.NewBuffer([]byte(command.inputCommand)))
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func execCommand(command *userInputCommand, pwd string) {
	var errBuffer bytes.Buffer
	cmd := exec.Command(command.inputCommand, command.options...)
	cmd.Dir = pwd
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = &errBuffer
	err := cmd.Run()
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
	io.Copy(os.Stdout, &errBuffer)
}
