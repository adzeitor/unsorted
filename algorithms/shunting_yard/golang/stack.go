package shunting_yard

type stack struct {
	values []string
}

func (stack *stack) Pop() string {
	top := stack.values[len(stack.values)-1]
	stack.values = stack.values[:len(stack.values)-1]
	return top
}

func (stack *stack) Push(value string)  {
	stack.values = append(stack.values, value)
}

func (stack *stack) IsEmpty() bool {
	return len(stack.values) == 0
}

func (stack *stack) Items() []string {
	return stack.values
}
