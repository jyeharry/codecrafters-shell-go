package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
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
		code, err := strconv.Atoi(args)
		if err != nil {
			fmt.Println("not a valid exit code")
			os.Exit(0)
		}
		os.Exit(code)
	case typeCmd:
		if slices.Contains(builtins, args) {
			fmt.Printf("%s is a shell builtin\n", args)
		} else if path, found := findExecutable(args); found {
			fmt.Println(path)
		} else {
			fmt.Printf("%s not found\n", args)
		}
	default:
		fmt.Printf("%s: command not found\n", strings.Trim(input, "\n"))
	}
}

func findExecutable(cmd string) (string, bool) {
	osPath := os.Getenv("PATH")
	paths := strings.Split(osPath, ":")

	for _, path := range paths {
		fullpath := filepath.Join(path, cmd)
		if _, err := os.Stat(fullpath); err == nil {
			return fullpath, true
		}
	}

	return "", false
}
