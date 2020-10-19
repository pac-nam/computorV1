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
	// i := 0
	for sqrtftoa(min) != sqrtftoa(max) {
		tmp := min + (max - min) / 2
		tmps := sqrtftoa(tmp)
		// fmt.Println(min, max, tmps, number)
		if tmps == sqrtftoa(min) || tmps == sqrtftoa(max) || tmp * tmp == number {
			return tmp
		}
		if tmp * tmp < number {
			min = tmp
		} else if tmp * tmp > number {
			max = tmp
		}
		// i++
		// if i > 10 {
		// 	return 0
		// }
	}
	return min
}

func	ftoa(nb float64) string {
	return fmt.Sprint(nb)
}