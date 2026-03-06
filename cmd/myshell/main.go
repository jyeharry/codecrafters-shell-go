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
	pwd     = "pwd"
	cd      = "cd"
)

var builtins = []string{echo, exit, typeCmd, pwd, cd}

func main() {
	for {
		fmt.Print("$ ")

		// Wait for user input
		input, err := (bufio.NewReader(os.Stdin).ReadString('\n'))
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
	case pwd:
		pwd, _ := os.Getwd()
		fmt.Println(pwd)
	case cd:
		if err := os.Chdir(args); err != nil {
			fmt.Printf("cd: %s: No such file or directory\n", args)
		}
	default:
		runExternalCommand(command, input, args)
	}
}

func runExternalCommand(command string, input string, args string) {
	argsSlice := strings.Split(args, " ")
	_, err := exec.LookPath(command)
	if err != nil {
		fmt.Printf("%s: command not found\n", strings.Trim(input, "\n"))
		return
	}
	cmd := exec.Command(command, argsSlice...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}
