package scanning

import (
	"fmt"
	"os"
	"strconv"
)

type Scanner struct {
	Source    string
	Tokens    []Token
	Start     int
	Current   int
	Line      int
	ErrorList []error
}

func (s *Scanner) ScanTokens() ([]Token, []error) {
	for !s.isAtEnd() {
		s.Start = s.Current
		s.scanToken()
	}
	s.Tokens = append(s.Tokens, Token{EOF, "", Literal{}, s.Line})
	return s.Tokens, s.ErrorList
}

func (s *Scanner) peek() TokenType {
	if s.isAtEnd() {
		return EOF
	}
	return TokenType(s.Source[s.Current])
}

func (s *Scanner) peekString() string {
	return s.Source[s.Current-1 : s.Current]
}

func (s *Scanner) addToken(token TokenType, literal Literal) {
	text := s.Source[s.Start:s.Current]

	s.Tokens = append(s.Tokens, Token{tokenType: token, lexeme: text, literal: literal, line: s.Line})
}

func (s *Scanner) isAtEnd() bool {
	return s.Current >= len(s.Source)
}

func (s *Scanner) peekAt(pos int) TokenType {
	return TokenType(s.Source[pos])
}
func (s *Scanner) advance() TokenType {
	s.Current++

	return TokenType(s.Source[s.Current-1])
}

func (s *Scanner) createString() {
	for s.peek() != PARENTHESES && !s.isAtEnd() {
		if s.peek() == NEWLINE {
			s.Line++
		}
		s.advance()
	}
	if s.isAtEnd() {
		s.ErrorList = append(s.ErrorList, fmt.Errorf("[line %d] Error: Unterminated string.", s.Line))
		return
	}
	// Consume the final "
	s.advance()
	// trim quotes and add string token
	str := s.Source[s.Start+1 : s.Current-1]

	s.addToken(STRING, Literal{STRING_LITERAL, str})
}
func (s *Scanner) createNumber() {
	for s.isDigit(s.peek()) {
		s.advance()
	}

	if s.peek() == DOT && s.isDigit(s.peekNext()) {
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
	s.addToken(NUMBER, Literal{NUMBER_LITERAL, num})
}

func (s *Scanner) PrintTokens(tokens []Token) {

	for _, err := range s.ErrorList {
		fmt.Fprintln(os.Stderr, err)
	}

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
	peekVal := s.peekAt(s.Current)
	if peekVal != expected {
		return false
	}

	s.Current++
	return true
}

func (s *Scanner) peekNext() TokenType {
	if s.Current+1 >= len(s.Source) {
		return EOF
	}
	return TokenType(s.Source[s.Current+1])
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
		s.Line++

	case PARENTHESES:
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

func (s *Scanner) isAlpha(c TokenType) bool {
	return (c >= "a" && c <= "z") ||
		(c >= "A" && c <= "Z") ||
		c == "_"
}

func (s *Scanner) isAlphaNumeric(c TokenType) bool {
	return s.isAlpha(c) || s.isDigit(c)
}

func (s *Scanner) createIdentifier() {
	for s.isAlphaNumeric(s.peek()) {
		s.advance()
	}
	str := s.Source[s.Start:s.Current]
	tokenType := TokenType(str)
	if tokenLoopUp[tokenType] != "" {
		s.addToken(tokenType, Literal{})
		return
	}
	s.addToken(IDENTIFIER, Literal{IDENTIFIER_LITERAL, str})
}
