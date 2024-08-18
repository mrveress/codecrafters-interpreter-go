package interpreter

import "fmt"

type AstPrinter struct {
}

func (ap AstPrinter) Print(expr Expr) {
	fmt.Println(ap.GetString(expr))
}

func (ap AstPrinter) GetString(expr Expr) string {
	return expr.accept(&ap).(string)
}

func (ap AstPrinter) visitUnary(expr *Unary) any {
	return ap.parenthesize(expr.Operator.Lexeme, expr.Right)
}

func (ap AstPrinter) visitBinary(expr *Binary) any {
	return ap.parenthesize(expr.Operator.Lexeme, expr.Left, expr.Right)
}

func (ap AstPrinter) visitLiteral(expr *Literal) any {
	if expr.Value == nil {
		return "nil"
	}
	return expr.String()
}

func (ap AstPrinter) visitGrouping(expr *Grouping) any {
	return ap.parenthesize("group", expr.Expression)
}

func (ap AstPrinter) parenthesize(name string, exprs ...Expr) string {
	result := "(" + name

	for _, expr := range exprs {
		result = result + " " + expr.accept(ap).(string)
	}
	result = result + ")"
	return result
}
