package main

import "fmt"

func PairWithSumSorted(list []int, sum int) bool {
	low := 0
	high := len(list) - 1

	for low < high {
		currentSum := list[low] + list[high]
		if currentSum == sum {
			return true
		} else if currentSum < sum {
			low++
		} else {
			high--
		}
	}

	return false
}

func PairWithSum(list []int, sum int) bool {
	dict := map[int]bool{}

	for i := 0; i < len(list); i++{
		currentElement := list[i]
		_, ok := dict[sum-currentElement]
		if ok {
			return true
		} else {
			dict[currentElement] = true
		}
	}
	return false
}

func main(){
	fmt.Println(PairWithSumSorted([]int{1, 2, 7, 3, 8}, 6))
	fmt.Println(PairWithSumSorted([]int{1, 2, 3, 4, 8}, 6))
	fmt.Println(PairWithSum([]int{4, 4, 4, 1}, 20))
	fmt.Println(PairWithSum([]int{4, 4, 16, 1}, 20))
}
