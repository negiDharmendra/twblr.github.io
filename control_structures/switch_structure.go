package control_structures

const add string = "+"
const subtract string = "-"
const multiply string = "*"
const divide = "/"

func calculate(op string, arg1, arg2 float32) float32 {
	switch {
	case op == add:
		return arg1 + arg2
	case op == subtract:
		return arg1 - arg2
	case op == multiply:
		return arg1 * arg2
	case op == divide:
		return arg1 / arg2
	}
	return 0
}
