package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type LottoNumber struct {
	Timestamp time.Time
	numbers   [6]int
	super     int
}

func main() {
	fmt.Println("Übung 04 - Dateien\nAufgabe 3")

	lottoNumbers := generateLottoNumbers(3)
	for _, n := range lottoNumbers {
		fmt.Println(n)
	}

	writeToLottoFile(lottoNumbers)
}

func generateLottoNumbers(amount int) []LottoNumber {
	var numbers []LottoNumber
	for amount > 0 {
		numbers = append(numbers, generateNumberSet())
		amount--

	}
	return numbers
}

func generateNumberSet() LottoNumber {
	drawTime := time.Now()
	rand.Seed(drawTime.UnixNano())
	var numbers [6]int
	var superNum int = 0

	for i := 0; i < 7; i++ {
		myNum := rand.Intn(49)
		if !contains(numbers, myNum) {
			if i == 6 {
				superNum = myNum
				break
			}
			numbers[i] = myNum
		} else {
			i--
		}
	}
	return LottoNumber{
		Timestamp: drawTime,
		numbers:   numbers,
		super:     superNum,
	}
}

func contains(array [6]int, value int) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}
	return false
}

func writeToLottoFile(numbers []LottoNumber) {
	csvfile, error := os.OpenFile("lottoNumbers.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if error != nil {
		fmt.Println("Couldn't open lottoNumbers.txt. Error:", error)
	}

	writer := csv.NewWriter(csvfile)

	for _, number := range numbers {
		writer.Write(
			[]string{
				number.Timestamp.Format(time.RFC3339Nano),
				fmt.Sprint(number.Timestamp.UnixNano()),
				fmt.Sprint(number.numbers[0]),
				fmt.Sprint(number.numbers[1]),
				fmt.Sprint(number.numbers[2]),
				fmt.Sprint(number.numbers[3]),
				fmt.Sprint(number.numbers[4]),
				fmt.Sprint(number.numbers[5]),
				fmt.Sprint(number.super)})
	}
	writer.Flush()
	csvfile.Close()
}
