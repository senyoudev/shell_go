package shell

import (
	"fmt"
	"os/exec"
	"strings"
	"os"
)


func (s *Shell) ExecuteCommand(input string) {
parts := strings.SplitN(input, " ", 2)
	command := parts[0]
	var arguments []string
	if len(parts) > 1 {
		arguments = strings.Fields(parts[1])
	}

	if command == "type" {
		s.handleTypeCommand(arguments)
		return
	}

	if builtin, exists := s.builtins[command]; exists {
		builtin.Execute(arguments)
		return
	}

	// Execute External command
	if err := s.executeExternal(command, arguments); err != nil {
		fmt.Printf("%s: command not found\n", command)
	}
}

func (s *Shell) executeExternal(command string, args []string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}


func (s *Shell) handleTypeCommand(args []string) {
		if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "type: missing argument")
		return
	}

	cmdName := args[0]

	// Check if it's a built-in command
	if s.IsBuiltin(cmdName) {
		fmt.Printf("%s is a shell builtin\n", cmdName)
		return
	}

	// Command not found
	fmt.Printf("%s: not found\n", cmdName)
}