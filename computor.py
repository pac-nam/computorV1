#!/usr/bin/python
import sys

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

def reduce_expression(parameters):
	elems = [0, 0, 0]
	i = 0
	neg = 1
	while (i < len(parameters)):
		if (parameters[i - 1] == "="):
			neg = -1
		degree = int(parameters[i + 2][2:])
		if (degree != 0 and degree != 1 and degree != 2):
			print("The polynomial degree '{0}' is invalid".format(degree))
			exit()
		multiplicator = float(parameters[i])
		if (i > 0 and parameters[i - 1] == "-"):
			multiplicator = -multiplicator
		elems[degree] += multiplicator * neg
		i += 4
	if (elems[2] != 0):
		print("Reduced form: {2} * X^2 + {1} * X + {0} = 0".format(elems[0], elems[1], elems[2]))
	elif (elems[1] != 0):
		print("Reduced form: {1} * X + {0} = 0".format(elems[0], elems[1]))
	else:
		print("Reduced form: {0} = 0".format(elems[0]))
	return elems[::-1]

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
	elems = reduce_expression(" ".join(sys.argv[1:]).split())
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
