package main

import (
	"fmt"
	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/parsing"
	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/scanning"
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

	switch command {
	case "tokenize":
		tokenize()
	case "evaluate":
		evaluate()
	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}

}

const (
	errExitCode        = 1
	lexicalErrExitCode = 65
)

func printErrorAndExit(err string, args ...any) {
	_, _ = fmt.Fprintf(os.Stderr, err, args...)
	os.Exit(errExitCode)
}

func openFile(path string) *os.File {
	file, err := os.Open(path)

	if err != nil {
		printErrorAndExit("Error opening file: %v\n", err)
	}

	return file
}

func evaluate() {
	filename := os.Args[2]
	file := openFile(filename)
	defer file.Close()

	scanner := parsing.NewScanner(file)
	scanner.Scan()
}

func tokenize() {
	filename := os.Args[2]
	rawfileContents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}
	source := string(rawfileContents)
	scanner := scanning.Scanner{Source: source, Tokens: []scanning.Token{}, Start: 0, Current: 0, Line: 1}
	var errorList []error
	var tokens []scanning.Token
	tokens, errorList = scanner.ScanTokens()
	scanner.PrintTokens(tokens)
	if len(errorList) != 0 {
		os.Exit(lexicalErrExitCode)
	}
}
