package main

import "fmt"

func main() {
	list := []int{5, 3, 6, 2, 10}
	list2 := []int{6, 46, 21, 6, 78}
	fmt.Println(selectionSort(list))
	fmt.Println(selectionSort(list2))
}

func selectionSort(list []int) []int {
	var sorted []int

	for _ = range list {
		leastIndex := findLeastIndex(list)
		sorted = append(sorted, list[leastIndex])
		list = append(list[:leastIndex], list[leastIndex+1:]...)
	}

	return sorted
}

func findLeastIndex(list []int) int {
	least := list[0]
	leastIndex := 0

	for i, v := range list {
		if v < least {
			least = v
			leastIndex = i
		}
	}

	return leastIndex
}

func findGreatestIndex(list []int) int {
	greatest := list[0]
	greatestIndex := len(list) - 1

	for i, v := range list {
		if v > greatest {
			greatest = v
			greatestIndex = i
		}
	}

	return greatestIndex
}
