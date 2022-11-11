package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Uebung 03 - Funktionen und Prozeduren")

	fmt.Println("Printing Multiplication table.") // Task 1
	printTable(makeMultiTable(12, 12))

	a := 4
	b := 7
	fmt.Println("Printing bigger value:")
	printBigger(a, b) // Task 2

	fmt.Println("The bigger return value was", returnBigger(a, b)) // Task 3

	a, b = swapValues(a, b)
	fmt.Println("The first return value was", a, "and the second return value was", b, ".") // Task 4

	fmt.Println("The iter fac of 3, 4, 5 are:", facIter(3), facIter(4), facIter(5))             // Task 5
	fmt.Println("The recurse fac of 3, 4, 5 are:", facRecurse(3), facRecurse(4), facRecurse(5)) // Task 6

	fmt.Println("The recursive function is slower than the Iterative function by a lot, since it needs to do stack operations and jump between different codepaths.") // Task 7

	fmt.Println("The iter fib of 3, 4, 5 is:", fibRecurse(3), fibRecurse(4), fibIter(5))       // Task 8
	fmt.Println("The recurse fib of 3, 4, 5 is:", fibRecurse(3), fibRecurse(4), fibRecurse(5)) // Task 8

	fmt.Println("The recursive function is slower than the Iterative function by a lot, since it needs to do stack operations and jump between different codepaths.") // Task 7

	fmt.Println("The Quadratic numbers up to 10:", genQuadNumbers(10))
	fmt.Println("Defining 100.000 Structs with quadratic numbers, this might take a while.")
	nums := genQuadNumbers(1000000)

	fmt.Println("Done! Writing them to a file 100 times...")
	var errors error
	for i := 0; i < 100; i++ {
		errors = os.WriteFile("quadNumsTo100000.txt", []byte(fmt.Sprintln(nums)), 0644)
	}
	fmt.Println("If an error occured, it will be displayed here:", errors)

	fmt.Println("That was rather slow! This is because the array was copied 100 times and then destroyed 100 times. Formatting the array into a better format for the application is faster.")
}

func makeMultiTable(sizex, sizey int) [][]int {
	a := make([][]int, sizey)

	for y := range a {
		a[y] = make([]int, sizex)
		for x := range a[y] {
			a[y][x] = (x + 1) * (y + 1)
		}
	}
	return a
}

func printTable(table [][]int) {
	for _, row := range table {
		for _, col := range row {
			fmt.Printf("%*d", 4, col)
		}
		fmt.Print("\n")
	}
}

func printBigger(a, b int) { // Task 2: Print the bigger value of the two input integers to the screen.
	if a < b {
		fmt.Println(b)
	} else {
		fmt.Println(a)
	}
}

func returnBigger(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func swapValues(a, b int) (int, int) {
	return b, a
}

func facIter(limit int) int {
	facValue := 1
	for i := limit; i > 0; i-- {
		facValue *= i
	}
	return facValue
}

func facRecurse(limit int) int {
	if limit < 2 {
		return 1
	}
	return facRecurse(limit-1) * limit
}

func fibIter(limit int) int {
	total, previous := 1, 0
	for limit > 0 {
		temp := total
		total += previous
		previous = temp
		limit--
	}
	return total
}

func fibRecurse(limit int) int {
	if limit <= 1 {
		return 1
	}
	return fibRecurse(limit-1) + fibRecurse(limit-2)

}

type quadNumber struct {
	base, quad int
}

func genQuadNumbers(limit int) []quadNumber {
	var quadNumbers []quadNumber
	for i := 0; i < limit; i++ {
		quadNumbers = append(quadNumbers, quadNumber{i, i * i})
	}
	return quadNumbers
}
