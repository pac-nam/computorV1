package parse

import (
	"strings"
)

func epurSpaces(args []string) string {
	args = strings.Split(strings.Join(args, " "), " ")
	for i := 0; i < len(args); i++ {
		for i < len(args) && args[i] == "" {
			args = append(args[:i], args[i+1:]...)
		}
	}
	return strings.Join(args, " ")
}

func validCharset(str string, charset string) bool {
	for _, c := range str {
		if !strings.ContainsRune(charset, c) {
			return false
		}
	}
	return true
}

func GetPolynom(args []string) ([3]float64, string) {
	str := epurSpaces(args[1:])
	if !validCharset(str, "0123456789 .*+-=X^") {
		return [3]float64{0, 0, 0}, "invalid character"
	}
	expression := strings.Split(str, "=")
	if len(expression) != 2 {
		return [3]float64{0, 0, 0}, "only one '=' is accepted"
	}
	left, err := halfExpression(expression[0])
	if err != "" {
		return [3]float64{0, 0, 0}, err
	}
	right, err := halfExpression(expression[1])
	if err != "" {
		return [3]float64{0, 0, 0}, err
	}
	left[0] -= right[0]
	left[1] -= right[1]
	left[2] -= right[2]
	return left, ""
}
