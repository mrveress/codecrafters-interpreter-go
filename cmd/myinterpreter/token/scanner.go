package token

import (
	"fmt"
	"os"
)

type Scanner struct {
	sourceRunes []rune
	source      string
	tokens      []Token
	start       int
	current     int
	line        int
}

func NewScanner(source string) Scanner {
	return Scanner{
		sourceRunes: []rune(source),
		source:      source,
		tokens:      make([]Token, 0),
		start:       0,
		current:     0,
		line:        1}
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) ScanTokens() []Token {
	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}

	s.tokens = append(s.tokens, NewToken(EOF, "", "null", s.line))
	return s.tokens
}

func (s *Scanner) scanToken() {
	c := s.advance()
	switch c {
	case '(':
		s.addToken(LEFT_PAREN)
	case ')':
		s.addToken(RIGHT_PAREN)
	case '{':
		s.addToken(LEFT_BRACE)
	case '}':
		s.addToken(RIGHT_BRACE)
	case ',':
		s.addToken(COMMA)
	case '.':
		s.addToken(DOT)
	case '-':
		s.addToken(MINUS)
	case '+':
		s.addToken(PLUS)
	case ';':
		s.addToken(SEMICOLON)
	case '*':
		s.addToken(STAR)
	case '\n':
		s.line++
	default:
		s.logError(c)
	}
}

func (s *Scanner) advance() rune {
	result := s.sourceRunes[s.current]
	s.current++
	return result
}

func (s *Scanner) addToken(t Type) {
	s.addTokenWithLiteral(t, "null")
}

func (s *Scanner) addTokenWithLiteral(t Type, literal string) {
	text := s.source[s.start:s.current]
	s.tokens = append(s.tokens, NewToken(t, text, literal, s.line))
}

func (s *Scanner) PrintLines() {
	for _, t := range s.tokens {
		fmt.Fprintf(os.Stdout, "%s\n", t)
	}
}

func (s *Scanner) logError(r rune) {
	fmt.Fprintf(os.Stderr, "[line %d] Error: Unexpected character: %c\n", s.line, r)
}
