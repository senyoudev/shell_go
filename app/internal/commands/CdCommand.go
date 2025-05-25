package commands

import (
	"fmt"
	"os"
	"strings"
	"path"
)


type CdCommand struct{}
func (c CdCommand) Name() string { return "cd" }
func (c CdCommand) Execute(args []string) error {
	path_arg := args[0]
	if strings.HasPrefix(path_arg, "~") {
		HOME := os.Getenv("HOME")
		if (len(HOME)) == 0 {
			fmt.Println("cd: $HOME is not set.")
		} else {
			path_arg = path.Join(HOME, path_arg[1:])
		}
	}
	err := os.Chdir(path_arg)
	if err != nil {
		fmt.Printf("%s: %s: No such file or directory\n", c.Name(), path_arg)
	}
	return nil
}