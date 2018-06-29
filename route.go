package main

type Route struct {
	DepartureStationId int
	ArrivalStationId   int
}

func unicStations(trains Trains) []Route {
	encountered := map[Route]bool{}
	var allStations []Route
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

func createRoute(stations []Route) []Route {
	return stations
}
