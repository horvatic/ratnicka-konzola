package commandProcessor

func execNonPipe(userInput string) {
	command := parseNonPipeCommand(userInput)
	execCommand(command)
}

func execPipe(userInput string) {
	commands := parsePipeCommand(userInput)
	execPipeCommands(commands)
}

func ExecCommands(userInput string) {
	if containsPipes(userInput) {
		execPipe(userInput)
	} else {
		execNonPipe(userInput)
	}
}
