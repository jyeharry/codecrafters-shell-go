package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	exit = "exit"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		input, err := (reader.ReadString('\n'))
		if err != nil {
			fmt.Println(err)
		}

		parsedInput := strings.Split(strings.Trim(input, "\n"), " ")
		command := parsedInput[0]
		args := parsedInput[1:]

		switch command {
		case exit:
			code, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println("not a valid exit code")
				os.Exit(0)
			}
			os.Exit(code)
		default:
			fmt.Printf("%s: command not found\n", strings.Trim(input, "\n"))
		}
	}
}
