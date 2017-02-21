package main

import "fmt"

func IsSubstring(a, b string) bool {
	bPtr := 0

	for i := range a {
		if a[i] == b[bPtr] {
			if bPtr == len(b) - 1 {
				return true
			}
			bPtr++
		} else if bPtr != 0 && a[i] != a[i-1] {
			return false
		}
	}
	return false
}

func IsAnagram(a, b string) bool {
	dict := map[rune]int{}

	for _, s:= range a {
		dict[s]++
	}

	for _, s:= range b {
		if dict[s]--; dict[s] < 0 {
			return false
		}
	}
	return true
}

func IsPalindrome(a string) bool {
	for i := range a {
		if a[i] == a[len(a) - 1 -i] {
			if i == len(a)/2 {
				return true
			}
		}
	}
	return false
}

func main(){
	fmt.Println(IsSubstring("aaaaabb", "abb"))
	fmt.Println(IsSubstring("abcde", "bce"))

	fmt.Println()

	fmt.Println(IsAnagram("bad", "dab"))
	fmt.Println(IsAnagram("bad", "daba"))

	fmt.Println()

	fmt.Println(IsPalindrome("aba"))
	fmt.Println(IsPalindrome("abba"))
	fmt.Println(IsPalindrome("abca"))
}
