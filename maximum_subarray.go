package main

import "fmt"

// Searches for maximum subarray that crosses the midpoint of given input
// and boundaries
//
// First iterates from the middle down to the start of the given input and
// continuously calculates the sum of the elements. Sets the maximum left index
// for the highest sum found.
//
// Then iterates from the middle up to the end of the given input and
// continuously calculates the sum of the elements. Sets the maximum right index
// for the highest sum found.
//
// Returns the highest left index and highest right index as well as the sum of
// both sums calculated
func MaximumCrossingSubarray(input []int, low, mid, high int) (int, int, int) {
	leftSum := MinInt
	sum := 0
	maxLeft := mid
	maxRight := mid
	for i := mid; i >= low; i-- {
		sum = sum + input[i]
		if sum > leftSum {
			leftSum = sum
			maxLeft = i
		}
	}
	rightSum := MinInt
	sum = 0
	for j := mid + 1; j <= high; j++ {
		sum = sum + input[j]
		if sum > rightSum {
			rightSum = sum
			maxRight = j
		}
	}

	return maxLeft, maxRight, leftSum + rightSum
}

// Searches for the maximum subarray for given input and boundaries
//
// If boundary indexes have the same value it returns the indexes and
// the value in given input at the index
//
// Otherwise recursively searches for maximum subarrays in left half and
// right half of input and for a maximum subarray that crosses the mid point
// of given input
//
// Returns the boundary indexes and sum of the subarray with the largest sum
func MaximumSubarrayInner(input []int, low, high int) (int, int, int) {
	if high == low {
		return low, high, input[low]
	}

	mid := (low + high) / 2
	leftLow, leftHigh, leftSum := MaximumSubarrayInner(input, low, mid)
	rightLow, rightHigh, rightSum := MaximumSubarrayInner(input, mid+1, high)
	crossLow, crossHigh, crossSum := MaximumCrossingSubarray(input, low, mid, high)

	if leftSum >= rightSum && leftSum >= crossSum {
		return leftLow, leftHigh, leftSum
	} else if rightSum >= leftSum && rightSum >= crossSum {
		return rightLow, rightHigh, rightSum
	} else {
		return crossLow, crossHigh, crossSum
	}
}

// Calculates the boundaries of input and kicks off the maximum subarray search
func MaximumSubarray(input []int) (int, int, int) {
	return MaximumSubarrayInner(input, 0, len(input)-1)
}

func main() {
	list := []int{111, 4, 6, -3, 8, -2, 2, -11}
	low, high, sum := MaximumSubarray(list)
	fmt.Printf("The maximum subarray %v with a sum of %v starts at index %v and ends at index %v \n", list[low:high+1], sum, low, high)
}
