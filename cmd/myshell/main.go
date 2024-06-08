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
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
		}

		command := strings.TrimSpace(input)

		if strings.HasPrefix(command, "exit") {
			args := strings.Fields(command)

			if len(args) == 2 {
				if exitCode, err := strconv.Atoi(args[1]); err == nil {
					os.Exit(exitCode)
				}
			}

			fmt.Println("Invalid exit command usage. Use: exit <code>")
			continue
		}

		if strings.HasPrefix(command, "echo") {
			arg := strings.TrimSpace(strings.TrimPrefix(command, "echo"))
			fmt.Print(arg)
			continue
		}

		fmt.Printf("%s: command not found\n", command)
	}
}
