package main

import (
	"fmt"
	"strconv"
)

type TokenType string

const (
	LEFT_PAREN  TokenType = "("
	RIGHT_PAREN TokenType = ")"
	LEFT_BRACE  TokenType = "{"
	RIGHT_BRACE TokenType = "}"

	COMMA     TokenType = ","
	DOT       TokenType = "."
	MINUS     TokenType = "-"
	PLUS      TokenType = "+"
	SEMICOLON TokenType = ";"
	SLASH     TokenType = "/"
	STAR      TokenType = "*"

	// One or two character tokens.
	BANG       TokenType = "!"
	BANG_EQUAL TokenType = "!="

	EQUAL       TokenType = "="
	EQUAL_EQUAL TokenType = "=="

	GREATER       TokenType = ">"
	GREATER_EQUAL TokenType = ">="

	LESS       TokenType = "<"
	LESS_EQUAL TokenType = "<="

	// Literals.
	IDENTIFIER TokenType = "."
	STRING     TokenType = "string"
	NUMBER     TokenType = "number"

	// Keywords.
	AND   TokenType = "and"
	CLASS TokenType = "class"
	ELSE  TokenType = "else"
	FALSE TokenType = "false"
	FUN   TokenType = "fun"
	FOR   TokenType = "for"
	IF    TokenType = "if"
	NIL   TokenType = "nil"
	OR    TokenType = "or"

	PRINT  TokenType = "print"
	RETURN TokenType = "return"
	SUPER  TokenType = "super"
	THIS   TokenType = "this"
	TRUE   TokenType = "true"
	VAR    TokenType = "var"
	WHILE  TokenType = "while"

	EOF TokenType = "null"
)

func (t TokenType) toString() string {
	return string(t)[1:]
}

type Token struct {
	tokenType TokenType
	lexeme    string
	literal   interface{}
	line      int
}

func (t Token) toString() string {
	return t.tokenType.toString() + " " + t.lexeme + " " + fmt.Sprint(t.literal)
}

type Scanner struct {
	source   string
	tokens   []Token
	start    int
	current  int
	line     int
	hasError bool
}

func (s *Scanner) scanTokens() []Token {
	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}
	s.tokens = append(s.tokens, Token{EOF, "", nil, s.line})
	return s.tokens
}

func (s *Scanner) peek() TokenType {
	return TokenType(strconv.Itoa(int(s.source[s.current])))
}

func (s *Scanner) addToken(token TokenType) {
	s.tokens = append(s.tokens, Token{token, "", nil, s.line})
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}
func (s *Scanner) advance() TokenType {
	s.current++
	return TokenType(strconv.Itoa(int(s.source[s.current-1])))
}
func (s *Scanner) scanToken() {
	c := s.advance()
	switch c {
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
	default:
		fmt.Println("[line 1] Error: Unexpected character: " + c.toString())
		hasError = true
	}
}
