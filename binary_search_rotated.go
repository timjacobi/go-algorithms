package main

import "fmt"

// Searches for given value on a rotated, sorted array, returns its index
func BinarySearchRotatedInner(list []int, value, start, end int) int {
	mid := ((end - start) / 2) + start

	// If the value is at the middle index return the index
	if list[mid] == value {
		return mid
	} else if list[mid] < list[start] && list[mid] < list[end] {
		// if the value of the middle point is smaller than start and end e.g. [8,1,2,3,4]
		// the pivot point must be left of the of the middle point
		if value >= list[mid] && value <= list[end] {
			// if the value is between the value of the middle point and end i.e. in [2,3,4]
			// continue search right of the middle point
			return BinarySearchRotatedInner(list, value, mid+1, end)
		} else {
			// continue search left of the middle point
			return BinarySearchRotatedInner(list, value, start, mid-1)
		}
	} else {
		// the pivot point must be right of the middle point
		if value >= list[start] && value <= list[mid] {
			// if the value is between the value of start and the middle point i.e. in [8,1,2]
			// continue search left of the middle point
			return BinarySearchRotatedInner(list, value, start, mid-1)
		} else {
			// continue search right of the middle point
			return BinarySearchRotatedInner(list, value, mid+1, end)
		}
	}

}

func BinarySearchRotated(list []int, value int) int {
	return BinarySearchRotatedInner(list, value, 0, len(list)-1)
}

func main() {
	list := []int{10, 20, 1, 3, 6, 7, 8}
	for _, value := range list {
		fmt.Println(BinarySearchRotated(list, value))
	}
}
