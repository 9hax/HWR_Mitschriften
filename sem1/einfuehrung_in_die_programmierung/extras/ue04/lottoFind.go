package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type LottoNumber struct {
	Date    time.Time
	numbers [6]int
	super   int
}

func main() {
	fmt.Println("Übung 04 - Dateien\nAufgabe 4")

	lottoNumbers := readLottoFile()
	//fmt.Println(lottoNumbers)

	y, m, d := readDate()

	printNumbers(y, m, d, lottoNumbers)
}

func validateNumber(number LottoNumber, nano int64) bool {
	rand.Seed(nano)
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
	if numbers == number.numbers && superNum == number.super {
		return true
	}
	fmt.Println("Invalid Number! Number from ", number.Date.String(), ":", number.numbers, "should be", numbers, "and the super", number.super, "should be", superNum, "!")
	return false
}

func contains(array [6]int, value int) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}
	return false
}

func readLottoFile() []LottoNumber {
	csvfile, error := os.OpenFile("lottoNumbers.txt", os.O_RDONLY, 0600)
	if error != nil {
		fmt.Println("Couldn't open lottoNumbers.txt. Error:", error)
	}
	defer csvfile.Close()

	reader := csv.NewReader(csvfile)

	var lottoNumbers []LottoNumber
	for {
		record, err := reader.Read()
		if err == nil {
			timeValue, _ := time.Parse(time.RFC3339Nano, record[0])
			unixNano, _ := strconv.Atoi(record[1])
			var numbers [6]int
			for i := 0; i < 6; i++ {
				numbers[i], _ = strconv.Atoi(record[i+2])
			}
			supernum, _ := strconv.Atoi(record[8])
			newNumber := LottoNumber{
				Date:    timeValue,
				numbers: numbers,
				super:   supernum}
			if validateNumber(newNumber, int64(unixNano)) {
				lottoNumbers = append(lottoNumbers, newNumber)
			}
		} else {
			break
		}
	}
	return lottoNumbers
}

func readDate() (int, int, int) {
	fmt.Print("Enter A Date with format yyyy-MM-dd: ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	data := strings.Split(text[0:10], "-")
	year, _ := strconv.Atoi(data[0])
	month, _ := strconv.Atoi(data[1])
	day, _ := strconv.Atoi(data[2])

	return year, month, day
}

func printNumbers(y int, m int, d int, lottoNumbers []LottoNumber) {
	for _, number := range lottoNumbers {
		ny, nm, nd := number.Date.Date()
		nmi := int(nm)
		if ny == y && nmi == m && nd == d {
			fmt.Println(number)
		}
	}
}
