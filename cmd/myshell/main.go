package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	builtins := map[string]bool{
		"echo": true,
		"exit": true,
		"type": true,
	}

	for {
		fmt.Fprint(os.Stdout, "$ ")

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
			fmt.Println(arg)
			continue
		}

		if strings.HasPrefix(command, "type") {
			args := strings.Fields(command)

			if len(args) == 2 {
				if builtins[args[1]] {
					fmt.Printf("%s is a shell builtin\n", args[1])
				} else {
					fmt.Printf("Invalid type command usage. Use type <command>")
				}
			}

			continue
		}

		fmt.Printf("%s: command not found\n", command)
	}
}
