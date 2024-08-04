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

	EOF TokenType = ""

	NEWLINE         TokenType = "\n"
	CARRIAGE_RETURN           = "\r"
	TABULATOR                 = "\t"
	WHITESPACE                = " "
	PARENTHESES               = "\""
)

func (t TokenType) String() string {
	switch t {
	case LEFT_PAREN:
		return "LEFT_PAREN"
	case RIGHT_PAREN:
		return "RIGHT_PAREN"
	case LEFT_BRACE:
		return "LEFT_BRACE"
	case RIGHT_BRACE:
		return "RIGHT_BRACE"
	case STAR:
		return "STAR"
	case DOT:
		return "DOT"
	case COMMA:
		return "COMMA"
	case PLUS:
		return "PLUS"
	case MINUS:
		return "MINUS"
	case SEMICOLON:
		return "SEMICOLON"
	case SLASH:
		return "SLASH"
	case BANG:
		return "BANG"
	case BANG_EQUAL:
		return "BANG_EQUAL"
	case EQUAL:
		return "EQUAL"
	case EQUAL_EQUAL:
		return "EQUAL_EQUAL"
	case GREATER:
		return "GREATER"
	case GREATER_EQUAL:
		return "GREATER_EQUAL"
	case LESS:
		return "LESS"
	case LESS_EQUAL:
		return "LESS_EQUAL"
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
	case NEWLINE:
		return "\n"
	}
	return string(t)
}

var tokenLoopUp = map[TokenType]string{
	LEFT_PAREN:  "(",
	RIGHT_PAREN: ")",
	LEFT_BRACE:  "{",
	RIGHT_BRACE: "}",

	COMMA:     ",",
	DOT:       ".",
	MINUS:     "-",
	PLUS:      "+",
	SEMICOLON: ";",
	SLASH:     "/",
	STAR:      "*",

	// One or two character tokens.
	BANG:       "!",
	BANG_EQUAL: "!=",

	EQUAL:       "=",
	EQUAL_EQUAL: "==",

	GREATER:       ">",
	GREATER_EQUAL: ">=",

	LESS:       "<",
	LESS_EQUAL: "<=",

	// Literals.
	IDENTIFIER: "identifier",
	STRING:     "string",
	NUMBER:     "number",

	// Keywords.
	AND:   "and",
	CLASS: "class",
	ELSE:  "else",
	FALSE: "false",
	FUN:   "fun",
	FOR:   "for",
	IF:    "if",
	NIL:   "nil",
	OR:    "or",

	PRINT:  "print",
	RETURN: "return",
	SUPER:  "super",
	THIS:   "this",
	TRUE:   "true",
	VAR:    "var",
	WHILE:  "while",

	EOF: "null",

	NEWLINE: "\\n"}

type LiteralType int

const (
	STRING_LITERAL LiteralType = iota
	NUMBER_LITERAL
)

type Literal struct {
	literalType LiteralType
	value       string
}
type Token struct {
	tokenType TokenType
	lexeme    string
	literal   Literal
	line      int
}

func (t Token) String() string {
	return t.tokenType.String() + " " + t.lexeme
}
