package scheme

type ExprT interface{}

type IntT int

type SymbolT string

type StringT string

type ListT []ExprT

func Int(n int) ExprT {
	return IntT(n)
}

func List(l ...ExprT) ListT {
	return l
}
