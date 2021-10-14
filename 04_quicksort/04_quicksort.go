package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	array := []int{10, 5, 2, 3, 48, 20, 5, 9, 89, 4, 69, 2}
	fmt.Println(quicksort(array))
}

func quicksort(array []int) []int {
	if len(array) < 2 {
		return array
	} else {
		// pivot := array[0] // we should chose a random element, to get the average case.
		rand.Seed(time.Now().UnixNano())
		pivot := array[rand.Intn(len(array))]

		var less []int
		var greater []int

		for _, v := range array[1:] {
			if v <= pivot {
				less = append(less, v)
			} else {
				greater = append(greater, v)
			}
		}

		less = append(quicksort(less), pivot)
		greater = quicksort(greater)
		return append(less, greater...)
	}
}

// Sum done with a for loop
/*
func sum(arr []int) int {
	var total int
	for _, x := range arr {
		total += x
	}
	return total
}
*/

// Sum done with recursion
/*
func sumRecursive(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	return arr[0] + sum(arr[1:])
}
*/

// Recusively count number of items in a list:
/*
func count(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	return 1 + count(arr[1:])
}
*/

// Find the greatest number in a list:
/*
func max(arr []int) int {
	if len(arr) == 2 {
		if arr[0] > arr[1] {
			return arr[0]
		} else {
			return arr[1]
		}
	}

	subMax := max(arr[1:])
	if arr[0] > subMax {
		return arr[0]
	} else {
		return subMax
	}
}
*/