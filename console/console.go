package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("% ")
		userInput, _ := reader.ReadString('\n')

		// convert CRLF to LF
		userInput = strings.Replace(userInput, "\n", "", -1)

		if userInput == "" {
			continue
		}

		if strings.ToLower(userInput) == "exit" {
			break
		}

		command := strings.Split(userInput, " ")

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
}
