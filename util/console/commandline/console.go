package commandline

import (
	"bufio"
	"fmt"
	"github.com/JasperStritzke/cubid-cloud/util/console/logger"
	"os"
	"strings"
)

func NewCommandLine() *CommandLine {
	var cmdLine *CommandLine
	helpCommand := Command{
		SubCommands: nil,
		Exec: func(args []string) {
			logger.Info("Befehle (" + fmt.Sprint(len(cmdLine.Commands)) + "):")
		},
		NotFound: nil,
	}

	cmdLine = &CommandLine{
		Commands: CommandList{"help": helpCommand},
	}

	return cmdLine
}

func (commandLine *CommandLine) Start() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		text := scanner.Text()

		textSplitBySpaces := strings.Split(text, " ")

		firstCommand := textSplitBySpaces[0]

		var args []string
		if len(textSplitBySpaces) == 1 {
			args = []string{}
		} else {
			args = textSplitBySpaces[1:]
		}

		commandLine.Invoke(firstCommand, args)
	}
}
