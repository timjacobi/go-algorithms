package main

import "fmt"

// Sorts given input in order descending
//
// Iterates over elements 1 - n and shifts
// all elements left of current element
// that are smaller than the current element one
// index to the right. Inserts current element
// at index the last element was moved from.
func InsertionSortReverse(input []int) {
	for j := 1; j < len(input); j++ {
		key := input[j]
		i := j - 1
		for i >= 0 && input[i] < key {
			input[i+1] = input[i]
			i = i - 1
		}
		input[i+1] = key

	}
}

func main() {
	list := []int{4, 7, 3, 1, 5, 8, 2, 6}
	InsertionSortReverse(list)
	fmt.Println(list)
}
