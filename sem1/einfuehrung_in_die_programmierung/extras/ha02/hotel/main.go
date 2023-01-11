package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"hotel/imaging"
	"image/color"
	"log"
	"os"
	"strconv"
	"strings"
)

var datasetFile string = "dataset_tiny.csv"

type Item struct {
	name        string
	description string
}

type RouteItem struct {
	item         Item
	nextID       uint64
	nextDistance uint64
}

var POIs []RouteItem

func main() {
	log.Print("Initializing Program. Please hold tight, this might take a while depending on the dataset size.")
	// This program has been tested with an AI generated dataset with 1000 POIs.
	// To decrease unneeded Data, the Dataset has been truncated to 5 POIs.
	//Uncomment the following line to load the big dataset.

	//datasetFile = "dataset.csv"

	datasetRecords := readDataset(datasetFile)
	POIs = parseDataset(datasetRecords)

	log.Print("Finished preparing dataset.")

	POIStart := requestRoutePOIIndex(1)
	POIEnd := requestRoutePOIIndex(2)
	//POIStart, POIEnd := 2, 4

	distance, waypoints := findShortestLink(POIStart, POIEnd)
	println("Going from POI", POIStart, "to POI", POIEnd, "is fastest when traveling through the following POIs:")
	fmt.Printf("%+v\n", waypoints)
	println("Going this way will be", distance, "kilometres.")

	drawRoute(POIStart, POIEnd, waypoints, distance)
}

func requestRoutePOIIndex(promptType int) int {
	if promptType == 1 {
		print("Please Enter the Name of the POI that you want to start your trip from: >")
	} else if promptType == 2 {
		print("Please Enter the Name of the POI you want to get to: >")
	} else {
		print("Please Enter the Name of a POI. >")
	}

	POIName, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	POIName = strings.TrimSpace(POIName)

	for id, poi := range POIs {
		if poi.item.name == POIName {
			return id
		}
	}
	fmt.Printf("The POI\"%s\"could not be found. Assuming POI id 0.", POIName)
	return 0
}

func drawRoute(start, end int, waypoints []RouteItem, distance uint64) {

	imaging.AddColorToPalette(color.RGBA{0x00, 0x00, 0xff, 0x80})
	imaging.AddColorToPalette(color.RGBA{0x60, 0x60, 0x60, 0xff})

	outputImage := imaging.CreateBasicGif(200, 200)

	// POI Icons
	imaging.DrawPOI(outputImage, 50, 25, 1)
	imaging.DrawPOI(outputImage, 100, 140, 1)

	// Waypoint Iconography
	imaging.DrawLine(outputImage, 100, 50, 130, 70, 9)
	imaging.DrawLine(outputImage, 130, 70, 50, 120, 9)
	imaging.DrawLine(outputImage, 50, 120, 95, 160, 9)

	// Waypoints
	imaging.DrawString(outputImage, 101, 95, "Traversing over", 5)
	imaging.DrawString(outputImage, 100, 94, "Traversing over", 1)

	imaging.DrawString(outputImage, 101, 103, strconv.Itoa(len(waypoints)-2)+" other POIs.", 5)
	imaging.DrawString(outputImage, 100, 102, strconv.Itoa(len(waypoints)-2)+" other POIs.", 1)

	imaging.DrawString(outputImage, 101, 111, "Distance: "+strconv.FormatUint(distance, 10)+"km", 10)
	imaging.DrawString(outputImage, 100, 110, "Distance: "+strconv.FormatUint(distance, 10)+"km", 1)

	// Start and End Names
	imaging.DrawString(outputImage, 101, 16, POIs[start].item.name, 10)
	imaging.DrawString(outputImage, 100, 15, POIs[start].item.name, 1)
	imaging.DrawString(outputImage, 101, 186, POIs[end].item.name, 10)
	imaging.DrawString(outputImage, 100, 185, POIs[end].item.name, 1)

	// Finalize GIF Image
	imaging.WriteGifImage(outputImage, "PoiInfo")
}

func readDataset(datasetFileName string) [][]string {
	log.Print("Loading ", datasetFile, ".")
	f, err := os.Open(datasetFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return records
}

func parseDataset(dataset [][]string) []RouteItem {
	POIs := make([]RouteItem, 0)
	for index, record := range dataset {
		var nextPOIDistance uint64
		var loadingError error
		if len(record) == 3 {
			nextPOIDistance, loadingError = strconv.ParseUint(record[2], 10, 64)
			if loadingError != nil {
				log.Fatal("DATASET LOADING ERROR! Cannot parse nextPOIDistance for POI on line "+strconv.Itoa(index)+".\n", loadingError)
			}
			POIs = append(POIs, RouteItem{Item{record[0], record[1]}, uint64(index + 1), nextPOIDistance})
		} else {
			nextPOIDistance, loadingError = strconv.ParseUint(record[1], 10, 64)
			if loadingError != nil {
				log.Fatal("DATASET LOADING ERROR! Cannot parse nextPOIDistance for POI on line "+strconv.Itoa(index)+".\n", loadingError)
			}
			POIs = append(POIs, RouteItem{Item{record[0], "No description provided by dataset."}, uint64(index + 1), nextPOIDistance})
		}
	}
	POIs[len(POIs)-1].nextID = 0
	return POIs
}

func findShortestLink(POIIndexA, POIIndexB int) (uint64, []RouteItem) {
	if POIIndexA == POIIndexB {
		// When traveling to the same POI, the distance is always 0.
		return 0, make([]RouteItem, 0)
	}

	swapped := false
	if POIIndexA > POIIndexB {
		// This makes the program always iterate the POIs in a positive direction.
		// Since the list of POIs is not a double linked list (one POI only knows the distance to the next POI), this makes calculations easier.
		POIIndexA, POIIndexB = POIIndexB, POIIndexA
		swapped = true
	}

	//println("Calculating Distance between POIs at index", POIIndexA, "and", POIIndexB)
	//fmt.Printf("%+v\n", POIs[POIIndexA])
	//fmt.Printf("%+v\n", POIs[POIIndexB])

	// Calculate the distance by iterating from the lower Index to the higher index.
	var distanceUp uint64 = 0
	POIsUp := make([]RouteItem, 0)
	for i := POIIndexA; i <= POIIndexB-1; i++ {
		distanceUp += POIs[i].nextDistance
		POIsUp = append(POIsUp, POIs[i])
	}
	POIsUp = append(POIsUp, POIs[POIIndexB])

	// Calculate the distance by iterating from the higher index to the end of the list,
	// then resume by iterating from the first element to the lower index.
	// This by defualt returns a swapped result in the list of waypoints.
	var distanceOverZero uint64 = 0
	POIsOverZero := make([]RouteItem, 0)
	for i := POIIndexB; i <= len(POIs)-1; i++ {
		// Iterate from higher index to end
		distanceOverZero += POIs[i].nextDistance
		POIsOverZero = append(POIsOverZero, POIs[i])
	}
	for i := 0; i < POIIndexA; i++ {
		// Iterate from 0 to lower index
		distanceOverZero += POIs[i].nextDistance
		POIsOverZero = append(POIsOverZero, POIs[i])
	}
	POIsOverZero = append(POIsOverZero, POIs[POIIndexA])

	if distanceUp < distanceOverZero {
		if swapped {
			// Reverse the list of traversed POIs if traveling from a high POI index to a lower one.
			for i, j := 0, len(POIsUp)-1; i < j; i, j = i+1, j-1 {
				POIsUp[i], POIsUp[j] = POIsUp[j], POIsUp[i]
			}
		}
		return distanceUp, POIsUp
	}

	if !swapped {
		// Reverse the list of traversed POIs if traveling from a high POI index to a lower one.
		for i, j := 0, len(POIsOverZero)-1; i < j; i, j = i+1, j-1 {
			POIsOverZero[i], POIsOverZero[j] = POIsOverZero[j], POIsOverZero[i]
		}
	}
	return distanceOverZero, POIsOverZero
}
