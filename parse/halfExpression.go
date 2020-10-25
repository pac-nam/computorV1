package parse

import (
	"strings"
)

const numberCharset = "0123456789"
const signCharset = "+-*^"

func missingSign(expr string) bool {
	canBeNumber := 1
	for _, c := range expr {
		//fmt.Println(string(c), " : canBenumber:", canBeNumber)
		if strings.ContainsRune(numberCharset, c) && canBeNumber != 0 {
			canBeNumber = 2
		} else if c == ' ' && canBeNumber == 2 {
			canBeNumber = 0
		} else if strings.ContainsRune(signCharset, c) {
			canBeNumber = 1
		} else if strings.ContainsRune(numberCharset, c) {
			return true
		}
	}
	return false
}

func invalidDot(expr string) bool {
	canBeDot := 1
	for _, c := range expr {
		//fmt.Println(string(c), " : canBeDot:", canBeDot)
		if c == '.' && canBeDot == 0 {
			canBeDot = 2
		} else if strings.ContainsRune(numberCharset, c) && canBeDot == 1 {
			canBeDot = 0
		} else if strings.ContainsRune("+-*^X ", c) {
			canBeDot = 1
		} else if c == '.' {
			return true
		}
	}
	return false
}

func doubleSign(expr string) bool {
	canBeSign := 1
	for _, c := range expr {
		//fmt.Println(string(c), " : canBeSign:", canBeSign)
		if strings.ContainsRune(signCharset, c) && canBeSign == 1 {
			canBeSign = 0
		} else if strings.ContainsRune(signCharset, c) {
			return true
		} else {
			canBeSign = 1
		}
	}
	return false
}

func halfExpression(expr string) ([3]float64, string) {
	res := [3]float64{0, 0, 0}
	if expr == "" {
		return [3]float64{0, 0, 0}, "missing a part of the expression"
	} else if missingSign(expr) {
		return [3]float64{0, 0, 0}, "missing sign"
	} else if invalidDot(expr) {
		return [3]float64{0, 0, 0}, "invalid dot"
	}
	expr = strings.Join(strings.Split(expr, " "), "")
	if doubleSign(expr) {
		return [3]float64{0, 0, 0}, "can not have double sign"
	}
	tab := strings.Split(expr, "+")
	for _, str := range tab {
		var err string
		res, err = parseFactor(str, res)
		if err != "" {
			return [3]float64{0, 0, 0}, err
		}
	}
	return res, ""
}
