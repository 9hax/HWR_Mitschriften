package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Uebung 04 - Dateien")

	fmt.Println("Aufgabe 01 - Dateien schreiben")
	writeFiles()

	fmt.Println("Aufgabe 02 - Dateien auslesen")
	readFiles()

}

func createFiles() [7]*os.File {
	var myFiles [7]*os.File
	for a := range myFiles {
		myFiles[a], _ = os.Create("testFiles/FileIO." + strconv.Itoa(a+1) + ".txt")
	}
	return myFiles
}

func writeFiles() {
	fmt.Println()
	myFiles := createFiles()

	myFiles[0].Write([]byte("Hallo"))
	myFiles[1].Write([]byte("1,2,3,4,5,6,7,8,9,10"))

	myBools := [10]bool{true, true, false, true, false, false, false, false, false, false}
	var boolBytes [10]byte
	for i := range boolBytes {
		if myBools[i] {
			boolBytes[i] = 1
		} else {
			boolBytes[i] = 0
		}
	}
	myFiles[2].Write(boolBytes[:])

	myInts := [10]int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var intBytes [10]byte
	for i := range intBytes {
		intBytes[i] = byte(myInts[i])
	}
	myFiles[3].Write(intBytes[:])

	myFloats := [10]float64{1.01, 2.02, 3.03, 4.04, 5.05, 6.06, 7.07, 8.08, 9.09, 10.10}
	var floatBytes [10]byte
	for i := range floatBytes {
		floatBytes[i] = byte(myFloats[i])
	}
	myFiles[4].Write(floatBytes[:])

	var moreOfIntBytes [512]byte
	for i := range moreOfIntBytes {
		moreOfIntBytes[i] = byte(i - 256)
	}
	myFiles[5].Write(moreOfIntBytes[:])

	moreStrings := []string{
		"This is my first string.\n",
		"This is my second string.\n",
		"This is my third string.\n",
		"This is my fourth string.\n",
		"This is my fifth string.\n",
		"This is my sixth string.\n",
		"This is my seventh string.\n",
		"This is my eighth string.\n",
		"This is my ninth string.\n",
		"This is my tenth string.\n"}
	for i := range moreStrings {
		myFiles[6].Write([]byte(moreStrings[i]))
	}

	for _, file := range myFiles {
		file.Close()
	}
}

func readFiles() {
	intFile, _ := os.Open("testFiles/FileIO.2.txt")
	stats, _ := intFile.Stat()
	fileContent := make([]byte, stats.Size())
	_, _ = intFile.Read(fileContent)
	fmt.Println(string(fileContent))

	moreIntsFile, _ := os.Open("testFiles/FileIO.6.txt")
	stats, _ = moreIntsFile.Stat()
	fileContent = make([]byte, stats.Size())
	_, _ = moreIntsFile.Read(fileContent)
	var stringContent string
	for _, integer := range fileContent {
		stringContent += strconv.Itoa(int(integer)) + ", "
	}
	fmt.Println(stringContent)

}
