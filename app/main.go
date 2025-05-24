package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {

	for {
		// Wait for user input
		fmt.Fprint(os.Stdout, "$ ")
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}

		command = strings.TrimSpace(command)
		parts := strings.SplitN(command, " ", 2) // command, args

		switch parts[0] {
		case "exit":
			exit_code, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Fprintln(os.Stderr, "Exit Code must be a number:", err)
				os.Exit(1)
			}
			os.Exit(exit_code)
		default:
			fmt.Println(command + ": command not found")
		}
	}
}
