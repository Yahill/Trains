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

func createRoute(stations []Route, start int) (route []Route, length int) {
	route = append(route, stations[start])
	visited := map[Route]bool{}

	for j := 0; j < 100; j++ {
		for i := range stations {
			if route[len(route)-1].ArrivalStationId == stations[i].DepartureStationId {
				if visited[stations[i]] != true {
					visited[stations[i]] = true
					route = append(route, stations[i])
				} else {
					next := stations[i]
					route, visited = searchVisited(route, stations, next, visited)
				}
			}
		}
	}
	length = len(visited)
	return
}

func searchVisited(route []Route, stations []Route, next Route, visited map[Route]bool) (nextRoute []Route, newVisited map[Route]bool) {
	for i := range stations {
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
