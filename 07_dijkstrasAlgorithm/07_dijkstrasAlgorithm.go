package main

import (
	"fmt"
	"math"
)

var (
	graph = make(map[string]map[string]float64)
	costs = make(map[string]float64)
	inf   = math.Inf(1)
	parents = make(map[string]string)
	processed []string
)

func main() {
	// define the graph
	graph["start"] = make(map[string]float64)
	graph["start"]["a"] = 6
	graph["start"]["b"] = 2

	graph["a"] = make(map[string]float64)
	graph["a"]["fin"] = 1

	graph["b"] = make(map[string]float64)
	graph["b"]["a"] = 3
	graph["b"]["fin"] = 5

	graph["fin"] = make(map[string]float64)
	graph["fin"] = nil

	// print only the map of map's graphKeys
	graphKeys := make([]string, len(graph["start"]))
	i := 0
	for k := range graph["start"] {
		graphKeys[i] = k
		i++
	}
	fmt.Println(graphKeys)

	fmt.Println(graph["start"]["a"])
	fmt.Println(graph["start"]["b"])

	// define the costs hash table
	costs["a"] = 6
	costs["b"] = 2
	costs["fin"] = inf

	// define the parents hash table
	parents["a"] = "start"
	parents["b"] = "start"
	parents["fin"] = ""

	// find the lowest cost node and run the algo
	node := findLowestCostNode(costs)
	for node != "" {
		cost := costs[node]
		neighbors := graph[node]

		for n := range neighbors {
			newCost := cost + neighbors[n]

			if costs[n] > newCost {
				costs[n] = newCost
				parents[n] = node
			}
		}
		processed = append(processed, node)
		node = findLowestCostNode(costs)
	}

	fmt.Println(costs)
}

func findLowestCostNode(costs map[string]float64) string {
	lowestCost := inf
	lowestCostNode := ""

	for node := range costs {
		cost := costs[node]

		if cost < lowestCost && !contains(processed, node) {
			lowestCost = cost
			lowestCostNode = node
		}
	}

	return lowestCostNode
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {

			return true
		}
	}

	return false
}