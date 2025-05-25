package commands

import (
	"fmt"
	"os"
	"strconv"
)

type ExitCommand struct{}
func(c ExitCommand) Name() string{
	return "exit"
}

func(c ExitCommand) Execute(args []string) error {
	exit_code, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Exit Code must be a number:", err)
		os.Exit(1)
	}
	os.Exit(exit_code)
	return nil
}