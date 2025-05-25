package shell

import (
	"fmt"
	"os/exec"
	"strings"
	"os"
)


func (s *Shell) ExecuteCommand(input string) {
	command, arguments, err := parseCommandLine(input)
	if err != nil {
		fmt.Printf("Unable to parse %s\n", input)
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

func parseCommandLine(input string) (string, []string, error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return "", nil, fmt.Errorf("empty command")
	}

	// Extract command
	firstSpaceIdx := strings.IndexFunc(input, func(r rune) bool { return r == ' ' || r == '\t' })
	if firstSpaceIdx == -1 {
		return input, []string{}, nil
	}

	command := input[:firstSpaceIdx]
	rest := input[firstSpaceIdx:]

	args, err := parseArguments(rest)
	if err != nil {
		return "", nil, err
	}

	return command, args, nil
}

func parseArguments(input string) ([]string, error) {
	var args []string
	var currentWord strings.Builder
	i := 0

	for i < len(input) {
		char := input[i]

		// Skip whitespace between words
		if char == ' ' || char == '\t' {
			if currentWord.Len() > 0 {
				args = append(args, currentWord.String())
				currentWord.Reset()
			}
			// Skip all consecutive whitespace
			for i < len(input) && (input[i] == ' ' || input[i] == '\t') {
				i++
			}
			continue
		}

		// Handle single quotes
		if char == '\'' {
			i++ 
			// Read until next single quote
			for i < len(input) && input[i] != '\'' {
				currentWord.WriteByte(input[i])
				i++
			}
			if i >= len(input) {
				return nil, fmt.Errorf("unclosed single quote")
			}
			i++
			continue
		}

		// Handle double quotes
		if char == '"' {
			i++
			for i < len(input) && input[i] != '"' {
				if input[i] == '\\' {
					if i+1 >= len(input) {
						return nil, fmt.Errorf("unterminated escape sequence")
					}

					nextChar := input[i + 1]
					switch nextChar {
					case '\\', '"', '$':
						currentWord.WriteByte(nextChar)
						i += 2 
					default:
						currentWord.WriteByte('\\')
						i++
					}
				} else {
					currentWord.WriteByte(input[i])
					i++
				}
			}
			if i >= len(input) {
				return nil, fmt.Errorf("unclosed double quote")
			}
			i++
			continue
		}

		// Handle backslash outside quotes (escape next character)
		if char == '\\' {
			if i+1 >= len(input) {
				return nil, fmt.Errorf("unterminated escape sequence")
			}
			// Add the escaped character literally
			currentWord.WriteByte(input[i+1])
			i += 2
			continue
		}

		// Regular character
		currentWord.WriteByte(char)
		i++
	}

	// Add final word if present
	if currentWord.Len() > 0 {
		args = append(args, currentWord.String())
	}

	return args, nil
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

	// Look for executables in PATH
	path, err := exec.LookPath(cmdName)
	if err == nil {
		fmt.Printf("%s is %s\n", cmdName, path)
		return
	}

	// Command not found
	fmt.Printf("%s: not found\n", cmdName)
}