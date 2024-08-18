package interpreter

import "fmt"

type Token struct {
	TokenType TokenType
	Lexeme    string
	Literal   any
	Line      int
}

func NewToken(tokenType TokenType, lexeme string, literal any, line int) Token {
	return Token{TokenType: tokenType, Lexeme: lexeme, Literal: literal, Line: line}
}

func (t Token) String() string {
	switch t.Literal.(type) {
	case float64:
		return fmt.Sprintf("%s %s %s", t.TokenType, t.Lexeme, num2str(t.Literal.(float64)))
	case nil:
		return fmt.Sprintf("%s %s %s", t.TokenType, t.Lexeme, "null")
	default:
		return fmt.Sprintf("%s %s %v", t.TokenType, t.Lexeme, t.Literal)
	}
}
