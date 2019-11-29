package ui

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/horvatic/ratnicka-konzola/console/commandProcessor"
)

func ReadInput() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("% ")
		userInput, _ := reader.ReadString('\n')

		userInput = strings.Replace(userInput, "\n", "", -1)

		if userInput == "" {
			continue
		}

		if strings.ToLower(userInput) == "exit" {
			break
		}

		commandProcessor.ExecCommands(userInput)
	}

}
