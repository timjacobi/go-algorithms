package main

import "fmt"

// Sorts the given input by iterating from the lowest
// element up to second to last input element and comparing
// it with all element that a right of the current element
// stores the index of the smallest. If the smalles element
// is not the current element it swaps the current element
// with the smallest element.
func SelectionSort(input []int) {
	for i := 0; i < len(input)-1; i++ {
		minIndex := i
		for j := minIndex; j < len(input); j++ {
			if input[j] < input[minIndex] {
				minIndex = j
			}
		}
		input[i], input[minIndex] = input[minIndex], input[i]
	}
}

func main() {
	list := []int{4, 7, 3, 1, 5, 8, 2, 6}
	SelectionSort(list)
	fmt.Println(list)
}

