package commandProcessor

func execNonPipe(userInput string, pwd string) {
	command := parseNonPipeCommand(userInput)
	execCommand(command, pwd)
}

func execPipe(userInput string, pwd string) {
	commands := parsePipeCommand(userInput)
	execPipeCommands(commands, pwd)
}

func execStreams(userInput string, pwd string) {
	commands := parseStreamCommand(userInput)
	execStreamCommand(commands, pwd)
}

func ExecCommands(userInput string, pwd string) {
	if containsPipes(userInput) {
		execPipe(userInput, pwd)
	} else if containsStreams(userInput) {
		execStreams(userInput, pwd)
	} else {
		execNonPipe(userInput, pwd)
	}
}
