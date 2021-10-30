package main

import "fmt"

func main() {
	states := []string{"mt", "wa", "or", "id", "nv", "ut", "ca", "az"}
	statesNeeded := make(map[string]struct{})
	for _, v := range states {
		statesNeeded[v] = struct{}{}
	}
	stations := make(map[string][]string)
	stations["kone"] = []string{"id", "nv", "ut"}
	stations["ktwo"] = []string{"wa", "id", "mt"}
	stations["kthree"] = []string{"or", "nv", "ca"}
	stations["kfour"] = []string{"nv", "ut"}
	stations["kfive"] = []string{"ca", "az"}

	finalStations := make(map[string]struct{})

	for len(statesNeeded) > 0 {
		var bestStation string
		statesCovered := make(map[string]struct{})

		for station, stationStates := range stations {
			covered := intersection(statesNeeded, stationStates)

			if len(covered) > len(statesCovered) {
				bestStation = station
				statesCovered = covered
			}

		}
		for k, _ := range statesCovered {
			delete(statesNeeded, k)
		}
		finalStations[bestStation] = struct{}{}
	}

	fmt.Println(finalStations)
}

func intersection(a map[string]struct{}, b []string) map[string]struct{} {
	intersection := make(map[string]struct{})
	for _, v := range b {

		if _, ok := a[v]; !ok {
			continue
		}
		intersection[v] = struct{}{}
	}

	return intersection
}