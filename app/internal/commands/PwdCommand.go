package commands

import (
	"fmt"
	"os"
)


type PwdCommand struct{}
func (c PwdCommand) Name() string { return "pwd" }
func (c PwdCommand) Execute(args []string) error {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Errorf("Something Went wrong while executing %s", c.Name())
	}
	fmt.Printf("%s\n", dir)
	return nil
}