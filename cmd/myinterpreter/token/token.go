package token

import "fmt"

type Token struct {
	tokenType Type
	lexeme    string
	literal   string
	line      int
}

func NewToken(tokenType Type, lexeme string, literal string, line int) Token {
	return Token{tokenType: tokenType, lexeme: lexeme, literal: literal, line: line}
}

func (t Token) String() string {
	return fmt.Sprintf("%s %s %s", t.tokenType, t.lexeme, t.literal)
}
