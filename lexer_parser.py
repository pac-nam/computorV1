
def	create_matrice():
	dico = {}
	dico["X"] = "^0123456789"
	dico["^"] = "012"
	dico["0"] = "X0123456789+-*"
	dico["1"] = "X0123456789+-*"
	dico["2"] = "X0123456789+-*"
	dico["3"] = "X0123456789+-*"
	dico["4"] = "X0123456789+-*"
	dico["5"] = "X0123456789+-*"
	dico["6"] = "X0123456789+-*"
	dico["7"] = "X0123456789+-*"
	dico["8"] = "X0123456789+-*"
	dico["9"] = "X0123456789+-*"
	dico["+"] = "X0123456789"
	dico["-"] = "X0123456789"
	dico["*"] = "X0123456789"

def check_spaces(expression):
	i = 0
	valid = 1
	if (expression[0] not in "0123456789X+-"):
		print("wrong character {0}".format(expression[0]))
		return (0)
	while (i < len(expression)):
		if (expression[i] in "X0123456789" and i + 2 < len(expression) and expression[i + 1] == " " and expression[i + 2] in "0123456789"):
			return (0)
		# elif (expression[i] in "+-*" and i + 2 < len(expression) and expression[i + 1] == " " and expression[i + 2] in "-+*")
	return (1)

def	parse_expression_part(expression, elems)
	if (expression.count("X") > 1):
		print("can not have more than one 'X' in product")
		return (0)
	
	degree = int(parameters[i + 2][2:])
	if (degree != 0 and degree != 1 and degree != 2):
		print("The polynomial degree '{0}' is invalid".format(degree))
		exit()
	multiplicator = float(parameters[i])
	if (i > 0 and parameters[i - 1] == "-"):
		multiplicator = -multiplicator
	elems[degree] += multiplicator * neg
	i += 4
	return elems

def	parse(expression, matrice):
	if (!(check_spaces(expression))):
		return (0)
	expression = "".join(expression).split())
	elems = [0, 0, 0]
	i = 0
	start = 0
	while (i < len(expression)):
		if (i + 1 < len(expression) and expression[i + 1] not in matrice[expression[i]]):
			print("wrong character {0} after {1}".format(expression[i + 1], expression[i]))
			return (0)
		if (expression[i] in "+-"):
			if (!(elems = parse_expression_part(expression[start:i], elems)))
				return (0)
			start = i
		i += 1
	return elems


def lexer_parser(parameters):
	LR_parts = parameters.split("=")
	if (len(LR_parts) != 2):
		print("only one '=' is necessary")
		return (0)
	matrice = create_matrice()
	if (!(left = parse(LR_parts[0], matrice)) or !(right = parse(LR_parts[0], matrice))):
		return (0)
	left[0] -= right[0]
	left[1] -= right[1]
	left[2] -= right[2]
	return (left)