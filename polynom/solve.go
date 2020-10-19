package polynom

// import "fmt"

func Solve(elems [3]float64) S_poly {
	poly := new(S_poly)
	poly.Elems = elems
	if poly.Elems[2] == 0 {
		if poly.Elems[1] == 0 {
			if poly.Elems[0] == 0 {
				// fmt.Println("degree 0")
				degreeZeroSuccess(poly)
			} else {
				// fmt.Println("fail 0")
				degreeZeroFail(poly)
			}
		} else {
			// fmt.Println("degree 1")
			degreeOne(poly)
		}
	} else {
		// fmt.Println("degree 2")
		degreeTwo(poly)
	}
	return *poly
}