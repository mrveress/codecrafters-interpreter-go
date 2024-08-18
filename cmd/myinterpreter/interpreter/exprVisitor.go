package interpreter

type ExprVisitor interface {
	visitUnary(*Unary) string
	visitBinary(*Binary) string
	visitLiteral(*Literal) string
	visitGrouping(*Grouping) string
}
