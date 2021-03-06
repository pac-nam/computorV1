package main

import (
	"computorV1/parse"
	"computorV1/polynom"
	"fmt"
	"os"
)

func SP(nb float64) float64 {
	if nb == 0 {
		return 0
	}
	return nb
}

func reducedForm(tab [3]float64, degree int) string {
	str := " = 0"
	if tab[0] < 0 {
		str = "- " + fmt.Sprint(SP(-tab[0])) + str
	} else {
		str = "+ " + fmt.Sprint(tab[0]) + str
	}
	if degree > 0 {
		if tab[1] < 0 {
			str = "- " + fmt.Sprint(SP(-tab[1])) + " * X " + str
		} else {
			str = "+ " + fmt.Sprint(tab[1]) + " * X " + str
		}
		if degree == 2 {
			if tab[2] < 0 {
				str = "- " + fmt.Sprint(SP(-tab[2])) + " * X^2 " + str
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
			str += "i + " + fmt.Sprint(r.Solutions[0].Real)
		} else {
			str += "i - " + fmt.Sprint(SP(-r.Solutions[0].Real))
		}
		str += "\n" + fmt.Sprint(r.Solutions[1].Imaginary)
		if r.Solutions[1].Real < 0 {
			str += "i + " + fmt.Sprint(r.Solutions[1].Real)
		} else {
			str += "i - " + fmt.Sprint(SP(-r.Solutions[1].Real))
		}
	} else {
		str += "null, the  solution is:"
		str += "\n" + fmt.Sprint(r.Solutions[0].Real)
	}
	return str
}

func printSolution(r polynom.S_poly) {
	str := reducedForm(r.Elems, r.Degree)
	str += "\nPolynomial degree " + fmt.Sprint(r.Degree)
	if r.Degree == 2 {
		str += discriminant(r)
	} else if r.Degree == 1 {
		str += fmt.Sprint("\nthe solution is: X = ", r.Solutions[0].Real)
	} else if r.Specific == 1 {
		str += "\nimpossible expression"
	} else if r.Specific == 2 {
		str += "\nall real number work"
	}
	fmt.Println(str)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing arguments\nUsage: " + os.Args[0] + " <expression>")
		return
	}
	// fmt.Println("hello")
	poly, error := parse.GetPolynom(os.Args)
	if error != "" {
		fmt.Println("ERROR: " + error)
		return
	}
	// fmt.Println(poly)
	result := polynom.Solve(poly)
	// fmt.Println(result)
	printSolution(result)
}
