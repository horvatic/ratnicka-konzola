package ui

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/horvatic/ratnicka-konzola/console/commandProcessor"
	"github.com/horvatic/ratnicka-konzola/console/customCommands"
	"github.com/horvatic/ratnicka-konzola/console/path"
)

func Start() {

	pwd := path.InitPath()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s& ", pwd)
		userInput, _ := reader.ReadString('\n')

		userInput = strings.Replace(userInput, "\n", "", -1)

		if userInput == "" {
			continue
		}

		if strings.HasPrefix(userInput, "cd") {
			pwd = customCommands.Cd(userInput, pwd)
			continue
		}

		if strings.ToLower(userInput) == "exit" {
			break
		}

		commandProcessor.ExecCommands(userInput, pwd)
	}

}
