package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func main() {

	for {
		// Wait for user input
		fmt.Fprint(os.Stdout, "$ ")
		command, err := readInput(bufio.NewReader(os.Stdin))
		if err != nil {
			printErrorAndExit("Error reading input", err)
		}
		executeCommand(command)
	}
}

func executeCommand(input string) {
	parts := strings.SplitN(input, " ", 2)
	command := parts[0]
	arguments := parts[1:]

		switch command {
		case "exit":
			exit_code, err := strconv.Atoi(arguments[1])
			if err != nil {
				fmt.Fprintln(os.Stderr, "Exit Code must be a number:", err)
				os.Exit(1)
			}
			os.Exit(exit_code)
		default:
			fmt.Println(command + ": command not found")
		}
}



func readInput(reader *bufio.Reader) (string, error) {
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		os.Exit(1)
	}

	return strings.TrimSpace(input), nil
}

func printErrorAndExit(msg string, err error) {
	fmt.Fprintln(os.Stderr, msg+":", err)
	os.Exit(1)
}