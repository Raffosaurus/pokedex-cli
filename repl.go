package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)

	strSlice := strings.Fields(lowered)

	return strSlice
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		if err := scanner.Err(); err != nil {
			fmt.Printf("Invalid input: %s\n", err)
		}

		line := scanner.Text()
		input := cleanInput(line)
		if len(input) == 0 {
			continue
		}

		fmt.Printf("Your command was: %v\n", input[0])
	}
}
