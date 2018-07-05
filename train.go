package main

import "fmt"

func main() {

	trains := readXML()
	unicStations := unicStations(trains)
	//fmt.Print("Unic Stations: ")
	//fmt.Println(unicStations)
	var route []Route

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
}
