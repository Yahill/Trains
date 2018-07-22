package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"time"
)

type Route struct {
	DepartureStationId int
	ArrivalStationId   int
}

type Trains struct {
	XMLName xml.Name `xml:"TrainLegs"`
	Train   []Train  `xml:"TrainLeg"`
}

type Train struct {
	TrainId             int     `xml:"TrainId,attr"`
	DepartureStationId  int     `xml:"DepartureStationId,attr"`
	ArrivalStationId    int     `xml:"ArrivalStationId,attr"`
	Price               float64 `xml:"Price,attr"`
	ArrivalTimeString   string  `xml:"ArrivalTimeString,attr"`
	DepartureTimeString string  `xml:"DepartureTimeString,attr"`
}

func main() {

	//route which will contain ready route for the passanger
	var route []int
	//route from departurte station to all other
	allRouteFromDepartureStation := []int{}

	//flags for options
	departureStation := flag.Int("Departure", 1, "Enter departure station")
	arrivalStation := flag.Int("Arrival", 1, "Enter arrival station")
	option := flag.String("option", "nill", "Choose option: cheapest ot fastest.")
	flag.Parse()

	//reading xml file
	trains := ReadXML()

	//nodes graph
	nodes := CreateGraph(trains)

	//count whole route from graph
	DeapthFirstSearch(*departureStation, nodes, func(node int) {
		allRouteFromDepartureStation = append(allRouteFromDepartureStation, node)
	})

	//create route to the arrival station
	for i := range allRouteFromDepartureStation {
		if allRouteFromDepartureStation[i] != *arrivalStation {
			buff := allRouteFromDepartureStation[i]
			route = append(route, buff)
		} else {
			buff := allRouteFromDepartureStation[i]
			route = append(route, buff)
			break
		}
	}

	//handling flag options
	switch *option {

	case "nill":
		fmt.Println("Choose option: -option=cheapest or -option=fastest.")

	case "cheapest":
		resultTrains := CheapestOption(route, trains)
		fmt.Println(route)
		for i := range resultTrains {
			fmt.Println("TrainID: " + strconv.Itoa(resultTrains[i].TrainId) + "\n" + "DepartureStationId: " + strconv.Itoa(resultTrains[i].DepartureStationId) + "\n" +
				"ArrivalStationId: " + strconv.Itoa(resultTrains[i].ArrivalStationId) + "\n" + "Price: " + strconv.FormatFloat(resultTrains[i].Price, 'f', 2, 32) + "\n" +
				"ArrivalTimeString: " + resultTrains[i].ArrivalTimeString + "\n" + "DepartureTimeString: " + resultTrains[i].DepartureTimeString + "\n")
		}

	case "fastest":
		resultTrains := FastestOption(route, trains)
		fmt.Println(route)
		for i := range resultTrains {
			fmt.Println("TrainID: " + strconv.Itoa(resultTrains[i].TrainId) + "\n" + "DepartureStationId: " + strconv.Itoa(resultTrains[i].DepartureStationId) + "\n" +
				"ArrivalStationId: " + strconv.Itoa(resultTrains[i].ArrivalStationId) + "\n" + "Price: " + strconv.FormatFloat(resultTrains[i].Price, 'f', 2, 32) + "\n" +
				"ArrivalTimeString: " + resultTrains[i].ArrivalTimeString + "\n" + "DepartureTimeString: " + resultTrains[i].DepartureTimeString + "\n")
		}
	}
}

func ReadXML() (trains Trains) {
	//reading xml file

	//open file
	xmlFile, err := os.Open("data.xml")
	//handle error if it happens
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Your xml file - successfully opened.")
	//closing xml
	defer xmlFile.Close()
	//read xmlFile as a byte array
	byteValue, _ := ioutil.ReadAll(xmlFile)

	//initializing of the Trains array
	//unmarshalling byteArray
	err = xml.Unmarshal(byteValue, &trains)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	return
}

func CreateGraph(trains Trains) map[int][]int {
	//searching for unic stations

	//map with visited stations
	encountered := map[Route]bool{}
	//structure with arrivalStationId and departureStationId
	var allStations []Route
	//structure with all unic stations
	var result []Route
	//create map
	var graph map[int][]int
	graph = make(map[int][]int)

	//create first structure with all the stations
	for i := 0; i < len(trains.Train); i++ {
		var buff Route

		buff.DepartureStationId = trains.Train[i].DepartureStationId
		buff.ArrivalStationId = trains.Train[i].ArrivalStationId

		allStations = append(allStations, buff)
	}

	//adding slice with unic routes
	for v := range allStations {
		if encountered[allStations[v]] == true {
			//do not add duplicate
		} else {
			//record this element to encountered
			encountered[allStations[v]] = true
			buff := allStations[v]
			//append to result slice
			result = append(result, buff)
		}
	}

	//create graph from unic stations
	for i := range result {
		graph[result[i].DepartureStationId] = append(graph[result[i].DepartureStationId], result[i].ArrivalStationId)
	}

	return graph
}

func DeapthFirstSearch(node int, nodes map[int][]int, fn func(int)) {
	//recursive with deapth-first search
	Recursive(node, map[int]bool{}, fn, nodes)
}

func Recursive(node int, currentStation map[int]bool, fn func(int), nodes map[int][]int) {
	//putting station to the map so we can understand that we have visited it
	currentStation[node] = true
	fn(node)
	//cheking from all graph elements every edge
	for _, n := range nodes[node] {
		//we have edge to the next station and it is not in map create recursion
		if _, ok := currentStation[n]; !ok {
			Recursive(n, currentStation, fn, nodes)
		}
	}
}

func CheapestOption(route []int, trains Trains) []Train {
	//trains which need to use
	var resultTrains []Train
	var buff Train

	//go through whole route
	for i := 0; i < len(route)-1; i++ {
		//buff price
		price := 1000000.00
		//check all trains
		for j := range trains.Train {
			//if train has this stations and his price is lowest than add it to trains
			if trains.Train[j].DepartureStationId == route[i] && trains.Train[j].ArrivalStationId == route[i+1] && trains.Train[j].Price < price {
				price = trains.Train[j].Price
				buff = trains.Train[j]
			}
		}
		//adding trains
		resultTrains = append(resultTrains, buff)
	}

	return resultTrains
}

func FastestOption(route []int, trains Trains) []Train {
	//trains which need to use
	var resultTrains []Train
	var buff Train

	//go through whole routes
	for i := 0; i < len(route)-1; i++ {
		//buff duration
		duration := 10000000000000.00
		//check all trains
		for j := range trains.Train {
			departureTime, _ := time.Parse("02-01-2006 15:04:05", "22-07-2018 "+trains.Train[j].DepartureTimeString)
			arrivalTime, _ := time.Parse("02-01-2006 15:04:05", "22-07-2018 "+trains.Train[j].ArrivalTimeString)
			difference := arrivalTime.Sub(departureTime).Seconds()
			result := math.Abs(difference)
			//if train has this stations and his duration the lowest add it to trains
			if trains.Train[j].DepartureStationId == route[i] && trains.Train[j].ArrivalStationId == route[i+1] && result < duration {
				duration = result
				buff = trains.Train[j]
			}
		}
		//adding to trains
		resultTrains = append(resultTrains, buff)
	}

	return resultTrains
}
