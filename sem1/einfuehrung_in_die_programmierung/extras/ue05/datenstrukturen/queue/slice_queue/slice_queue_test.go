package slice_queue

import (
	"testing"
)

func TestPushPop(t *testing.T) {
	if !IsEmpty() {
		Clear()
	}
	if !IsEmpty() {
		t.Error("Queue is supposed to be empty.")
	}
	// Set up data
	Push("A")
	Push("B")
	firstReturn := Pop()
	if firstReturn != "A" {
		t.Error("Queue was supposed to return A, but was", firstReturn)
	}
	secondReturn := Pop()
	if secondReturn != "B" {
		t.Error("Queue was supposed to return B, but was", secondReturn)
	}
	if !IsEmpty() {
		t.Error("Queue is supposed to be empty at end of test.")
	}
}
