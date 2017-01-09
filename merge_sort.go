package main

import "fmt"

func Merge(a, b []int) []int {
	result := make([]int, len(a)+len(b))
	idxa := 0
	idxb := 0
	idxresult := 0
	for idxa < len(a) && idxb < len(b) {
		if a[idxa] < b[idxb] {
			result[idxresult] = a[idxa]
			idxa++

		} else {
			result[idxresult] = b[idxb]
			idxb++
		}
		idxresult++
	}

	for idxa < len(a) {
		result[idxresult] = a[idxa]
		idxa++
		idxresult++
	}

	for idxb < len(b) {
		result[idxresult] = b[idxb]
		idxb++
		idxresult++
	}

	return result
}

func MergeSort(input []int) []int {
	var a, b []int
	middle := len(input) / 2
	if(len(input) > 2) {
		a = MergeSort(input[0:middle])
		b = MergeSort(input[middle:len(input)])
	} else {
		a = input[0:1]
		b = input[1:len(input)]
	}
	return Merge(a, b)
}


func main() {
	list := []int{4, 7, 3, 1, 5, 8, 2, 6}
	fmt.Println(MergeSort(list))
}