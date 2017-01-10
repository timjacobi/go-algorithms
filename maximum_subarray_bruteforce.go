package main

import "fmt"

func MaximumSubarrayBrutforce(input []int) (int, int, int) {
	var left, right int
	sum := MinInt

	for i := 0; i < len(input); i++ {
		currentSum := 0
		for j := i; j < len(input); j++ {
			currentSum = currentSum + input[j]
			if currentSum > sum {
				sum = currentSum
				left = i
				right = j
			}
		}
	}

	return left, right, sum

}

func main() {
	list := []int{111, 4, 6, -3, 8, -2, 2, -11}
	low, high, sum := MaximumSubarrayBrutforce(list)
	fmt.Printf("The maximum subarray %v with a sum of %v starts at index %v and ends at index %v \n", list[low:high+1], sum, low, high)
}
