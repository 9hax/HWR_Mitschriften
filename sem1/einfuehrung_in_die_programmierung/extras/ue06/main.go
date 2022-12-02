package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Uebung 06 - Stacking Pointer Magic")

	myIntPointer := new(int)
	*myIntPointer = 7
	fmt.Println(*myIntPointer)

	myRunePointer := new(rune)
	*myRunePointer = 'A'
	fmt.Printf("%s\n", string(*myRunePointer))

	// Wenn ein Pointer "In den Wald" zeigt, dann hat er einen Wert, welcher außerhalb des Speichers liegt, auf den Zugegriffen werden kann.

	myRunePointer = nil
	//myRuneVariable := *myRunePointer // This causes a panic because myRunePointer points to nil.
	//fmt.Println(myRuneVariable)

	var pointers [100000]*[1048576]int
	for i := 0; i < len(pointers); i++ {
		var bigValues [1048576]int
		for j := 0; j < len(bigValues); j++ {
			bigValues[j] = j
		}
		pointers[i] = &bigValues
	}

	fmt.Print("Press Enter to end program.")
	bufio.NewScanner(os.Stdin).Scan()

	for i := 0; i < len(pointers); i++ {
		fmt.Printf("%p", pointers[i])
	}
}
