package main

import "fmt"

func RecursiveInsertionSort(input []int, n int) []int {
	if (n > 1){
		return RecursiveInsertionSort(input, n - 1)
	} else {
		key := input[n]
		i := n - 1
		for i >= 0 && input[i] > key {
			input[i + 1] = input[i]
			i = i - 1
		}
		input[i + 1] = key
	}
}

func main() {
	list := []int{4, 7, 3, 1, 5, 8, 2, 6}
	fmt.Println(RecursiveInsertionSort(list, len(list)-1))
}
