package array_stack

import "errors"

var stack [10]string
var stackPointer int = -1

func Push(s string) error {
	if stackPointer > len(stack) {
		return errors.New("stack is full")
	}
	stackPointer++
	stack[stackPointer] = s
	return nil
}

func Pop() (string, error) {
	if stackPointer < 0 {
		return "", errors.New("stack is empty2")
	}
	element := stack[stackPointer]
	stackPointer--
	return element, nil
}

func IsEmpty() bool {
	return stackPointer < 0
}
