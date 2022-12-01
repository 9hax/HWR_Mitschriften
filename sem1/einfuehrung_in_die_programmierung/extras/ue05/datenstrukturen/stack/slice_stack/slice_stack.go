package slice_stack

var stack = make([]string, 0)

func Push(s string) {
	stack = append(stack, s)
}

func Pop() string {
	var x string
	x, stack = stack[len(stack)-1], stack[:len(stack)-1]
	return x
}

func IsEmpty() bool {
	return len(stack) == 0
}
