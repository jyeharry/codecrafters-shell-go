package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

var Commands = []string{}

func main() {
	fmt.Fprint(os.Stdout, "$ ")

	reader := bufio.NewReader(os.Stdin)
	// Wait for user input
	command, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)	
	} 

	if !slices.Contains(Commands, command) {
		fmt.Printf("%s: command not found\n", strings.Trim(command, "\n"))
	}
}
