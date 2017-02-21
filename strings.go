package main

import (
	"fmt"
	"math"
)

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

func StringIntersection(a, b string) string {
	result := ""
	for _, v := range a {
		for _, v2 := range b {
			if v == v2 {
				result = result + string(v)
			}
		}
	}
	return result
}

func StringIntersectionLinear(a, b string) string {
	dict := make(map[string]interface{})
	result := ""
	for _, v := range a {
		dict[string(v)] = struct{}{}
	}

	for _, v := range b {
		if dict[string(v)] != nil {
			result = result + string(v)
		}
	}

	return result
}

func StringDuplicateCharacters(a string) string {
	dict := make(map[string]interface{})
	result := ""
	for _, v := range a {
		if dict[string(v)] != nil {
			result = result + string(v)
		} else {
			dict[string(v)] = struct{}{}
		}
	}
	return result
}

func StringIsAnagram(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	dict := make(map[string]int)
	for _, v := range a {
		if _, ok := dict[string(v)]; ok {
			dict[string(v)]++
		} else {
			dict[string(v)] = 1
		}
	}

	for _, v := range b {
		if _, ok := dict[string(v)]; !ok {
			return false
		} else if dict[string(v)]--; dict[string(v)] < 0 {
			return false
		}
	}
	return true
}

func IntToEnglishAux(n float64) string {
	words := map[float64]string{
		0:  "zero",
		1:  "one",
		2:  "two",
		3:  "three",
		4:  "four",
		5:  "five",
		6:  "six",
		7:  "seven",
		8:  "eight",
		9:  "nine",
		10: "ten",
		11: "eleven",
		12: "twelve",
		13: "thirteen",
		14: "fourteen",
		15: "fifteen",
		16: "sixteen",
		17: "seventeen",
		18: "eighteen",
		19: "nineteen",
		20: "twenty",
		30: "thirty",
		40: "fourty",
		50: "fifty",
		60: "sixty",
		70: "seventy",
		80: "eighty",
		90: "ninety",
	}

	var result string

	lastTwoDigits := math.Remainder(n, 100)

	if lastTwoDigits < 20 && lastTwoDigits != 0 {
		result = words[lastTwoDigits]
	} else {
		ones := math.Remainder(lastTwoDigits, 10)
		if ones > 0 {
			result = words[ones]
		}
		tens := lastTwoDigits - ones
		if tens > 0 {
			result = words[tens] + result
		}
	}

	hundreds := (n - lastTwoDigits)/100

	if hundreds > 0 {
		result = words[hundreds] + "hundred" + result
	}


	return fmt.Sprint(result)
}

func IntToEnglish(n int) string {
	number := float64(n)

	var result string

	words := []string{"", "thousand", "million", "billion", "trillion"}

	for i:=0; number > 0; i++ {
		current := math.Remainder(number, 1000)
		result = IntToEnglishAux(current) + words[i] + " " + result
		number = (number - current)/1000
	}

	return result
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

	fmt.Println(StringIntersection("ab", "abc"))
	fmt.Println(StringIntersectionLinear("aaaab", "aabc"))
	fmt.Println(StringDuplicateCharacters("abccd"))
	fmt.Println(StringIsAnagram("abaca", "bcaaa"))
}
