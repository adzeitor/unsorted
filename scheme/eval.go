package scheme

import "fmt"

var globalEnvironment = DefaultEnvironment()

type Builtin func(args []ExprT, env Environment) (ExprT, Environment)

// Eval evaluate expression in global environment. Be careful with several
// instances, in this case use EvalInEnvironment.
func Eval(s string) ExprT {
	result, newEnvironment := EvalInEnvironment(s, globalEnvironment)
	globalEnvironment = newEnvironment
	return result
}

func EvalInEnvironment(s string, env Environment) (result ExprT, resultEnv Environment) {
	// FIXME: is exceptions here good idea? Maybe move to repl or exception free code?
	defer func() {
		err := recover()
		if err != nil {
			result = StringT(fmt.Sprintf("exception: %v", err))
			// FIXME: good?
			resultEnv = env
		}
	}()

	parsed, _, ok := Parse(s)
	if !ok {
		return StringT("parse error"), env
	}
	return eval(parsed, env)
}

type Lambda struct {
	Env        Environment
	Parameters []ExprT
	Body       ExprT
}

func (lambda Lambda) MakeArgEnv(arguments []ExprT) Environment {
	env := make(Environment)
	for i, value := range arguments {
		// FIXME: parametrs out of index
		env[lambda.Parameters[i].(SymbolT)] = value
	}
	return env
}

func applyLambda(lambda Lambda, arguments []ExprT, env Environment) (ExprT, Environment) {
	closureEnv := env.Extend(lambda.Env)
	closureEnv = closureEnv.Extend(lambda.MakeArgEnv(arguments))
	return eval(lambda.Body, closureEnv)
}

func evalArguments(args []ExprT, env Environment) ([]ExprT, Environment) {
	evaledArgs := make([]ExprT, len(args))
	for i, arg := range args {
		evaledArgs[i], env = eval(arg, env)
	}
	return evaledArgs, env
}

func evalList(list ListT, env Environment) (ExprT, Environment) {
	if len(list) == 0 {
		return List(), env
	}
	head := list[0]

	// FIXME: add cons?
	switch head {
	case SymbolT("quote"):
		return list[1], env
	case SymbolT("="):
		value1, _ := eval(list[1], env)
		value2, _ := eval(list[2], env)
		// FIXME: introduce BoolT type
		// FIXME: recursive lists comparison?
		return value1 == value2, env
	case SymbolT("null?"):
		arg, env := eval(list[1], env)
		if l, ok := arg.(ListT); ok {
			return len(l) == 0, env
		}
		return false, env
	case SymbolT("if"):
		condition, _ := eval(list[1], env)
		if condition.(bool) {
			return eval(list[2], env)
		} else {
			return eval(list[3], env)
		}
	case SymbolT("define"):
		name := list[1].(SymbolT)
		value, _ := eval(list[2], env)
		env[name] = value
		return name, env
	case SymbolT("cons"):
		car, env := eval(list[1], env)
		cdr, env := eval(list[2], env)
		// FIXME: optmize and allocate once
		l := List(car)
		l = append(l, cdr.(ListT)...)
		return l, env
	case SymbolT("cdr"):
		arg, env := eval(list[1], env)
		// FIXME: panic and copy?
		return arg.(ListT)[1:], env
	case SymbolT("car"):
		arg, env := eval(list[1], env)
		// FIXME: panic and copy?
		return arg.(ListT)[0], env
	case SymbolT("lambda"):
		return Lambda{
			Env:        env.Copy(),
			Parameters: list[1].(ListT),
			Body:       list[2],
		}, env
	}

	head, env = eval(head, env)
	switch head.(type) {
	case Builtin:
		return head.(Builtin)(list[1:], env)
	case Lambda:
		arguments, env := evalArguments(list[1:], env)
		return applyLambda(head.(Lambda), arguments, env)
	}
	return nil, env
}

type Environment map[SymbolT]ExprT

func (env Environment) Copy() Environment {
	copied := make(Environment, len(env))
	for k, v := range env {
		copied[k] = v
	}
	return copied
}

func (env Environment) Extend(extension Environment) Environment {
	newEnv := env.Copy()
	// add extension to new environment
	for k, v := range extension {
		newEnv[k] = v
	}
	return newEnv
}

func DefaultEnvironment() Environment {
	env := make(Environment)
	addBultin(env)
	return env
}

func eval(expr ExprT, env Environment) (ExprT, Environment) {
	switch value := expr.(type) {
	case IntT:
		return value, env
	case StringT:
		return value, env
	case SymbolT:
		// FIXME: what if not defined?
		if v := env[value]; v != nil {
			return v, env
		}
	case ListT:
		return evalList(value, env)
	}
	return nil, env
}
