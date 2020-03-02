package polynom

import (
	"fmt"
	"strconv"
)

type S_solution struct {
	Real					float64			//	real part of the solution
	Imaginary				float64			//	imaginary part of the solution for Delta < 0
}

func (solution S_solution) String() string {
	return ftoa(solution.Imaginary) + "i + " +
		ftoa(solution.Real)
}

//	A polynom have the form: "aX^2 + bX + c"

type S_poly struct {
	Elems					[3]float64		//	Elem[0] == c | Elem[1] == b | Elem[2] == a
	Degree					int				//	degree of the polynom
	Delta					float64			//	discriminant of the polynom (b^2 - 4ac)
	Solutions				[]S_solution	//	array of solutions
	Specific				byte			//	special variable see SPECIFIC field
}

//	SPECIFIC
//	usualy set to 0
//	set to 1 if the polynom can't be solved
//	set to 2 if all reals work. In this case Solutions will be empty

func (poly S_poly) String() string {
	res := "{\n\tpolynom: " + ftoa(poly.Elems[2]) + "X^2 + " +
		ftoa(poly.Elems[1]) + "X + " +
		ftoa(poly.Elems[0]) + " = 0" +
		"\n\tdegree: " + strconv.Itoa(poly.Degree) +
		"\n\tdelta: " + ftoa(poly.Delta) +
		"\n\t" + strconv.Itoa(len(poly.Solutions)) + " solutions:\n\t{"
	for i, solution := range poly.Solutions {
		res += "\n\t\t" + strconv.Itoa(i) + ": " + fmt.Sprint(solution)
	}
	res += "\n\t}\n\tspecific: " + strconv.Itoa(int(poly.Specific)) + "\n}"
	return res
}