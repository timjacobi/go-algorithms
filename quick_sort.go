package main

import (
	"fmt"
)

func Partition(input []int, p, r int) int {
	i := p - 1
	for j := p; j < r; j++ {
		if input[j] < input[r] {
			input[i+1], input[j] = input[j], input[i+1]
			i = i + 1
		}
	}
	input[i+1], input[r] = input[r], input[i+1]

	return i + 1
}

func QuickSortInner(input []int, p, r int) {
	if p < r {
		q := Partition(input, p, r)
		QuickSortInner(input, p, q-1)
		QuickSortInner(input, q+1, r)
	}
}

func QuickSort(input []int) {
	QuickSortInner(input, 0, len(input)-1)
}

func main() {
	list := []int{2, 8, 7, 1, 3, 5, 6, 4}

	QuickSort(list)
	fmt.Println(list)
}
