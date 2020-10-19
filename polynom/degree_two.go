package polynom
// import "fmt"

func rmZeroNeg(nb float64) float64 {
	if nb == 0 {
		return 0
	}
	return nb
}

func posDelta(poly *S_poly, sqrtDelta float64) {
	poly.Solutions = append(poly.Solutions,
		S_solution{Real: rmZeroNeg((-poly.Elems[1] + sqrtDelta) / (2 * poly.Elems[2]))},
		S_solution{Real: rmZeroNeg((-poly.Elems[1] - sqrtDelta) / (2 * poly.Elems[2]))})
}

func negDelta(poly *S_poly, sqrtDelta float64) {
	poly.Solutions = append(poly.Solutions,
		S_solution{Real: rmZeroNeg(-poly.Elems[1] / (2 * poly.Elems[2])),
			Imaginary: sqrtDelta / (2 * poly.Elems[2])},
		S_solution{Real: rmZeroNeg(-poly.Elems[1] / (2 * poly.Elems[2])),
			Imaginary: sqrtDelta / (2 * poly.Elems[2])})
}

func nulDelta(poly *S_poly) {
	poly.Solutions = append(poly.Solutions,
		S_solution{Real: rmZeroNeg(-poly.Elems[1] / (2 * poly.Elems[2]))})
}

func degreeTwo(poly *S_poly) {
	poly.Degree = 2
	poly.Delta = poly.Elems[1] * poly.Elems[1] - 4 * poly.Elems[0] * poly.Elems[2]
	// fmt.Println(poly.Delta)
	if poly.Delta > 0 {
		posDelta(poly, sqrt(poly.Delta))
	} else if poly.Delta < 0 {
		negDelta(poly, sqrt(-poly.Delta))
	} else {
		nulDelta(poly,)
	}
}