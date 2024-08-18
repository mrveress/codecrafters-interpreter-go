package interpreter

type Grouping struct {
	Expression Expr
}

func (s Grouping) accept(v ExprVisitor) any {
	return v.visitGrouping(&s)
}
