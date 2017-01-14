package main

import (
	"fmt"
)

// Partitions given input around a pivot element
// so that all elements left of the pivot element
// are smaller than the pivot element and all
// elements right of the pivot element are larger
// than the pivot element. Returns the index of
// the pivot element.
func Partition(input []int, p, r int) int {
	i := p - 1
	for j := p; j < r; j++ {
		if input[j] <= input[r] {
			input[i+1], input[j] = input[j], input[i+1]
			i = i + 1
		}
	}
	input[i+1], input[r] = input[r], input[i+1]

	return i + 1
}

// Sorts the given input in given boundaries
//
// First calculates a partition point using Partition
// and then recurses on subarrays left and right of
// partition point, i.e. [low boundary, partition point)
// and (partition point, high boundary]
func QuickSortInner(input []int, p, r int) {
	if p < r {
		q := Partition(input, p, r)
		QuickSortInner(input, p, q-1)
		QuickSortInner(input, q+1, r)
	}
}

// Calculates the boundaries of input and kicks off the sorting
func QuickSort(input []int) {
	QuickSortInner(input, 0, len(input)-1)
}

func main() {
	list := []int{2, 8, 7, 1, 3, 5, 6, 4}

	QuickSort(list)
	fmt.Println(list)
}
