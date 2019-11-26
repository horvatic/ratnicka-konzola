package services

import (
	"fmt"
	"os"
	"os/exec"
)

func execCommand(command []string) {
	if len(command) == 1 {
		cmd := exec.Command(command[0])
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			fmt.Printf("%s\n", err)
		}
	} else {
		options := command[1:len(command)]
		cmd := exec.Command(command[0], options...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			fmt.Printf("%s\n", err)
		}
	}
}
