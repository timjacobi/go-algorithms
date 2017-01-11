package main

import "fmt"

func Left(i int) int {
	return i*2 + 1
}

func Right(i int) int {
	return i*2 + 2
}

func MaxHeapify(input []int, i int) {
	var largest int
	left := Left(i)
	right := Right(i)

	if left < len(input) && input[left] > input[i] {
		largest = left
	} else {
		largest = i
	}
	if right < len(input) && input[right] > input[largest] {
		largest = right
	}
	if i != largest {
		input[i], input[largest] = input[largest], input[i]
		MaxHeapify(input, largest)
	}
}

func BuildMaxHeap(input []int) {
	for i := len(input) / 2; i >= 0; i-- {
		MaxHeapify(input, i)
	}
}

func HeapSort(input []int) {
	BuildMaxHeap(input)
	for i := len(input) - 1; i > 0; i-- {
		input[0], input[i] = input[i], input[0]
		MaxHeapify(input[0:i], 0)
	}
}

func main() {
	list := []int{5, 3, 4, 1, 6, 2}
	HeapSort(list)
	fmt.Println(list)
}
