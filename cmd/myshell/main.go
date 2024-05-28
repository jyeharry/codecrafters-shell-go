package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	echo = "echo"
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
		default:
			fmt.Printf("%s: command not found\n", strings.Trim(input, "\n"))
		}
	}
}
