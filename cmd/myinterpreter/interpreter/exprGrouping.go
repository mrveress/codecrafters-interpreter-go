package interpreter

type Grouping struct {
	Expression Expr
}

func (s Grouping) accept(v ExprVisitor) string {
	return v.visitGrouping(&s)
}
