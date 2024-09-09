package scanning

import (
	"fmt"
	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/tokens"
	"os"
	"strconv"
)

type Scanner struct {
	Source    string
	Tokens    []tokens.Token
	Start     int
	Current   int
	Line      int
	ErrorList []error
}

func (s *Scanner) ScanTokens() ([]tokens.Token, []error) {
	for !s.isAtEnd() {
		s.Start = s.Current
		s.scanToken()
	}
	s.Tokens = append(s.Tokens, tokens.Token{tokens.EOF, "", tokens.Literal{}, s.Line})
	return s.Tokens, s.ErrorList
}

func (s *Scanner) peek() tokens.TokenType {
	if s.isAtEnd() {
		return tokens.EOF
	}
	return tokens.TokenType(s.Source[s.Current])
}

func (s *Scanner) peekString() string {
	return s.Source[s.Current-1 : s.Current]
}

func (s *Scanner) addToken(token tokens.TokenType, literal tokens.Literal) {
	text := s.Source[s.Start:s.Current]

	s.Tokens = append(s.Tokens, tokens.Token{TokenType: token, Lexeme: text, Literal: literal, Line: s.Line})
}

func (s *Scanner) isAtEnd() bool {
	return s.Current >= len(s.Source)
}

func (s *Scanner) peekAt(pos int) tokens.TokenType {
	return tokens.TokenType(s.Source[pos])
}
func (s *Scanner) advance() tokens.TokenType {
	s.Current++

	return tokens.TokenType(s.Source[s.Current-1])
}

func (s *Scanner) createString() {
	for s.peek() != tokens.PARENTHESES && !s.isAtEnd() {
		if s.peek() == tokens.NEWLINE {
			s.Line++
		}
		s.advance()
	}
	if s.isAtEnd() {
		s.ErrorList = append(s.ErrorList, fmt.Errorf("[Line %d] Error: Unterminated string.", s.Line))
		return
	}
	// Consume the final "
	s.advance()
	// trim quotes and add string token
	str := s.Source[s.Start+1 : s.Current-1]

	s.addToken(tokens.STRING, tokens.Literal{tokens.STRING_LITERAL, str})
}
func (s *Scanner) createNumber() {
	for s.isDigit(s.peek()) {
		s.advance()
	}

	if s.peek() == tokens.DOT && s.isDigit(s.peekNext()) {
		s.advance()
		for s.isDigit(s.peek()) {
			s.advance()
		}
	}
	str := s.Source[s.Start:s.Current]

	num, err := strconv.ParseFloat(str, 64)
	if err != nil {
		s.ErrorList = append(s.ErrorList, fmt.Errorf("[Line %d] Error: Invalid number.", s.Line))
		return
	}
	s.addToken(tokens.NUMBER, tokens.Literal{tokens.NUMBER_LITERAL, num})
}

func (s *Scanner) PrintTokens(tokens []tokens.Token) {

	for _, err := range s.ErrorList {
		fmt.Fprintln(os.Stderr, err)
	}

	for _, token := range tokens {
		fmt.Println(token.String())
	}
}

func (s *Scanner) isDigit(c tokens.TokenType) bool {
	return c >= "0" && c <= "9"
}

func (s *Scanner) match(expected tokens.TokenType) bool {
	if s.isAtEnd() {
		return false
	}
	peekVal := s.peekAt(s.Current)
	if peekVal != expected {
		return false
	}

	s.Current++
	return true
}

func (s *Scanner) peekNext() tokens.TokenType {
	if s.Current+1 >= len(s.Source) {
		return tokens.EOF
	}
	return tokens.TokenType(s.Source[s.Current+1])
}

func (s *Scanner) scanToken() {
	c := s.advance()
	switch c {
	case tokens.LEFT_PAREN:
		s.addToken(tokens.LEFT_PAREN, tokens.Literal{})
	case tokens.RIGHT_PAREN:
		s.addToken(tokens.RIGHT_PAREN, tokens.Literal{})
	case tokens.LEFT_BRACE:
		s.addToken(tokens.LEFT_BRACE, tokens.Literal{})
	case tokens.RIGHT_BRACE:
		s.addToken(tokens.RIGHT_BRACE, tokens.Literal{})
	case tokens.STAR:
		s.addToken(tokens.STAR, tokens.Literal{})
	case tokens.DOT:
		s.addToken(tokens.DOT, tokens.Literal{})
	case tokens.COMMA:
		s.addToken(tokens.COMMA, tokens.Literal{})
	case tokens.PLUS:
		s.addToken(tokens.PLUS, tokens.Literal{})
	case tokens.MINUS:
		s.addToken(tokens.MINUS, tokens.Literal{})
	case tokens.SEMICOLON:
		s.addToken(tokens.SEMICOLON, tokens.Literal{})
	case tokens.SLASH:
		if s.match(tokens.SLASH) {
			for s.peek() != tokens.NEWLINE && !s.isAtEnd() {
				s.advance()
			}
		} else {
			s.addToken(tokens.SLASH, tokens.Literal{})
		}

	case tokens.EOF:
		s.addToken(tokens.EOF, tokens.Literal{})
	case tokens.EQUAL:
		if s.match(tokens.EQUAL) {
			s.addToken(tokens.EQUAL_EQUAL, tokens.Literal{})
		} else {
			s.addToken(tokens.EQUAL, tokens.Literal{})
		}
	case tokens.BANG:
		if s.match(tokens.EQUAL) {
			s.addToken(tokens.BANG_EQUAL, tokens.Literal{})
		} else {
			s.addToken(tokens.BANG, tokens.Literal{})
		}
	case tokens.GREATER:
		if s.match(tokens.EQUAL) {
			s.addToken(tokens.GREATER_EQUAL, tokens.Literal{})
		} else {
			s.addToken(tokens.GREATER, tokens.Literal{})
		}
	case tokens.LESS:
		if s.match(tokens.EQUAL) {
			s.addToken(tokens.LESS_EQUAL, tokens.Literal{})
		} else {
			s.addToken(tokens.LESS, tokens.Literal{})
		}
	case tokens.CARRIAGE_RETURN, tokens.WHITESPACE, tokens.TABULATOR:

	case tokens.NEWLINE:
		s.Line++

	case tokens.PARENTHESES:
		s.createString()
	default:
		if s.isDigit(c) {
			s.createNumber()
			return
		} else if s.isAlpha(c) {
			s.createIdentifier()
			return
		}
		s.ErrorList = append(s.ErrorList, fmt.Errorf("[Line %d] Error: Unexpected character: %s", s.Line, s.peekString()))

	}
}

func (s *Scanner) isAlpha(c tokens.TokenType) bool {
	return (c >= "a" && c <= "z") ||
		(c >= "A" && c <= "Z") ||
		c == "_"
}

func (s *Scanner) isAlphaNumeric(c tokens.TokenType) bool {
	return s.isAlpha(c) || s.isDigit(c)
}

func (s *Scanner) createIdentifier() {
	for s.isAlphaNumeric(s.peek()) {
		s.advance()
	}
	str := s.Source[s.Start:s.Current]
	tokenType := tokens.TokenType(str)
	if tokens.TokenLoopUp[tokenType] != "" {
		s.addToken(tokenType, tokens.Literal{})
		return
	}
	s.addToken(tokens.IDENTIFIER, tokens.Literal{
		LiteralType: tokens.IDENTIFIER_LITERAL,
		Value:       str,
	})
}
