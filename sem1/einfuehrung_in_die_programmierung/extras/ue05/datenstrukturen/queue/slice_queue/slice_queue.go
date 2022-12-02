package slice_queue

var queue = make([]string, 0)

func Push(s string) {
	queue = append(queue, s)
}

func Pop() string {
	var x string
	x, queue = queue[0], queue[1:]
	return x
}

func IsEmpty() bool {
	return len(queue) == 0
}

func Clear() {
	queue = make([]string, 0)
}
