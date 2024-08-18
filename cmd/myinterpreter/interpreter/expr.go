package interpreter

type Expr interface {
	accept(ExprVisitor) any
}
