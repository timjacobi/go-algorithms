package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

type CutRodFn func(map[int]int, int) int

func MemoizeCutRod(cutRodFn CutRodFn) func(int) int{
	memo := make(map[int]int)
	memo[1] = 1
	memo[2] = 5
	memo[3] = 8
	memo[4] = 9
	memo[5] = 10
	memo[6] = 17
	memo[7] = 17
	memo[8] = 20
	memo[9] = 24
	memo[10] = 30

	return func (n int) int {
		v, ok := memo[n]
		if ok {
			return v
		} else {
			return cutRodFn(memo, n)
		}
	}
}

func CutRod(p map[int]int, n int) int {
	if n == 0 {
		return 0
	}
	q := -1

	for i := 1; i <= n; i++ {
		q = max(q, p[i] + CutRod(p, n-i))
	}

	return q
}

func PowerSet(set []int) [][]int {
	var result [][]int

	if len(set) == 0 {
		result = append(result, []int{})
	} else {
		end := len(set) -1
		allSubSets := PowerSet(set[:end])

		for i := 0; i < len(allSubSets); i++ {
			subset := allSubSets[i]
			subsetplus := append(subset, set[end])
			result = append(result, subset, subsetplus)

		}
	}

	return result
}

func RecursiveMultiply(x, y int) int {
	if x == 0 {
		return 0
	}

	return y + RecursiveMultiply(x - 1, y)
}

func Permutations(input string) []string {
	var result []string
	if len(input) == 1 {
		result = append(result, input)
	} else {
		// Take the last letter off the word
		letter := string(input[len(input)-1])
		subString := string(input[:len(input)-1])

		// Calculate all the permutations for the remainder
		subPermutations := Permutations(subString)

		// Iterate over them
		for i := 0; i < len(subPermutations); i++ {
			subPermutation := subPermutations[i]

			// Insert letter at each possible place in the permutation
			// and add the result to the list of permutations
			for j := 0; j <= len(subPermutation); j++ {
				result = append(result, string(subPermutation[j:])+letter+string(subPermutation[:j]))
			}
		}
	}

	return result
}

func ValidParens(amount int) []string {
	var result []string
	if amount == 1 {
		result = append(result, "()")
	} else {
		subSolutions := ValidParens(amount - 1)

		for i := 0; i < len(subSolutions); i++ {
			subSolution := subSolutions[i]
			result = append(result, "()"+subSolution)
			result = append(result, "("+subSolution+")")
			if "()"+subSolution != subSolution+"()"{
				result = append(result, subSolution+"()")
			}
		}
	}

	return result
}

func Coins(amount int) int {
	coins := map[int]bool{1:true, 5:true, 10:true, 25:true}

	ways := map[int]int{0:0, 1:1}
}

//func makeChange(amount int, denoms []int, index int) int {
//	if index >= len(denoms) -1 {
//		return 1
//	}
//	denomAmount := denoms[index]
//	ways := 0
//	for i:= 0; i * denomAmount <= amount; i++ {
//		amountRemaining := amount - i*denomAmount
//		ways += makeChange(amountRemaining, denoms, index + 1)
//	}
//	return ways
//}
//func makeChangeW(amount int) int {
//	denoms := []int{25, 10, 5, 1}
//	return makeChange(amount, denoms, 0)
//}


func main(){
	//p := map[int]int{0:0, 1:1, 2:5, 3:8, 4:9, 5:10, 6:17, 7:17, 8:20, 9:24, 10:30}
	//fmt.Println(CutRod(p, 10))
	//
	//q := MemoizeCutRod(CutRod)
	//
	//fmt.Println(q(4))
	//
	//fmt.Println(PowerSet([]int{1, 2, 3}))
	//
	//fmt.Println(RecursiveMultiply(5, 8))
	//
	//fmt.Println(Permutations("aba"))
	//
	//fmt.Println(ValidParens(10))

	for i:=1; i <= 100; i++ {
		fmt.Printf("%v: %v %v\n", i, makeChangeW(i), Coins(i))
	}
}