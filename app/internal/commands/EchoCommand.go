package commands

import (
	"fmt"
	"strings"
)


type EchoCommand struct{}
func (c EchoCommand) Name() string { return "echo" }
func (c EchoCommand) Execute(args []string) error {
	output := strings.Join(args, " ")
	fmt.Println(output)
	return nil
}