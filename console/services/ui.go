package services

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Ui() {
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

		command := strings.Split(userInput, " ")
		execCommand(command)
	}

}
