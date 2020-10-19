package polynom

func degreeOne(poly *S_poly) {
	poly.Degree = 1
	tmp := -poly.Elems[0] / poly.Elems[1]
	if tmp == 0 {
		tmp = 0
	}
	poly.Solutions = append(poly.Solutions, S_solution{Real: tmp})
}