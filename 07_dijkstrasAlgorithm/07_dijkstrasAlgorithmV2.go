package main

import (
	"fmt"
	"math"
)

// I've left these commented out, more verbose map constructions for reference.
var (
	infV2       = math.Inf(1)
	processedV2 = make(map[string]struct{})
)

func main() {
	// define the graphV2
	graphV2 := map[string]map[string]float64{
		"start": map[string]float64{
			"a": 6,
			"b": 2,
		},
		"a": map[string]float64{
			"fin": 1,
		},
		"b": map[string]float64{
			"a":   3,
			"fin": 5,
		},
		"fin": {},
	}

	// define the costsV2 hash table
	costsV2 := map[string]float64{
		"a":   6,
		"b":   2,
		"fin": infV2,
	}

	// define the parentsV2 hash table
	parentsV2 := map[string]string{
		"a":   "start",
		"b":   "start",
		"fin": "",
	}

	// find the lowest cost nodeV2 and run the algo
	nodeV2 := findLowestCostNodeV2(costsV2)
	for nodeV2 != "" {
		cost := costsV2[nodeV2]
		neighbors := graphV2[nodeV2]

		for n := range neighbors {
			newCost := cost + neighbors[n]

			if costsV2[n] > newCost {
				costsV2[n] = newCost
				parentsV2[n] = nodeV2
			}
		}
		processedV2[nodeV2] = struct{}{}
		nodeV2 = findLowestCostNodeV2(costsV2)
	}

	fmt.Println(costsV2)
}

func findLowestCostNodeV2(costs map[string]float64) string {
	lowestCost := infV2
	lowestCostNode := ""

	for node := range costs {
		cost := costs[node]

		if cost < lowestCost && !containsV2(node) {
			lowestCost = cost
			lowestCostNode = node
		}
	}

	return lowestCostNode
}

func containsV2(node string) bool {
	_, ok := processedV2[node]
	if ok {
		return true
	}

	return false
}
