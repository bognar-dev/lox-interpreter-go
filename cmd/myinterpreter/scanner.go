package main

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

	PRINT  TokenType = " "
	RETURN TokenType = " "
	SUPER  TokenType = " "
	THIS   TokenType = " "
	TRUE   TokenType = " "
	VAR    TokenType = " "
	WHILE  TokenType = " "

	EOF TokenType = " "
)

func (t TokenType) toString() string {

}

type Token struct {
	tokenType TokenType
	lexeme    string
	literal   interface{}
	line      int
}

func (t Token) toString() string {
	return t.tokenType + " " + t.lexeme + " " + t.literal
}

type Scanner struct {
	source  string
	tokens  []Token
	start   int
	current int
	line    int
}

func (s *Scanner) scanTokens() []Token {
	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}
	s.tokens = append(s.tokens, Token{EOF, "", nil, s.line})
	return s.tokens
}
