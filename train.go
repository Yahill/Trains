package main

import (
	"fmt"
)

func main() {

	trains := readXML()
	unicStations := unicStations(trains)
	route := createRoute(unicStations)
	fmt.Println(route)
}
