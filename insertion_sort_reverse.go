package main

import "fmt"

func InsertionSortReverse(input []int) []int {
	for j := 1; j < len(input); j++ {
		key := input[j]
		i := j - 1
		for i >= 0 && input[i] < key {
			input[i+1] = input[i]
			i = i -1
		}
		input[i+1] = key

	}
	return input
}

func main() {
	list := []int{4,7,3,1,5,8,2,6}
	fmt.Println(InsertionSortReverse(list))
}
