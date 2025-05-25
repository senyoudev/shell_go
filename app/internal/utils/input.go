package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadInput(reader *bufio.Reader) (string, error) {
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		os.Exit(1)
	}

	return strings.TrimSpace(input), nil
}

func PrintErrorAndExit(msg string, err error) {
	fmt.Fprintln(os.Stderr, msg+":", err)
	os.Exit(1)
}