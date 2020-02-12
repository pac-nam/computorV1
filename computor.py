#!/usr/bin/python
import sys
import lexer_parser

def sqrt(nb):
	low = float(1)
	high = float(1)
	if (nb < 1):
		while(low * low > nb):
			high = low
			low = low / 2
	else:
		while (high * high < nb):
			low = high
			high = high * 2
	while (str(high) != str(low)):
		tmp = high - (high - low) / 2
		if (tmp * tmp >= nb):
			high = tmp
		if (tmp * tmp <= nb):
			low = tmp
	return high


def second_degree(e):
	D = float(e[1] * e[1] - 4 * e[0] * e[2])
	if (D > 0):
		print("Discriminant ({0}) is strictly positive, the two solutions are:".format(D))
		print((-e[1] - sqrt(D)) / (2 * e[0]))
		print((-e[1] + sqrt(D)) / (2 * e[0]))
	elif (D == 0):
		print("Discriminant is null, the solution is:")
		print(e[1] / (2 * e[0]))
	else:
		print("Discriminant ({0}) is strictly negative, the two complexe solutions are:".format(D))
		print("{0} * i + {1}".format(-sqrt(-D) / (2 * e[0]), e[1] / (2 * e[0])))
		print("{0} * i + {1}".format(sqrt(-D) / (2 * e[0]), e[1] / (2 * e[0])))

def main():
	if (len(sys.argv) < 2):
		print("Missing equation")
		return
	if (!(elems = lexer_parser(" ".join(sys.argv[1:]).split())))
		print("parsing error")
		return
	if (elems[0] != 0):
		print("Reduced form: {0} * X^2 + {1} * X + {2} = 0".format(elems[0], elems[1], elems[2]))
	elif (elems[1] != 0):
		print("Reduced form: {0} * X + {1} = 0".format(elems[1], elems[2]))
	else:
		print("Reduced form: {0} = 0".format(elems[2]))
	if (elems[0] != 0):
		print("Polynomial degree: 2")
		second_degree(elems)
	elif (elems[1] != 0):
		print("Polynomial degree: 1\nThe solution is:")
		if (elems[2] == 0):
			print("0")
		else:
			print(-elems[2] / elems[1])
	elif (elems[2] != 0):
		print("Polynomial degree: 0\nUnsolvable")
	else:
		print("0 = 0 all real numbers will do the job")

main()
