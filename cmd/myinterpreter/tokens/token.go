package tokens

import (
	"fmt"
	"strings"
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

	// One or two character Tokens.
	BANG       TokenType = "!"
	BANG_EQUAL TokenType = "!="

	EQUAL       TokenType = "="
	EQUAL_EQUAL TokenType = "=="

	GREATER       TokenType = ">"
	GREATER_EQUAL TokenType = ">="

	LESS       TokenType = "<"
	LESS_EQUAL TokenType = "<="

	// Literals.
	IDENTIFIER TokenType = "IDENTIFIER"
	STRING     TokenType = "STRING"
	NUMBER     TokenType = "NUMBER"

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

var TokenLoopUp = map[TokenType]string{
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

	// One or two character Tokens.
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
	NONE               LiteralType = 0
	STRING_LITERAL                 = 1
	NUMBER_LITERAL                 = 2
	IDENTIFIER_LITERAL             = 3
)

type Literal struct {
	LiteralType LiteralType
	Value       interface{}
}
type Token struct {
	TokenType TokenType
	Lexeme    string
	Literal   Literal
	Line      int
}

func (l Literal) String() string {
	switch l.LiteralType {
	case NUMBER_LITERAL:
		if l.Value == float64(int(l.Value.(float64))) {
			return fmt.Sprintf("%.1f", l.Value) // Ensures 1234.0 for whole numbers
		} else {
			return fmt.Sprintf("%g", l.Value) // Keeps the precision for non-whole numbers
		}
	default:
		return fmt.Sprintf("\"%v\"", l.Value)
	}
}

func (t Token) String() string {
	if t.Literal.LiteralType == STRING_LITERAL {
		return fmt.Sprintf("%s %s %s", t.TokenType.String(), t.Literal.String(), strings.Trim(t.Literal.String(), "\""))
	}
	if t.Literal.LiteralType == NUMBER_LITERAL {
		return fmt.Sprintf("%s %s %s", t.TokenType.String(), t.Lexeme, t.Literal.String())
	}
	if t.Literal.LiteralType == IDENTIFIER_LITERAL {
		return fmt.Sprintf("%s %s null", t.TokenType.String(), t.Literal.Value)
	}
	return t.TokenType.String() + " " + t.Lexeme + " null"
}
