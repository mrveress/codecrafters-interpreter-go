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
	errorsCount int
}

func NewScanner(source string) Scanner {
	return Scanner{
		sourceRunes: []rune(source),
		source:      source,
		tokens:      make([]Token, 0),
		start:       0,
		current:     0,
		line:        1,
		errorsCount: 0}
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
	case '=', '!', '<', '>':
		s.addComplexToken(c)
	case '/':
		s.addSlashOrIgnoreComment()
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

func (s *Scanner) addComplexToken(c rune) {
	if s.matchNext('=') {
		s.current++
		switch c {
		case '=':
			s.addToken(EQUAL_EQUAL)
		case '!':
			s.addToken(BANG_EQUAL)
		case '<':
			s.addToken(LESS_EQUAL)
		case '>':
			s.addToken(GREATER_EQUAL)
		}
	} else {
		switch c {
		case '=':
			s.addToken(EQUAL)
		case '!':
			s.addToken(BANG)
		case '<':
			s.addToken(LESS)
		case '>':
			s.addToken(GREATER)
		}
	}
}

func (s *Scanner) addSlashOrIgnoreComment() {
	if s.matchNext('/') {
		//Means that this is comment, need to skip until new line
		s.skipUntil('\n')
	} else {
		s.addToken(SLASH)
	}
}

func (s *Scanner) skipUntil(c rune) {
	for !s.isAtEnd() && !s.matchNext(c) {
		s.current++
	}
}

func (s *Scanner) matchNext(c rune) bool {
	if s.current >= len(s.source) {
		return false
	}
	return s.sourceRunes[s.current] == c
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
	s.errorsCount++
	fmt.Fprintf(os.Stderr, "[line %d] Error: Unexpected character: %c\n", s.line, r)
}

func (s *Scanner) GetExitCode() int {
	result := 0
	if s.errorsCount > 0 {
		result = 65
	}
	return result
}
