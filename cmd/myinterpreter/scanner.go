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
	IDENTIFIER TokenType = "identifier"
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

	EOF     TokenType = "null"
	UNKNOWN TokenType = "unknown"
)

func (t TokenType) toString() string {
	switch t {
	case LEFT_PAREN:
		return "("
	case RIGHT_PAREN:
		return ")"
	case LEFT_BRACE:
		return "{"
	case RIGHT_BRACE:
		return "}"
	case STAR:
		return "*"
	case DOT:
		return "."
	case COMMA:
		return ","
	case PLUS:
		return "+"
	case MINUS:
		return "-"
	case SEMICOLON:
		return ";"
	case SLASH:
		return "/"
	case BANG:
		return "!"
	case BANG_EQUAL:
		return "!="
	case EQUAL:
		return "="
	case EQUAL_EQUAL:
		return "=="
	case GREATER:
		return ">"
	case GREATER_EQUAL:
		return ">="
	case LESS:
		return "<"
	case LESS_EQUAL:
		return "<="
	case IDENTIFIER:
		return "IDENTIFIER"
	case STRING:
		return "STRING"
	case NUMBER:
		return "NUMBER"
	case AND:
		return "AND"
	case CLASS:
		return "CLASS"
	case ELSE:
		return "ELSE"
	case FALSE:
		return "FALSE"
	case FUN:
		return "FUN"
	case FOR:
		return "FOR"
	case IF:
		return "IF"
	case NIL:
		return "NIL"
	case OR:
		return "OR"
	case PRINT:
		return "PRINT"
	case RETURN:
		return "RETURN"
	case SUPER:
		return "SUPER"
	case THIS:
		return "THIS"
	case TRUE:
		return "TRUE"
	case VAR:
		return "VAR"
	case WHILE:
		return "WHILE"
	case EOF:
		return "EOF"
	}
	return string(t)
}

type Token struct {
	tokenType TokenType
	lexeme    string
	literal   interface{}
	line      int
}

func (t Token) toString() string {
	return t.tokenType.toString() + "  " + t.lexeme
}

type Scanner struct {
	source   string
	tokens   []Token
	start    int
	current  int
	line     int
	errorMsg error
}

func (s *Scanner) scanTokens() ([]Token, error) {
	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}
	s.tokens = append(s.tokens, Token{EOF, "null", nil, s.line})
	return s.tokens, s.errorMsg
}

func (s *Scanner) peek() TokenType {
	return TokenType(s.source[s.current-1 : s.current])
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

func (s *Scanner) printTokens(tokens []Token) {
	for _, token := range tokens {
		fmt.Println(token.toString())
	}
}

func (s *Scanner) scanToken() {
	c := s.advance()
	if c != UNKNOWN {
		s.addToken(c)
		return
	}

	if s.errorMsg == nil {
		s.errorMsg = fmt.Errorf("[line " + strconv.Itoa(s.line) + "] Error: Unexpected character: " + s.peek().toString())
	} else {
		s.errorMsg = fmt.Errorf("%w \n %s", s.errorMsg, "[line "+strconv.Itoa(s.line)+"] Error: Unexpected character: "+s.peek().toString())
	}

}
