package commandline

import "strings"

type CommandList map[string]Command

type CommandLine struct {
	Commands CommandList
}

func (commandLine *CommandLine) RegisterCommand(command Command) *CommandLine {
	commandLine.Commands[strings.ToLower(command.Name)] = command
	return commandLine
}

func (commandLine *CommandLine) Invoke(cmd string, args []string) {
	foundCommand, exists := commandLine.Commands[strings.ToLower(cmd)]

	if !exists {
		commandLine.Invoke("help", args)
		return
	}

	foundCommand.Invoke(args)
}
