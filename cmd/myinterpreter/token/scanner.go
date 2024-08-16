package token

import (
	"fmt"
	"os"
	"strconv"
)

var NUMBERS = []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

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
	return s.current >= len(s.sourceRunes)
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
	case '"':
		s.addString()
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		s.addNumber()
	case '\n':
		s.line++
	case '\t', ' ':
		//Just skip
	default:
		s.logErrorRune(c)
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
	if s.matchCurrent('=') {
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
	if s.matchCurrent('/') {
		//Means that this is comment, need to skip until new line
		s.skipUntilNotMatches('\n')
	} else {
		s.addToken(SLASH)
	}
}

func (s *Scanner) addString() {
	s.skipUntilNotMatches('"', '\n')
	if s.isAtEnd() {
		s.logError("Unterminated string.")
		return
	}
	c := s.sourceRunes[s.current]
	if c == '\n' {
		s.logError("Unterminated string.")
	} else if c == '"' {
		s.current++
		s.addTokenWithLiteral(STRING, s.source[s.start+1:s.current-1])
	} else {
		panic("Something really wrong")
	}
}

func (s *Scanner) addNumber() {
	s.skipUntilMatches(NUMBERS...)
	if !s.matchCurrent('.') {
		s.addNumberToken()
	} else {
		if s.matchNext(NUMBERS...) {
			s.current++ //Skip dot as part of number
			s.skipUntilMatches(NUMBERS...)
			s.addNumberToken()
		} else {
			s.addNumberToken()
		}
	}
}

func (s *Scanner) addNumberToken() {
	n, _ := strconv.ParseFloat(s.source[s.start:s.current], 64)
	s.addTokenWithLiteral(NUMBER, n)
}

func (s *Scanner) skipUntilNotMatches(runes ...rune) {
	for s.current < len(s.sourceRunes) && !s.matchCurrent(runes...) {
		s.current++
	}
}

func (s *Scanner) skipUntilMatches(runes ...rune) {
	for s.current < len(s.sourceRunes) && s.matchCurrent(runes...) {
		s.current++
	}
}

func (s *Scanner) matchCurrent(runes ...rune) bool {
	return s.matchAtIndex(s.current, runes...)
}

func (s *Scanner) matchNext(runes ...rune) bool {
	return s.matchAtIndex(s.current+1, runes...)
}

func (s *Scanner) matchAtIndex(index int, runes ...rune) bool {
	if index >= len(s.sourceRunes) || index < 0 {
		return false
	}
	result := false
	for _, r := range runes {
		if s.sourceRunes[index] == r {
			result = true
			break
		}
	}
	return result
}

func (s *Scanner) addTokenWithLiteral(t Type, literal any) {
	text := s.source[s.start:s.current]
	s.tokens = append(s.tokens, NewToken(t, text, literal, s.line))
}

func (s *Scanner) PrintLines() {
	for _, t := range s.tokens {
		fmt.Fprintf(os.Stdout, "%s\n", t)
	}
}

func (s *Scanner) logError(message string) {
	s.errorsCount++
	fmt.Fprintf(os.Stderr, "[line %d] Error: %s\n", s.line, message)
}

func (s *Scanner) logErrorRune(r rune) {
	s.logError(fmt.Sprintf("Unexpected character: %c", r))
}

func (s *Scanner) GetExitCode() int {
	result := 0
	if s.errorsCount > 0 {
		result = 65
	}
	return result
}
