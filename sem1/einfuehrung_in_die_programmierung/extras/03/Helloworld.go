package main

import (
	"fmt"
	"math"
	"strconv"
	"unsafe"
)

func main() {
	fmt.Printf("Hello, World!\n")

	f := 3.14159265

	fmt.Printf(fmt.Sprint(f))

	var value int16 = 30000
	fmt.Println("The Value 30000 converted from int16 to int32 is: ", int32(value))

	var value2 int32 = 40000
	fmt.Println("The Value 40000 converted from int32 to int16 has an overflow and is: ", int16(value2))

	fmt.Println("\nTask 4")

	outputStringLength(" Hallo welt")
	outputStringLength("ÄÖÜäöü\\")

	fmt.Println("\nTask 5")

	fmt.Println(strconv.Atoi("31415"))
	fmt.Println(strconv.Itoa(312039186))
	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Println("The if statement ran, as the value passed to it was evaluated as true.")
	}
	c, _ := strconv.ParseFloat("3.1415112142341231412", 128)
	fmt.Println(c)

	d, _ := strconv.ParseInt("1249740186492", 10, 64)
	fmt.Println(d)

	e := []string{"123173829", "32341351753", "125834"}
	var sum int64 = 0
	for _, element := range e {
		fmt.Println(element)
		temp, _ := strconv.ParseInt(element, 10, 64)
		sum += temp
	}
	fmt.Println("The grand total sum is: ", sum)
	var test float64 = 123.123141519
	fmt.Println(strconv.FormatBool(true))
	fmt.Println(strconv.FormatFloat(test, 'E', -1, 32))
	fmt.Println(strconv.FormatInt(-123, 16))
	fmt.Println(strconv.FormatUint(134151, 16))

	fmt.Println("Zusatz1")

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, element := range numbers {
		println(element, " ", element*element)
	}

	fmt.Println("Zusatz2")

	var y int
	println("The size of an int is ", unsafe.Sizeof(y), " bytes long.")

	if 3 == int8(math.Sqrt(3)*math.Sqrt(3)) {
		println("int32 is good!")
	}
	println((math.Sqrt(3) * math.Sqrt(3)))
}

func outputStringLength(output string) {
	fmt.Println(output, " has a length of ", len(output))
}
