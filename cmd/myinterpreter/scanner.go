package main

import (
	"fmt"
	"strconv"
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
	s.tokens = append(s.tokens, Token{EOF, "", nil, s.line})
	return s.tokens, s.errorList
}

func (s *Scanner) peek() TokenType {
	return TokenType(s.source[s.current-1 : s.current])
}

func (s *Scanner) peekString() string {
	return s.source[s.current-1 : s.current]
}

func (s *Scanner) addToken(token TokenType) {
	s.tokens = append(s.tokens, Token{token, tokenLoopUp[token], nil, s.line})
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

func (s *Scanner) printTokens(tokens []Token) {
	for _, token := range tokens {
		fmt.Println(token.String() + " null")
	}
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

func (s *Scanner) scanToken() {
	c := s.advance()
	switch c {
	case LEFT_PAREN:
		s.addToken(LEFT_PAREN)
	case RIGHT_PAREN:
		s.addToken(RIGHT_PAREN)
	case LEFT_BRACE:
		s.addToken(LEFT_BRACE)
	case RIGHT_BRACE:
		s.addToken(RIGHT_BRACE)
	case STAR:
		s.addToken(STAR)
	case DOT:
		s.addToken(DOT)
	case COMMA:
		s.addToken(COMMA)
	case PLUS:
		s.addToken(PLUS)
	case MINUS:
		s.addToken(MINUS)
	case SEMICOLON:
		s.addToken(SEMICOLON)
	case SLASH:
		s.addToken(SLASH)
	case EOF:
		s.addToken(EOF)
	case EQUAL:
		if s.match(EQUAL) {
			s.addToken(EQUAL_EQUAL)
		} else {
			s.addToken(EQUAL)
		}

	default:
		s.errorList = append(s.errorList, fmt.Errorf("[line %s] Error: Unexpected character: %s", strconv.Itoa(s.line), s.peekString()))

	}
}
