package array_queue

var queue [10]string
var queueWrite int
var queueRead int

func init() {
	queueRead, queueWrite = 0, 0
}

func Push(s string) {
}

/*func Pop() string {
	var x string
	x, queue = queue[0], queue[1:]
	return x
}*/

func IsEmpty() bool {
	return len(queue) == 0
}
