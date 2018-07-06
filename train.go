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
	var route []Route

	option := flag.String("option", "nill", "Choose option: cheapest ot fastest.")

	flag.Parse()

	//reading xml file
	trains := readXML()
	//finding unic stations
	unicStations := unicStations(trains)

	//searching for the best route
	for i := range unicStations {
		length := 200000
		buff, buffLength := createRoute(unicStations, i)

		if buffLength < length {
			route = buff
			length = buffLength
		}
	}
	fmt.Print("Route: ")
	fmt.Println(route)

	//check option
	switch *option {

	case "nill":
		fmt.Println("Choose option: -option=cheepest or -option=fastest.")

	case "cheapest":
		cheepestWay := cheapest(route, trains)
		for i := range cheepestWay {
			fmt.Println("TrainID: " + strconv.Itoa(cheepestWay[i].TrainId) + "\n" + "DepartureStationId: " + strconv.Itoa(cheepestWay[i].DepartureStationId) + "\n" +
				"ArrivalStationId: " + strconv.Itoa(cheepestWay[i].ArrivalStationId) + "\n" + "Price: " + strconv.FormatFloat(cheepestWay[i].Price, 'f', 2, 32) + "\n" +
				"ArrivalTimeString: " + cheepestWay[i].ArrivalTimeString + "\n" + "DepartureTimeString: " + cheepestWay[i].DepartureTimeString + "\n")
		}

	case "fastest":
		fastestWay := fastest(route, trains)
		for i := range fastestWay {
			fmt.Println("TrainID: " + strconv.Itoa(fastestWay[i].TrainId) + "\n" + "DepartureStationId: " + strconv.Itoa(fastestWay[i].DepartureStationId) + "\n" +
				"ArrivalStationId: " + strconv.Itoa(fastestWay[i].ArrivalStationId) + "\n" + "Price: " + strconv.FormatFloat(fastestWay[i].Price, 'f', 2, 32) + "\n" +
				"ArrivalTimeString: " + fastestWay[i].ArrivalTimeString + "\n" + "DepartureTimeString: " + fastestWay[i].DepartureTimeString + "\n")

		}
	}
}

//reading xml file
func readXML() (trains Trains) {

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

//searching for unic stations
func unicStations(trains Trains) []Route {
	//map with visited stations
	encountered := map[Route]bool{}
	//structure with arrivalStationId and departureStationId
	var allStations []Route
	//structure with all unic stations
	var result []Route

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
	return result
}

//create the best route to visit all stations
func createRoute(stations []Route, start int) (route []Route, length int) {
	//appending starting station
	route = append(route, stations[start])
	//map with visited stations
	visited := map[Route]bool{}

	for j := 0; j < 100; j++ {
		//going through all the stations
		for i := range stations {
			//cheking if stations connected
			if route[len(route)-1].ArrivalStationId == stations[i].DepartureStationId {
				//if we haven't visited this station yet append it and put into map with visited stations
				if visited[stations[i]] != true {
					//put in a map
					visited[stations[i]] = true
					//appending to the route
					route = append(route, stations[i])
				} else {
					//if we haven't got unvisited stations, taking next station
					next := stations[i]
					//checking if we have station connected between this stations in vivited map
					route, visited = searchVisited(route, stations, next, visited)
				}
			}
		}
	}
	length = len(visited)
	return
}

//checking visited route
func searchVisited(route []Route, stations []Route, next Route, visited map[Route]bool) (nextRoute []Route, newVisited map[Route]bool) {
	for i := range stations {
		//checking if we have station connected with previous and next station
		if route[len(route)-1].ArrivalStationId == stations[i].DepartureStationId && next.DepartureStationId == stations[i].ArrivalStationId {
			if visited[stations[i]] == true {
				visited[stations[i]] = true
				route = append(route, stations[i])
			}
		}
	}
	nextRoute = route
	newVisited = visited
	return
}

//searching for the cheapest trains
func cheapest(route []Route, trains Trains) (trainsFinal []Train) {
	var price float64
	var buff Train

	for i := range route {
		price = 1000000.00
		for j := range trains.Train {
			//checking trains with the same route and looking for cheapest one
			if route[i].ArrivalStationId == trains.Train[j].ArrivalStationId && route[i].DepartureStationId == trains.Train[j].DepartureStationId && trains.Train[j].Price < price {
				buff = trains.Train[j]
				price = trains.Train[j].Price
			}
		}
		trainsFinal = append(trainsFinal, buff)
	}
	return
}

//searching for the fastest trains
func fastest(route []Route, trains Trains) (trainsFinal []Train) {
	var time float64
	var buff Train

	for i := range route {
		time = 1000000.00
		for j := range trains.Train {
			//getting how long train goes between stations
			buffTime := duration(trains.Train[j].ArrivalTimeString, trains.Train[j].DepartureTimeString)
			//checking trains with the same route and looking for the fastest one
			if route[i].ArrivalStationId == trains.Train[j].ArrivalStationId && route[i].DepartureStationId == trains.Train[j].DepartureStationId && buffTime < time {
				buff = trains.Train[j]
				time = buffTime
			}
		}
		trainsFinal = append(trainsFinal, buff)
	}
	return
}

//counting how long train goes between stations
func duration(arrival, departure string) (result float64) {
	var layout string
	var arrivalBuff string
	var departureBuff string

	layout = "2006-01-02T15:04:05.000Z"

	arrivalBuff = "2017-08-31T" + arrival + ".000Z"
	departureBuff = "2017-08-31T" + departure + ".000Z"
	arrivalTime, _ := time.Parse(layout, arrivalBuff)
	departureTime, _ := time.Parse(layout, departureBuff)

	difference := arrivalTime.Sub(departureTime).Seconds()
	result = math.Abs(difference)

	return
}
