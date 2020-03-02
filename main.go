package main

import (
	"computorV1/parse"
	"computorV1/polynom"
	"fmt"
	"os"
)

func reducedForm(tab [3]float64, degree int) string {
	str := " = 0"
	if tab[0] < 0 {
		str = "- " + fmt.Sprint(-tab[0]) + str
	} else {
		str = "+ " + fmt.Sprint(tab[0]) + str
	}
	if degree > 0 {
		if tab[1] < 0 {
			str = "- " + fmt.Sprint(-tab[1]) + " * X " + str
		} else {
			str = "+ " + fmt.Sprint(tab[1]) + " * X " + str
		}
		if degree == 2 {
			if tab[2] < 0 {
				str = "- " + fmt.Sprint(-tab[2]) + " * X^2 " + str
			} else {
				str = "+ " + fmt.Sprint(tab[2]) + " * X^2 " + str
			}
		}
	}
	return "Reduced form: " + str
}

func discriminant(r polynom.S_poly) string {
	str := "\nDiscriminant (" + fmt.Sprint(r.Delta) + ") is "
	if r.Delta > 0 {
		str += "strictly positive, the two solutions are:"
		str += "\n" + fmt.Sprint(r.Solutions[0].Real)
		str += "\n" + fmt.Sprint(r.Solutions[1].Real)
	} else if r.Delta < 0 {
		str += "strictly negative, the two imaginary solutions are:"
		str += "\n" + fmt.Sprint(r.Solutions[0].Imaginary)
		if r.Solutions[0].Real < 0 {
			str += " + " + fmt.Sprint(r.Solutions[0].Real)
		} else {
			str += " - " + fmt.Sprint(-r.Solutions[0].Real)
		}
		str += "\n" + fmt.Sprint(r.Solutions[1])
		if r.Solutions[1].Real < 0 {
			str += " + " + fmt.Sprint(r.Solutions[1].Real)
		} else {
			str += " - " + fmt.Sprint(-r.Solutions[1].Real)
		}
	} else {
		str += "null, the  solution is:"
		str += "\n" + fmt.Sprint(r.Solutions[0].Real)
	}
	return str
}

func printSolution(r polynom.S_poly) {
	str := reducedForm(r.Elems, r.Degree)
	str += "\nPolynomial degree" + fmt.Sprint(r.Degree)
	str += discriminant(r)
	fmt.Println(str)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing arguments\nUsage: " + os.Args[0] + " <expression>")
		return
	}
	poly, error := parse.GetPolynom(os.Args)
	if error != "" {
		fmt.Println("ERROR: " + error)
		return
	}
	result := polynom.Solve(poly)
	//fmt.Println(result)
	printSolution(result)
}
