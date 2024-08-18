package interpreter

type Unary struct {
	Operator Token
	Right    Expr
}

func (s Unary) accept(v ExprVisitor) string {
	return v.visitUnary(&s)
}
