package interpreter

type Binary struct {
	Left     Expr
	Operator Token
	Right    Expr
}

func (s Binary) accept(v ExprVisitor) any {
	return v.visitBinary(&s)
}
