package main

import "fmt"

// Adds the bits in a and b
// Assumes a and b are int lists of same length which contain 1 and 0 only
//
// Iterates over both lists from the back and calculates the resulting value
// for each place. If both operands are 1 it sets the carry.
func AddBits(a, b []int) []int {
	c := make([]int, len(a)+1)
	carry := 0
	for j := 1; j <= len(a); j++ {
		c[len(c)-j] = (a[len(a)-j] ^ b[len(b)-j]) ^ carry
		carry = a[len(a)-j] & b[len(b)-j]
	}
	c[0] = carry

	return c
}

func main() {
	a := []int{0, 0, 0, 1, 1}
	b := []int{0, 0, 0, 1, 1}
	fmt.Println(AddBits(a, b))
}
