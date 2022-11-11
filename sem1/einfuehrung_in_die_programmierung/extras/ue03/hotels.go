package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Hotel struct {
	name                              string
	nextHotelIndex, nextHotelDistance int
}

var hotels [5]Hotel

func main() {
	fmt.Println("HotelFinder2000")

	fmt.Println("Defining Hotels...")

	genHotels()

	distance := 0

	for _, hotel := range hotels {
		distance += hotel.nextHotelDistance
	}
	fmt.Println("The distance between all hotels is", distance, "km.")

	distance = 0
	for i := 4; i >= 0; i-- {
		distance += hotels[i].nextHotelDistance
	}
	fmt.Println("The distance between all hotels is", distance, "km.")

	fmt.Println("The distance between hotel 2 and 5 is ", findShortestDistance(1, 4), "km.")
	fmt.Println("The distance between hotel 5 and 2 is ", findShortestDistance(4, 1), "km.")

	hotel1 := getHotelIndexFromName("Input the name of the first Hotel. >")
	hotel2 := getHotelIndexFromName("Input the name of the second Hotel. >")

	if hotel1 < 0 || hotel2 < 0 {
		fmt.Println("At least one of the hotels could not be found.")
	} else {
		fmt.Println("The distance between hotel 5 and 2 is ", findShortestDistance(hotel1, hotel2), "km.")
	}

}

func genHotels() {
	hotels[0] = Hotel{"Alpenrose", 1, 11}
	hotels[1] = Hotel{"Budget-Hotel", 2, 5}
	hotels[2] = Hotel{"Rabenstein", 3, 11}
	hotels[3] = Hotel{"Hohenstein", 4, 5}
	hotels[4] = Hotel{"Grand-Hotel", -1, 0}
}

func findShortestDistance(hotelIndexA, hotelIndexB int) int {
	if hotelIndexA == hotelIndexB {
		return 0
	}
	distance := 0
	if hotelIndexA > hotelIndexB {
		hotelIndexA, hotelIndexB = hotelIndexB, hotelIndexA
	}
	for i := hotelIndexA; i <= hotelIndexB; i++ {
		distance += hotels[i].nextHotelDistance
	}
	return distance
}

func getHotelIndexFromName(prompt string) int {
	fmt.Print(prompt)
	hotelName, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	hotelName = strings.TrimSpace(hotelName)
	for index, hotel := range hotels {
		if hotelName == hotel.name {
			return index
		}
	}
	return -1
}
