package interpreter

type ExprVisitor interface {
	visitUnary(*Unary) any
	visitBinary(*Binary) any
	visitLiteral(*Literal) any
	visitGrouping(*Grouping) any
}
