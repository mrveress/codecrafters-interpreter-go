package interpreter

import (
	"fmt"
	"os"
)

type Parser struct {
	Tokens  []Token
	Current int
}

func NewParser(tokens []Token) Parser {
	return Parser{tokens, 0}
}

func (p *Parser) match(types ...TokenType) bool {
	for _, t := range types {
		if p.check(t) {
			p.advance()
			return true
		}
	}
	return false
}

func (p *Parser) check(t TokenType) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().TokenType == t
}

func (p *Parser) advance() Token {
	if !p.isAtEnd() {
		p.Current = p.Current + 1
	}
	return p.previous()
}

func (p *Parser) peek() Token {
	return p.Tokens[p.Current]
}

func (p *Parser) isAtEnd() bool {
	return p.peek().TokenType == EOF
}

func (p *Parser) previous() Token {
	return p.Tokens[p.Current-1]
}

//-------------------------------------------------

func (p *Parser) Parse() Expr {
	return p.expression()
}

//-------------------------------------------------

func (p *Parser) expression() Expr {
	return p.equality()
}

func (p *Parser) equality() Expr {
	expr := p.comparison()
	for p.match(BANG_EQUAL, EQUAL_EQUAL) {
		expr = Binary{expr, p.previous(), p.comparison()}
	}
	return expr
}

func (p *Parser) comparison() Expr {
	expr := p.term()
	for p.match(GREATER, GREATER_EQUAL, LESS, LESS_EQUAL) {
		expr = Binary{expr, p.previous(), p.term()}
	}
	return expr
}

func (p *Parser) term() Expr {
	expr := p.factor()
	for p.match(MINUS, PLUS) {
		expr = Binary{expr, p.previous(), p.factor()}
	}
	return expr
}

func (p *Parser) factor() Expr {
	expr := p.unary()
	for p.match(SLASH, STAR) {
		expr = Binary{expr, p.previous(), p.unary()}
	}
	return expr
}

func (p *Parser) unary() Expr {
	if p.match(BANG, MINUS) {
		return Unary{p.previous(), p.unary()}
	}
	return p.primary()
}

func (p *Parser) primary() Expr {
	if p.match(FALSE) {
		return Literal{false}
	}
	if p.match(TRUE) {
		return Literal{true}
	}
	if p.match(NIL) {
		return Literal{nil}
	}

	if p.match(NUMBER, STRING) {
		return Literal{p.previous().Literal}
	}

	if p.match(LEFT_PAREN) {
		if p.match(RIGHT_PAREN) {
			p.error("Empty parentheses.")
		}
		expr := p.expression()
		p.consume(RIGHT_PAREN, "Unmatched parentheses.")
		return Grouping{expr}
	}

	p.error("Incorrect or empty input.")
	panic("Parser.primary(): Unhandled Exception")
}

func (p *Parser) consume(t TokenType, message string) Token {
	if p.check(t) {
		return p.advance()
	}
	p.error(message)
	return Token{} //Just placeholder
}

func (p *Parser) error(message string) {
	fmt.Fprintf(os.Stderr, "Error: %s\n", message)
	os.Exit(65)
}
