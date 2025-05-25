package commands

import (
	"fmt"
	"os"
)


type CdCommand struct{}
func (c CdCommand) Name() string { return "cd" }
func (c CdCommand) Execute(args []string) error {
	path := args[0]
	err := os.Chdir(path)
	if err != nil {
		fmt.Printf("%s: %s: No such file or directory\n", c.Name(), path)
	}
	return nil
}