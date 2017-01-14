package main

import "fmt"

// Searches for a value in given input and start and end boundaries
// Assumes input is sorted in ascending order
//
// Calculates the middle index using the start and end boundaries
// If value at middle index is the value searched for it returns the middle index
//
// If value at middle index is lower than value searched for it recursively searches
// input in [start, middle)
//
// If value at middle index is greater than value searched for it recursively searches
// input in (middle, end]
func BinarySearchInner(input []int, value, start, end int) int {
	if start == end {
		return start
	} else {
		middle := ((end - start) / 2) + start
		if input[middle] == value {
			return middle
		} else if input[middle] > value {
			return BinarySearchInner(input, value, start, middle-1)
		} else {
			return BinarySearchInner(input, value, middle+1, end)
		}
	}
}

// Calculates the boundaries of input and kicks off the binary search
func BinarySearch(input []int, value int) int {
	return BinarySearchInner(input, value, 0, len(input)-1)
}

func main() {
	list := []int{3, 26, 41, 57, 61, 63, 80, 93}

	for _, value := range list {
		fmt.Println(BinarySearch(list, value))
	}
}
