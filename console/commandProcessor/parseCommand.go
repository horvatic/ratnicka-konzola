package commandProcessor

import (
	"regexp"
	"strings"
)

const Pipe = "|"

func splitCommand(userInput string) []string {
	space := regexp.MustCompile(`\s+`)
	s := space.ReplaceAllString(userInput, " ")
	return strings.Split(s, " ")
}

func buildOptions(userOptions []string) []string {
	var options []string
	catMode := false
	offset := 1
	for index, option := range userOptions {
		if strings.HasSuffix(option, "\\") {
			if catMode {
				options[index-offset] = options[index-offset] + " " + strings.ReplaceAll(option, "\\", "")
				offset = offset + 1
			} else {
				options = append(options, strings.ReplaceAll(option, "\\", ""))
			}
			catMode = true
		} else {
			if catMode {
				options[index-offset] = options[index-offset] + " " + strings.ReplaceAll(option, "\\", "")
			} else {
				options = append(options, strings.ReplaceAll(option, "\\", ""))
			}
			catMode = false
		}
	}
	return options
}

func parseNonPipeCommand(userInput string) *userInputCommand {
	command := splitCommand(userInput)
	if len(command) == 1 {
		return newUserInputCommand(command[0], nil)
	}
	return newUserInputCommand(command[0], buildOptions(command[1:]))
}

func parsePipeCommand(userInput string) []*userInputCommand {
	commands := strings.Split(userInput, Pipe)
	var parsedCommands []*userInputCommand
	for _, userCommand := range commands {
		cleanCommand := strings.Trim(userCommand, " ")
		command := strings.Split(cleanCommand, " ")
		if len(command) == 1 {
			parsedCommands = append(parsedCommands, newUserInputCommand(command[0], nil))
		} else {
			parsedCommands = append(parsedCommands, newUserInputCommand(command[0], buildOptions(command[1:])))
		}
	}
	return parsedCommands
}

func containsPipes(command string) bool {
	return strings.Contains(command, Pipe)
}
