package main

import (
	"fmt"
)

func main() {
	println("Uebungsstunde 2 zu Go-Programmierung in Einführung in die Programmierung.")
	/* 4.11.2022, 08:45 Uhr. */

	var tenElementsArray [10]int

	for i := 0; i < 10; i++ {
		tenElementsArray[i] = i * i
	}

	fmt.Println(tenElementsArray)

	/* This would fail:*/
	//println(tenElementsArray[10])
	/* When using a competent IDE with language formatting and debugging tools, you will be told that this will cause an out-of-bounds exception while coding.*/

	println("Print the array using range")
	for _, element := range tenElementsArray {
		println(element)
	}

	println("Print the array using iterator")
	for i := 0; i < 10; i++ {
		println(tenElementsArray[i])
	}

	println("Print the array using for-break")
	i := 0
	for {
		println(tenElementsArray[i])
		i++
		if i > 9 {
			break
		}
	}
	i = 0

	println("Print the array using headloop")
	i = 0
	for i < 10 {
		println(tenElementsArray[i])
		i++
	}

}
