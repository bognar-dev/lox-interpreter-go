package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Scanner struct {
	source    string
	tokens    []Token
	start     int
	current   int
	line      int
	errorList []error
}

func (s *Scanner) scanTokens() ([]Token, []error) {
	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}
	s.tokens = append(s.tokens, Token{EOF, "", Literal{}, s.line})
	return s.tokens, s.errorList
}

func (s *Scanner) peek() TokenType {
	return TokenType(s.source[s.current-1 : s.current])
}

func (s *Scanner) peekString() string {
	return s.source[s.current-1 : s.current]
}

func (s *Scanner) addToken(token TokenType, literal Literal) {
	s.tokens = append(s.tokens, Token{token, tokenLoopUp[token], literal, s.line})
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) peekAt(pos int) TokenType {
	return TokenType(s.source[pos])
}
func (s *Scanner) advance() TokenType {
	s.current++

	return TokenType(s.source[s.current-1])
}

func (s *Scanner) createString() {
	s.advance()
	for s.peek() != PARENTHESES && !s.isAtEnd() {
		if s.peek() == NEWLINE {
			s.line++
		}

		s.advance()
	}
	str := s.source[s.start:s.current]
	if s.isAtEnd() && str[len(str)-1] != '"' {
		s.errorList = append(s.errorList, fmt.Errorf("[line %d] Error: Unterminated string.", s.line))
		return
	}
	s.addToken(STRING, Literal{STRING_LITERAL, str})
}
func (s *Scanner) createNumber() {
	for s.isDigit(s.peek()) && !s.isAtEnd() {
		s.advance()
	}
	var next TokenType
	// Look for a fractional part.
	if s.current >= len(s.source) {
		next = s.peek()
	} else {
		next = s.peekNext()
	}

	curr := s.peek()

	if curr == DOT && s.isDigit(next) {
		// Consume the "."
		s.advance()

		for s.isDigit(s.peek()) && !s.isAtEnd() {
			s.advance()
		}
	}
	str := strings.Trim(s.source[s.start:s.current], "\r")

	trimmedStr := strings.TrimSuffix(str, ".")
	floatVal, err := strconv.ParseFloat(trimmedStr, 64)
	if err != nil {
		s.errorList = append(s.errorList, fmt.Errorf("[line %d] Error: Invalid number.", s.line))
		return
	}
	s.addToken(NUMBER, Literal{NUMBER_LITERAL, floatVal})
}

func (s *Scanner) printTokens(tokens []Token) {
	for _, token := range tokens {
		fmt.Println(token.String())
	}
}

func (s *Scanner) isDigit(c TokenType) bool {
	return c >= "0" && c <= "9"
}

func (s *Scanner) match(expected TokenType) bool {
	if s.isAtEnd() {
		return false
	}
	peekVal := s.peekAt(s.current)
	if peekVal != expected {
		return false
	}

	s.current++
	return true
}

func (s *Scanner) peekNext() TokenType {
	return TokenType(s.source[s.current])
}

func (s *Scanner) scanToken() {
	c := s.advance()
	switch c {
	case LEFT_PAREN:
		s.addToken(LEFT_PAREN, Literal{})
	case RIGHT_PAREN:
		s.addToken(RIGHT_PAREN, Literal{})
	case LEFT_BRACE:
		s.addToken(LEFT_BRACE, Literal{})
	case RIGHT_BRACE:
		s.addToken(RIGHT_BRACE, Literal{})
	case STAR:
		s.addToken(STAR, Literal{})
	case DOT:
		s.addToken(DOT, Literal{})
	case COMMA:
		s.addToken(COMMA, Literal{})
	case PLUS:
		s.addToken(PLUS, Literal{})
	case MINUS:
		s.addToken(MINUS, Literal{})
	case SEMICOLON:
		s.addToken(SEMICOLON, Literal{})
	case SLASH:
		if s.match(SLASH) {
			for s.peek() != NEWLINE && !s.isAtEnd() {
				s.advance()
			}
			s.line++
		} else {
			s.addToken(SLASH, Literal{})
		}

	case EOF:
		s.addToken(EOF, Literal{})
	case EQUAL:
		if s.match(EQUAL) {
			s.addToken(EQUAL_EQUAL, Literal{})
		} else {
			s.addToken(EQUAL, Literal{})
		}
	case BANG:
		if s.match(EQUAL) {
			s.addToken(BANG_EQUAL, Literal{})
		} else {
			s.addToken(BANG, Literal{})
		}
	case GREATER:
		if s.match(EQUAL) {
			s.addToken(GREATER_EQUAL, Literal{})
		} else {
			s.addToken(GREATER, Literal{})
		}
	case LESS:
		if s.match(EQUAL) {
			s.addToken(LESS_EQUAL, Literal{})
		} else {
			s.addToken(LESS, Literal{})
		}
	case CARRIAGE_RETURN, WHITESPACE, TABULATOR:

	case NEWLINE:
		s.line++

	case PARENTHESES:
		s.createString()
	default:
		if s.isDigit(c) {
			s.createNumber()
			return
		}
		s.errorList = append(s.errorList, fmt.Errorf("[line %d] Error: Unexpected character: %s", s.line, s.peekString()))

	}
}
