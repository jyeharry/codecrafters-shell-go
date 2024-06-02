package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strconv"
	"strings"
)

const (
	echo    = "echo"
	exit    = "exit"
	typeCmd = "type"
)

var builtins = []string{echo, exit, typeCmd}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		input, err := (reader.ReadString('\n'))
		if err != nil {
			fmt.Println(err)
		}

		handleCommand(input)
	}
}

func handleCommand(input string) {
	command, args, _ := strings.Cut(strings.Trim(input, "\n"), " ")
	switch command {
	case echo:
		fmt.Println(args)
	case exit:
		if code, err := strconv.Atoi(args); err != nil {
			if args != "" {
				fmt.Println("not a valid exit code")
			}
			os.Exit(0)
		} else {
			os.Exit(code)		
		}
	case typeCmd:
		if slices.Contains(builtins, args) {
			fmt.Printf("%s is a shell builtin\n", args)
		} else if path, err := exec.LookPath(args); err == nil {
			fmt.Println(path)
		} else {
			fmt.Printf("%s not found\n", args)
		}
	default:
		runExternalCommand(command, input, args)
	}
}

func runExternalCommand(command string, input string, args string) {
	path, err := exec.LookPath(command)
	if err != nil {
		fmt.Printf("%s: command not found\n", strings.Trim(input, "\n"))
		return
	}
	cmd := exec.Command(path, args)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}
