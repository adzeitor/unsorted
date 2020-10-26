package shunting_yard

// ToPostfix converts infix expressions to reverse polish notation.
func ToPostfix(expr []string) []string {
	var (
		output    stack
		operators stack
	)
	for _, e := range expr {
		switch {
		case isOperator(e):
			operators.Push(e)
		case e == "(":
			continue
		case e == ")":
			output.Push(operators.Pop())
		default:
			output.Push(e)
		}
	}
	// At the end of reading the expression,
	// pop all operators off the stack and onto the output.
	for !operators.IsEmpty() {
		output.Push(operators.Pop())
	}
	return output.Items()
}

func isOperator(s string) bool {
	switch s {
	case "+", "-", "*", "/":
		return true
	default:
		return false
	}
}
