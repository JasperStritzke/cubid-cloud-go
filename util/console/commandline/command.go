package commandline

import "github.com/JasperStritzke/cubid-cloud/util/console/logger"

type Command struct {
	Name        string
	Description string
	SubCommands CommandList
	Exec        func(args []string)
	NotFound    func()
}

func (command *Command) Invoke(args []string) {
	if len(command.SubCommands) == 0 || len(args) <= 1 {
		if command.Exec != nil {
			command.Exec(args)
		}

		return
	}

	subCommand, exists := command.SubCommands[args[0]]

	if !exists {
		if command.NotFound != nil {
			command.NotFound()
			return
		}

		logger.Warn("Befehl nicht gefunden!")
		return
	}

	var newArgs []string
	if len(args) > 1 {
		newArgs = args[1:]
	} else {
		newArgs = nil
	}

	subCommand.Invoke(newArgs)
}
