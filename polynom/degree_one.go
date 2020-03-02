package polynom

func degreeOne(poly *S_poly) {
	poly.Degree = 1
	poly.Solutions = append(poly.Solutions, S_solution{Real: -poly.Elems[0] / poly.Elems[1]})
}