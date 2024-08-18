package interpreter

import (
	"fmt"
	"os"
	"strconv"
)

type Scanner struct {
	SourceRunes []rune
	Source      string
	Tokens      []Token
	Start       int
	Current     int
	Line        int
	ErrorsCount int
}

func NewScanner(source string) Scanner {
	return Scanner{
		SourceRunes: []rune(source),
		Source:      source,
		Tokens:      make([]Token, 0),
		Start:       0,
		Current:     0,
		Line:        1,
		ErrorsCount: 0}
}

func (s *Scanner) isAtEnd() bool {
	return s.Current >= len(s.SourceRunes)
}

func (s *Scanner) ScanTokens() []Token {
	for !s.isAtEnd() {
		s.Start = s.Current
		s.scanToken()
	}

	s.Tokens = append(s.Tokens, NewToken(EOF, "", nil, s.Line))
	return s.Tokens
}

func (s *Scanner) scanToken() {
	c := s.advance()
	switch c {
	case '=', '!', '<', '>':
		s.addComplexToken(c)
	case '/':
		s.addSlashOrIgnoreComment()
	case '"':
		s.addString()
	case '\n':
		s.Line++
	case '\t', ' ':
		//Just skip
	default:
		{
			switch {
			case isOneRuneToken(c):
				s.addToken(RuneTokens[c])
			case isNumeric(c):
				s.addNumber()
			case isAlpha(c):
				s.addIdentifierOrKeyword()
			default:
				s.logErrorRune(c)
			}
		}
	}
}

func isOneRuneToken(c rune) bool {
	switch c {
	case '(', ')', '{', '}', ',', '.', '-', '+', ';', '*':
		return true
	default:
		return false
	}
}
func isNumeric(c rune) bool      { return c >= '0' && c <= '9' }
func isAlpha(c rune) bool        { return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_' }
func isAlphaNumeric(c rune) bool { return isNumeric(c) || isAlpha(c) }

func (s *Scanner) advance() rune {
	result := s.SourceRunes[s.Current]
	s.Current++
	return result
}

func (s *Scanner) addToken(t TokenType) {
	s.addTokenWithLiteral(t, nil)
}

func (s *Scanner) addComplexToken(c rune) {
	if s.matchCurrent('=') {
		s.Current++
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
		//Means this is comment, need to skip until new line
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
	c := s.SourceRunes[s.Current]
	if c == '\n' {
		s.logError("Unterminated string.")
	} else if c == '"' {
		s.Current++
		s.addTokenWithLiteral(STRING, s.Source[s.Start+1:s.Current-1])
	} else {
		panic("Something really wrong")
	}
}

func (s *Scanner) addNumber() {
	s.skipUntilPredicate(isNumeric)
	if !s.matchCurrent('.') {
		s.addNumberToken()
	} else {
		if s.matchNextPredicate(isNumeric) {
			s.Current++ //Skip dot as part of number
			s.skipUntilPredicate(isNumeric)
			s.addNumberToken()
		} else {
			s.addNumberToken()
		}
	}
}

func (s *Scanner) addNumberToken() {
	n, _ := strconv.ParseFloat(s.Source[s.Start:s.Current], 64)
	s.addTokenWithLiteral(NUMBER, n)
}

func (s *Scanner) addIdentifierOrKeyword() {
	s.skipUntilPredicate(isAlphaNumeric)
	s.addIdentifierOrKeywordToken()
}

func (s *Scanner) skipUntilNotMatches(runes ...rune) {
	for !s.isAtEnd() && !s.matchCurrent(runes...) {
		s.Current++
	}
}

func (s *Scanner) skipUntilMatches(runes ...rune) {
	for !s.isAtEnd() && s.matchCurrent(runes...) {
		s.Current++
	}
}

func (s *Scanner) skipUntilPredicate(pred func(rune) bool) {
	for !s.isAtEnd() && pred(s.SourceRunes[s.Current]) {
		s.Current++
	}
}

func (s *Scanner) matchCurrent(runes ...rune) bool {
	return s.matchAtIndex(s.Current, runes...)
}

func (s *Scanner) matchCurrentPredicate(pred func(rune) bool) bool {
	return s.matchAtIndexPredicate(s.Current, pred)
}

func (s *Scanner) matchNext(runes ...rune) bool {
	return s.matchAtIndex(s.Current+1, runes...)
}

func (s *Scanner) matchNextPredicate(pred func(rune) bool) bool {
	return s.matchAtIndexPredicate(s.Current+1, pred)
}

func (s *Scanner) matchAtIndex(index int, runes ...rune) bool {
	if index >= len(s.SourceRunes) || index < 0 {
		return false
	}
	result := false
	for _, r := range runes {
		if s.SourceRunes[index] == r {
			result = true
			break
		}
	}
	return result
}

func (s *Scanner) matchAtIndexPredicate(index int, pred func(rune) bool) bool {
	if index >= len(s.SourceRunes) || index < 0 {
		return false
	}
	return pred(s.SourceRunes[index])
}

func (s *Scanner) addIdentifierOrKeywordToken() {
	text := s.Source[s.Start:s.Current]
	val, ok := Keywords[text]
	if ok {
		s.addTokenWithLiteral(val, nil)
	} else {
		s.addTokenWithLiteral(IDENTIFIER, nil)
	}
}

func (s *Scanner) addTokenWithLiteral(t TokenType, literal any) {
	text := s.Source[s.Start:s.Current]
	s.Tokens = append(s.Tokens, NewToken(t, text, literal, s.Line))
}

func (s *Scanner) PrintLines() {
	for _, t := range s.Tokens {
		Fprintf(os.Stdout, "%s\n", t)
	}
}

func (s *Scanner) logError(message string) {
	s.ErrorsCount++
	Fprintf(os.Stderr, "[line %d] Error: %s\n", s.Line, message)
}

func (s *Scanner) logErrorRune(r rune) {
	s.logError(fmt.Sprintf("Unexpected character: %c", r))
}

func (s *Scanner) GetExitCode() int {
	result := 0
	if s.ErrorsCount > 0 {
		result = 65
	}
	return result
}
