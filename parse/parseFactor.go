package parse

import (
	"fmt"
	"strconv"
	"strings"
)

func parseX(expr string) (float64, int, string) {
	var err string
	multiplicator := float64(1)
	index := 0
	Xstart := strings.Index(expr, "X")
	if Xstart != 0 {
		multiplicator, err = parseNb(expr[:Xstart])
		if err != "" {
			return 0, 0, err
		}
	}
	if Xstart+1 == len(expr) {
		index = 1
	} else if Xstart+3 == len(expr) {
		if expr[Xstart+1] != '^' {
			return 0, 0, "wrong character after X"
		} else if expr[Xstart+2] == '0' {
			index = 0
		} else if expr[Xstart+2] == '1' {
			index = 1
		} else if expr[Xstart+2] == '2' {
			index = 2
		}
	} else {
		return 0, 0, "error with X power"
	}
	return multiplicator, index, ""
}

func parseNb(expr string) (float64, string) {
	if strings.ContainsRune(expr, '^') {
		return 0, "invalid character '^'"
	} else if strings.Count(expr, ".") > 1 {
		return 0, "multiple dots"
	}
	nb, err := strconv.ParseFloat(expr, 64)
	if err != nil {
		return 0, fmt.Sprint(err)
	}
	return nb, ""
}

func parseOnlymultiple(expr string, res [3]float64) ([3]float64, string) {
	var err string
	resNb, tmp := float64(1), float64(0)
	resIndex := 0
	if strings.Count(expr, "X") > 1 {
		return res, "Invalid expression type X * X"
	}
	tab := strings.Split(expr, "*")
	for _, str := range tab {
		if strings.ContainsRune(str, 'X') {
			tmp, resIndex, err = parseX(str)
			if err != "" {
				return res, err
			}
			resNb *= tmp
		} else {
			tmp, err := parseNb(str)
			if err != "" {
				return res, err
			}
			resNb *= tmp
		}
	}
	res[resIndex] = resNb
	return res, ""
}

func parseFactor(expr string, res [3]float64) ([3]float64, string) {
	var err string
	tab := strings.Split(expr, "-")
	//fmt.Println(tab)
	if tab[0] != "" {
		res, err = parseOnlymultiple(tab[0], res)
	}
	if err != "" {
		return res, err
	}
	for i := 1; i < len(tab); i++ {
		if tab[i] != "" {
			res2, err := parseOnlymultiple(tab[i], [3]float64{0, 0, 0})
			if err != "" {
				return res, err
			}
			res[0] -= res2[0]
			res[1] -= res2[1]
			res[2] -= res2[2]
		}
	}
	return res, ""
}
