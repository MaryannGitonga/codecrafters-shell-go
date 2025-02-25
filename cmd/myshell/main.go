package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {

	builtins := map[string]bool{
		"echo": true,
		"exit": true,
		"type": true,
		"pwd":  true,
		"cd":   true,
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

		if strings.HasPrefix(command, "pwd") {
			dir, err := os.Getwd()

			if err != nil {
				fmt.Fprintln(os.Stderr, "Error getting current directory:", err)
			} else {
				fmt.Println(dir)
			}

			continue
		}

		if strings.HasPrefix(command, "cd") {
			args := strings.Fields(command)

			if len(args) != 2 {
				fmt.Println("Invalid cd command usage. Use: cd <directory>")
				continue
			}

			path := args[1]

			if path == "~" {
				path = os.Getenv("HOME")
			}

			if err := os.Chdir(path); err != nil {
				fmt.Printf("cd: %s: No such file or directory\n", path)
			}

			continue
		}

		if strings.HasPrefix(command, "type") {
			args := strings.Fields(command)

			if len(args) == 2 {
				if builtins[args[1]] {
					fmt.Printf("%s is a shell builtin\n", args[1])
				} else {
					found := false
					pathEnv := os.Getenv("PATH")
					paths := strings.Split(pathEnv, ":")

					for _, path := range paths {
						executablePath := filepath.Join(path, args[1])

						if _, err := os.Stat(executablePath); err == nil {
							fmt.Printf("%s is %s\n", args[1], executablePath)
							found = true
							break
						}
					}

					if !found {
						fmt.Printf("%s: not found\n", args[1])
					}
				}
			}

			continue
		}

		// Execute programs
		args := strings.Fields(command)
		program := args[0]
		found := false
		pathEnv := os.Getenv("PATH")
		paths := strings.Split(pathEnv, ":")

		for _, path := range paths {
			executablePath := filepath.Join(path, program)

			if _, err := os.Stat(executablePath); err == nil {
				cmd := exec.Command(executablePath, args[1:]...)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr

				if err := cmd.Run(); err != nil {
					fmt.Printf("Error running command: %v\n", err)
				}

				found = true
				break
			}
		}

		if found {
			continue
		}

		fmt.Printf("%s: command not found\n", command)
	}
}
