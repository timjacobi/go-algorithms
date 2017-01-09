package main

import "fmt"

func SelectionSort(input []int) []int {
	for i := 0; i < len(input)-1; i++ {
		minIndex := i
		for j := minIndex; j < len(input); j++ {
			if input[j] < input[minIndex] {
				minIndex = j
			}
		}
		input[i], input[minIndex] = input[minIndex], input[i]
	}
	return input
}

func main() {
	list := []int{4, 7, 3, 1, 5, 8, 2, 6}
	fmt.Println(SelectionSort(list))
}
