package scheme

func addBultin(env Environment) {
	env["+"] = Builtin(plusBuiltin)
	env["-"] = Builtin(minusBuiltin)
	env["*"] = Builtin(multBuiltin)
}

func plusBuiltin(args []ExprT, env Environment) (ExprT, Environment) {
	n1, _ := eval(args[0], env)
	n2, _ := eval(args[1], env)
	return Int(int(n1.(IntT)) + int(n2.(IntT))), env
}

func minusBuiltin(args []ExprT, env Environment) (ExprT, Environment) {
	n1, _ := eval(args[0], env)
	n2, _ := eval(args[1], env)
	return Int(int(n1.(IntT)) - int(n2.(IntT))), env
}

func multBuiltin(args []ExprT, env Environment) (ExprT, Environment) {
	n1, _ := eval(args[0], env)
	n2, _ := eval(args[1], env)
	return Int(int(n1.(IntT)) * int(n2.(IntT))), env
}
