package interpreter

type Unary struct {
	Operator Token
	Right    Expr
}

func (s Unary) accept(v ExprVisitor) any {
	return v.visitUnary(&s)
}
