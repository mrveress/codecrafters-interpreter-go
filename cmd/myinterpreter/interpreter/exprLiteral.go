package interpreter

import "fmt"

type Literal struct {
	Value any
}

func (s Literal) accept(v ExprVisitor) string {
	return v.visitLiteral(&s)
}

func (s Literal) String() string {
	switch s.Value.(type) {
	case nil:
		return "nil"
	case string:
		return s.Value.(string)
	case int:
		return fmt.Sprintf("%d", s.Value.(int))
	case float64:
		return num2str(s.Value.(float64))
	case bool:
		if s.Value.(bool) {
			return "true"
		} else {
			return "false"
		}
	default:
		return "LITERAL_ERROR"
	}
}
