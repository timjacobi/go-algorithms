package main

import "fmt"

// Given an index for a node it returns the index of the node's left child
func Left(i int) int {
	return i*2 + 1
}

// Given an index for a node it returns the index of the node's right child
func Right(i int) int {
	return i*2 + 2
}

// Generates a max heap for root node in input at given index
// Assumes children of node are roots of max heaps
//
// First compares value of node with value of left child and stores
// the index of the larger of the two
//
// Then compares the value of the larger node to the value of the
// right child and store the index of the larger of the two
//
// Finally if any of the children node's values was larger it swaps
// the value of the child node with the value of the root node and
// recurses on the child node to ensure the max heap property is
// kept for the subtree
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

// Builds a max heap from the given input
//
// It iterates from the most left parent of the leaf nodes
// (which start at |input|/2) up to the root and calls BuildMaxHeap
// on each node
func BuildMaxHeap(input []int) {
	for i := len(input) / 2; i >= 0; i-- {
		MaxHeapify(input, i)
	}
}

// Sorts the given input
//
// First builds a max heap from the given input then
// Continuously swaps the first and last element in the
// heap, removes the last element from the heap and
// restores the max heap property for the rest of the
// heap until no more elements are left and the list
// is sorted
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
