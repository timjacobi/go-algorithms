package main

import "fmt"

func AddBits(a, b []int) []int{
	c := make([]int, len(a)+1)
	carry := 0
	for j:= 1; j <= len(a); j++ {
		c[len(c)-j] = Abs(Abs(a[len(a)-j] - b[len(b)-j])-carry)
		if a[len(a)-j] == 1 && b[len(b)-j] == 1 {
			carry = 1
		} else {
			carry = 0
		}
	}
	c[0] = carry

	return c
}

func main() {
	a := []int{1,0,0,0,0}
	b := []int{1,0,0,0,0}
	fmt.Println(AddBits(a, b))
}
