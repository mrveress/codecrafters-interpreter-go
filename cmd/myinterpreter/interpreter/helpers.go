package interpreter

import "fmt"

func num2str(num float64) string {
	if isIntegralNumber(num) {
		return fmt.Sprintf("%v.0", num)
	} else {
		return fmt.Sprintf("%v", num)
	}
}

func isIntegralNumber(num float64) bool {
	return num == float64(int(num))
}
