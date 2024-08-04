package main

import (
	"fmt"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they"l be visible when running tests.
	//fmt.Fprintln(os.Stderr, "Logs from your program will appear here!")

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}

	command := os.Args[1]

	if command != "tokenize" {
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}

	filename := os.Args[2]
	rawfileContents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}
	source := string(rawfileContents)
	scanner := Scanner{source: source, tokens: []Token{}, start: 0, current: 0, line: 1}
	var errorList []error
	var tokens []Token
	tokens, errorList = scanner.scanTokens()
	scanner.printTokens(tokens)
	if len(errorList) != 0 {
		for _, err := range errorList {
			fmt.Fprintln(os.Stderr, err)
		}
		os.Exit(65)
	}

}
