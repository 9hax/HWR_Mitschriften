package main

import (
	"fmt"
)

func main() {
	fmt.Println("For Increment test")

	var Liste [10]int = [10]int{10, 9, 8, 7, 5, 6, 4, 3, 1, 2}

	for _, element := range Liste {
		fmt.Println(element)
		element++
	}

	for _, element := range Liste {
		fmt.Println(element)
	}

}
