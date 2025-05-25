package commands

import (
	"fmt"
)

type TypeCommand struct{}
func(c TypeCommand) Name() string{
	return "type"
}

func(c TypeCommand) Execute(args[] string) error {
	return fmt.Errorf("type_command_needs_special_handling")
}