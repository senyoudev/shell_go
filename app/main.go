package main

import (
	"bufio"
	"fmt"
	"os"
	"github.com/codecrafters-io/shell-starter-go/app/internal/shell"
	"github.com/codecrafters-io/shell-starter-go/app/internal/utils"
)


func main() {

	shell := shell.NewShell()

	for {
		// Wait for user input
		fmt.Fprint(os.Stdout, "$ ")
		command, err := utils.ReadInput(bufio.NewReader(os.Stdin))
		if err != nil {
			utils.PrintErrorAndExit("Error reading input", err)
		}
		shell.ExecuteCommand(command)
	}
}



