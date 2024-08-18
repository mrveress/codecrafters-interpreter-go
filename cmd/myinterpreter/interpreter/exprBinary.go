package interpreter

type Binary struct {
	Left     Expr
	Operator Token
	Right    Expr
}

func (s Binary) accept(v ExprVisitor) string {
	return v.visitBinary(&s)
}
