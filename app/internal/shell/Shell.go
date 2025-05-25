package shell

import (
	"github.com/codecrafters-io/shell-starter-go/app/internal/commands"
)

type Shell struct{
	builtins map[string]commands.Command
}

func NewShell() *Shell{
	shell := &Shell{
		builtins: make(map[string]commands.Command),
	}

	// Register commands
	shell.registerBuiltin(commands.ExitCommand{})
	shell.registerBuiltin(commands.EchoCommand{})
	shell.registerBuiltin(commands.TypeCommand{})

	return shell
}

func (s *Shell) registerBuiltin(cmd commands.Command) {
	s.builtins[cmd.Name()] = cmd
}

func (s *Shell) IsBuiltin(name string) bool {
	_, exists := s.builtins[name]
	return exists
}


