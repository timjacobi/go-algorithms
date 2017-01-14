package main

import "fmt"

// Merges given inputs in ascending order
//
// Iterates over given inputs and adds the smaller value
// at current index to result until one of the inputs
// is exhausted.
// Then iterates of rest of the input that wasn't exhausted
// and adds remaining elements to the result.
func Merge(a, b []int) []int {
	result := make([]int, len(a)+len(b))
	idxa := 0
	idxb := 0
	idxresult := 0
	for idxa < len(a) && idxb < len(b) {
		if a[idxa] < b[idxb] {
			result[idxresult] = a[idxa]
			idxa++

		} else {
			result[idxresult] = b[idxb]
			idxb++
		}
		idxresult++
	}

	for idxa < len(a) {
		result[idxresult] = a[idxa]
		idxa++
		idxresult++
	}

	for idxb < len(b) {
		result[idxresult] = b[idxb]
		idxb++
		idxresult++
	}

	return result
}

// Sorts given input
//
// Splits the input in two halves and then recurses
// on them until their size is 1. Then merges the
// two halves using Merge and returns the result.
func MergeSort(input []int) []int {
	if len(input) == 1 {
		return input
	}

	middle := len(input) / 2
	a := MergeSort(input[0:middle])
	b := MergeSort(input[middle:len(input)])

	return Merge(a, b)
}

func main() {
	list := []int{4, 7, 3, 1, 5, 8, 2, 6}
	fmt.Println(MergeSort(list))
}
