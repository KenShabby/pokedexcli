package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		command := scanner.Text()
		command_slice := cleanInput(command)
		fmt.Println("Your command was:", command_slice[0])
	}
}

func cleanInput(text string) []string {

	lower := strings.ToLower(text)
	words := strings.Fields(lower)

	return words
}
