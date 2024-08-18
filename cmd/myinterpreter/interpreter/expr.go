package interpreter

type Expr interface {
	accept(ExprVisitor) string
}
