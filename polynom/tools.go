package polynom

import (
	"fmt"
)

func	sqrtftoa(nb float64) string {
	str := fmt.Sprint(nb)
	if len(str)  > 15 {
		return str[:15]
	}
	return str
}

func	sqrt(number float64) float64 {
	if number < 0 {
		return 0
	}
	min, max := float64(0), number
	for sqrtftoa(min) != sqrtftoa(max) {
		//fmt.Println(min, max)
		tmp := min + (max - min) / 2
		tmps := sqrtftoa(tmp)
		if tmps == sqrtftoa(min) || tmps == sqrtftoa(max) {
			return tmp
		}
		if tmp * tmp < number {
			min = tmp
		} else if tmp * tmp > number {
			max = tmp
		}
	}
	return min
}

func	ftoa(nb float64) string {
	return fmt.Sprint(nb)
}