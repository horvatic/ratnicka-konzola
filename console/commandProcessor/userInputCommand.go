package commandProcessor

type userInputCommand struct {
	inputCommand string
	options      []string
}

func newUserInputCommand(command string, userOptions []string) *userInputCommand {
	return &userInputCommand{inputCommand: command, options: userOptions}
}
