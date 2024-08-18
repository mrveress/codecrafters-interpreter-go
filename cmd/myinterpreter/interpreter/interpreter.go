package interpreter

import (
	"fmt"
)

type Interpreter struct{}

func NewInterpreter() Interpreter {
	return Interpreter{}
}

func (i *Interpreter) visitLiteral(expr *Literal) any {
	return expr.Value
}

func (i *Interpreter) visitGrouping(expr *Grouping) any {
	return i.evaluate(&expr.Expression)
}

func (i *Interpreter) visitUnary(expr *Unary) any {
	right := i.evaluate(&expr.Right)
	switch expr.Operator.TokenType {
	case MINUS:
		i.checkNumberOperand(expr.Operator, right)
		return -(right.(float64))
	case BANG:
		return !i.isTruthy(right)
	default:
		return nil
	}
}

func (i *Interpreter) visitBinary(expr *Binary) any {
	left := i.evaluate(&expr.Left)
	right := i.evaluate(&expr.Right)

	switch expr.Operator.TokenType {
	case MINUS:
		i.checkNumberOperands(expr.Operator, left, right)
		return (left.(float64)) - (right.(float64))
	case SLASH:
		i.checkNumberOperands(expr.Operator, left, right)
		return (left.(float64)) / (right.(float64))
	case STAR:
		i.checkNumberOperands(expr.Operator, left, right)
		return (left.(float64)) * (right.(float64))
	case PLUS:
		i.checkNumbersOrStrings(expr.Operator, left, right)
		switch left.(type) {
		case float64:
			return (left.(float64)) + (right.(float64))
		case string:
			return (left.(string)) + (right.(string))
		default:
			return nil
		}
	case GREATER:
		i.checkNumberOperands(expr.Operator, left, right)
		return (left.(float64)) > (right.(float64))
	case GREATER_EQUAL:
		i.checkNumberOperands(expr.Operator, left, right)
		return (left.(float64)) >= (right.(float64))
	case LESS:
		i.checkNumberOperands(expr.Operator, left, right)
		return (left.(float64)) < (right.(float64))
	case LESS_EQUAL:
		i.checkNumberOperands(expr.Operator, left, right)
		return (left.(float64)) <= (right.(float64))

	case BANG_EQUAL:
		return !i.isEqual(left, right)
	case EQUAL_EQUAL:
		return i.isEqual(left, right)
	default:
		return nil
	}
}

func (i *Interpreter) evaluate(expr *Expr) any {
	return (*expr).accept(i)
}

func (i *Interpreter) isTruthy(obj any) bool {
	if obj == nil {
		return false
	}
	switch obj.(type) {
	case bool:
		return obj.(bool)
	default:
		return true
	}
}

func (i *Interpreter) isEqual(a any, b any) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil {
		return false
	}
	return a == b
}

func (i *Interpreter) checkNumberOperand(operator Token, operand any) {
	switch operand.(type) {
	case float64:
		return
	default:
		i.errorWithOperatorAndCode(operator, "Operand must be a number.", 70)
	}
}

func (i *Interpreter) checkNumberOperands(operator Token, left any, right any) {
	switch left.(type) {
	case float64:
		switch right.(type) {
		case float64:
			return
		default:
			i.errorWithOperatorAndCode(operator, "Operands must be numbers.", 70)
		}
	default:
		i.errorWithOperatorAndCode(operator, "Operands must be numbers.", 70)
	}
}

func (i *Interpreter) checkNumbersOrStrings(operator Token, left any, right any) {
	switch left.(type) {
	case float64:
		switch right.(type) {
		case float64:
			return
		default:
			i.errorWithOperatorAndCode(operator, "Operands must be two numbers or two strings.", 70)
		}
	case string:
		switch right.(type) {
		case string:
			return
		default:
			i.errorWithOperatorAndCode(operator, "Operands must be two numbers or two strings.", 70)
		}
	default:
		i.errorWithOperatorAndCode(operator, "Operands must be two numbers or two strings.", 70)
	}
}

func (i *Interpreter) errorWithOperatorAndCode(operator Token, message string, code int) {
	Errorf(code, "Interpretator Error at '%s': %s ", operator.Lexeme, message)
}

func (i *Interpreter) Interpret(expression Expr) string {
	value := i.evaluate(&expression)
	return i.stringify(value)
}

func (i *Interpreter) stringify(object any) string {
	if object == nil {
		return "nil"
	}
	return fmt.Sprintf("%v", object)
}
