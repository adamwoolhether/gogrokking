package main

import "fmt"

func main() {
	list := []int{1, 3, 5, 7, 9}

	fmt.Println(binarySearch(list, 3))
	fmt.Println(binarySearch(list, 7))
	fmt.Println(binarySearch(list, 2))

}

// binarySearch takes a sorted list of elements, and a target element.
// It returns the index position of the target element within the list.
// -1 is returned if the number is not found.
func binarySearch(list []int, item int) int {
	low := 0
	high := len(list)

	for low <= high {
		mid := (low + high) / 2
		guess := list[mid]

		if guess == item {
			return mid
		}
		if guess > item {
			high = mid -1
		} else {
			low = mid + 1
		}
	}

	return -1
}
