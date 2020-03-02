package polynom

func degreeZeroSuccess(poly *S_poly) {
	poly.Degree = 0
	poly.Specific = 2
}
func degreeZeroFail(poly *S_poly) {
	poly.Degree = 0
	poly.Specific = 1
}