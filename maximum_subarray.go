package main

import "fmt"

func MaximumCrossingSubarray(input []int, low, mid, high int) (int, int, int) {
	leftSum := MinInt
	sum := 0
	maxLeft := high
	maxRight := low
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

func MaximumSubarrayInner(input []int, low, high int) (int, int, int){
	if high == low {
		return low, high, input[low]
	}

	mid := (low+high)/2
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

func MaximumSubarray(input []int) (int, int, int){
	return MaximumSubarrayInner(input, 0, len(input)-1)
}

func main(){
	list := []int{111, 4, 6, -3, 8, -2, 2, -11}
	low, high, sum := MaximumSubarray(list)
	fmt.Printf("The maximum subarray %v with a sum of %v starts at index %v and ends at index %v", list[low:high+1], sum, low, high)
}