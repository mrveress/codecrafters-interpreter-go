package token

import "fmt"

type Token struct {
	tokenType Type
	lexeme    string
	literal   any
	line      int
}

func NewToken(tokenType Type, lexeme string, literal any, line int) Token {
	return Token{tokenType: tokenType, lexeme: lexeme, literal: literal, line: line}
}

func (t Token) String() string {
	switch t.literal.(type) {
	case float64:
		if t.isIntegralNumber() {
			return fmt.Sprintf("%s %s %v.0", t.tokenType, t.lexeme, t.literal)
		} else {
			return fmt.Sprintf("%s %s %v", t.tokenType, t.lexeme, t.literal)
		}
	default:
		return fmt.Sprintf("%s %s %v", t.tokenType, t.lexeme, t.literal)
	}
}

func (t *Token) isIntegralNumber() bool {
	switch t.literal.(type) {
	case float64:
		{
			l := t.literal.(float64)
			return l == float64(int(l))
		}
	default:
		return false
	}
}
