package main

import (
	"fmt"
	"os"
)

type TokenType rune

const (
	LEFT_PAREN  TokenType = '('
	RIGHT_PAREN TokenType = ')'
	LEFT_BRACE  TokenType = '{'
	RIGHT_BRACE TokenType = '}'

	COMMA     TokenType = ','
	DOT       TokenType = '.'
	MINUS     TokenType = '-'
	PLUS      TokenType = '+'
	SEMICOLON TokenType = ';'
	SLASH     TokenType = '/'
	STAR      TokenType = '*'

	// One or two character tokens.
	BANG       TokenType = '!'
	BANG_EQUAL TokenType = ' '

	EQUAL       TokenType = '='
	EQUAL_EQUAL TokenType = ' '

	GREATER       TokenType = '>'
	GREATER_EQUAL TokenType = ' '

	LESS       TokenType = '<'
	LESS_EQUAL TokenType = ' '

	// Literals.
	IDENTIFIER TokenType = '.'
	STRING     TokenType = ' '
	NUMBER     TokenType = ' '

	// Keywords.
	AND   TokenType = ' '
	CLASS TokenType = ' '
	ELSE  TokenType = ' '
	FALSE TokenType = ' '
	FUN   TokenType = ' '
	FOR   TokenType = ' '
	IF    TokenType = ' '
	NIL   TokenType = ' '
	OR    TokenType = ' '

	PRINT  TokenType = ' '
	RETURN TokenType = ' '
	SUPER  TokenType = ' '
	THIS   TokenType = ' '
	TRUE   TokenType = ' '
	VAR    TokenType = ' '
	WHILE  TokenType = ' '

	EOF TokenType = ' '
)

type Token struct {
	tokenType TokenType
	lexme     string
}
type Scanner struct {
	source string
	tokens []string
}

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
	fileContents := string(rawfileContents)
	for _, current := range fileContents {
		token := TokenType(current)
		switch token {
		case LEFT_PAREN:
			fmt.Println("LEFT_PAREN ( null")
		case RIGHT_PAREN:
			fmt.Println("RIGHT_PAREN ) null")
		case LEFT_BRACE:
			fmt.Println("LEFT_BRACE { null")
		case RIGHT_BRACE:
			fmt.Println("RIGHT_BRACE } null")
		case STAR:
			fmt.Println("STAR * null")
		case DOT:
			fmt.Println("DOT . null")
		case COMMA:
			fmt.Println("COMMA , null")
		case PLUS:
			fmt.Println("PLUS + null")
		case MINUS:
			fmt.Println("MINUS - null")
		case SEMICOLON:
			fmt.Println("SEMICOLON ; null")
		case SLASH:
			fmt.Println("SLASH / null")
		}

	}
	fmt.Println("EOF  null")
}
