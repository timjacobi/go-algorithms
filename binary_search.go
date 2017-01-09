package main

import "fmt"

func BinarySearchInner(input []int, value, start, end int) int {
	if end-start == 0 {
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

func BinarySearch(input []int, value int) int {
	return BinarySearchInner(input, value, 0, len(input)-1)
}

func main() {
	list := []int{3, 26, 41, 57, 61, 63, 80, 93}

	for _, value := range list {
		fmt.Println(BinarySearch(list, value))
	}
}
