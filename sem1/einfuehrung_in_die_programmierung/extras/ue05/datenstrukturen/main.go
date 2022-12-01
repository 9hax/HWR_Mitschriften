package main

import (
	"datenstrukturen/queue/slice_queue"
	"datenstrukturen/stack/array_stack"
	"datenstrukturen/stack/slice_stack"
)

func main() {
	println("Programm für die Nutzung von abgefahrenen Datenstrukturen")

	use_slice_stack()
	use_array_stack()
	use_slice_queue()

}

func use_slice_stack() {
	if slice_stack.IsEmpty() {
		println("The stack is currently empty.")
	}

	slice_stack.Push("Hello World!")
	slice_stack.Push("This is a second element in the stack.")

	if !slice_stack.IsEmpty() {
		println("The stack is currently not empty.")
	}

	println(slice_stack.Pop())
	println(slice_stack.Pop())
}

func use_array_stack() {
	if array_stack.IsEmpty() {
		println("The stack is currently empty.")
	}

	array_stack.Push("Hello World!")
	array_stack.Push("This is a second element in the stack.")

	if !array_stack.IsEmpty() {
		println("The stack is currently not empty.")
	}

	element, err := array_stack.Pop()
	if err == nil {
		println(element)
	}
	element, err = array_stack.Pop()
	if err == nil {
		println(element)
	}
}

func use_slice_queue() {
	if slice_queue.IsEmpty() {
		println("The stack is currently empty.")
	}

	slice_queue.Push("Hello World!")
	slice_queue.Push("This is a second element in the stack.")

	if !slice_queue.IsEmpty() {
		println("The stack is currently not empty.")
	}

	println(slice_queue.Pop())
	println(slice_queue.Pop())
}
