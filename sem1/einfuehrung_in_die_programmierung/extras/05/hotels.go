package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Hotel struct {
	name,
	distancetoNextHotel string
}

func main() {
	fmt.Println("Uebungsstunde 2 zu Go-Programmierung in Einführung in die Programmierung.")
	/* 4.11.2022, 08:45 Uhr. */

	var hotels [5]Hotel

	hotels[0] = Hotel{"Hotel Baltic Stralsund", "115km"}
	hotels[1] = Hotel{"Forsthaus Ecktannen", "112km"}
	hotels[2] = Hotel{"Hotel Haus Chorin", "38km"}
	hotels[3] = Hotel{"Comfort Hotel Bernau", "12km"}
	hotels[4] = Hotel{"Hotel Alt-Karow", "12km"}

	print("Hotelname: ")
	hotelFound := false
	hotelName, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	hotelName = strings.TrimSpace(hotelName)

	for _, element := range hotels {
		if element.name == hotelName {
			hotelFound = true
			fmt.Printf("The next Hotel is only %s away!", element.distancetoNextHotel)
			break
		}
	}
	if !hotelFound {
		fmt.Printf("The Hotel\"%s\"could not be found.", hotelName)
	}
}
