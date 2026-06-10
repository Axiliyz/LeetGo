package main

import (
	"fmt"
	"unicode"
)

// func prepareString(s string) string {
// 	return strings.Map(func(r rune) rune {
// 		if unicode.IsLetter(r) || unicode.IsDigit(r) {
// 			return unicode.ToLower(r)
// 		}
// 		return -1
// 	}, s)
// }

func isPalindrome(str string) bool {
	// str := prepareString(s)
	left := 0
	right := len(str) - 1
	for left < right {
		for !unicode.IsLetter(rune(str[left])) && !unicode.IsDigit(rune(str[left])) && left < right {
			left++
		}
		for !unicode.IsLetter(rune(str[right])) && !unicode.IsDigit(rune(str[right])) && left < right {
			right--
		}
		if unicode.ToLower(rune(str[left])) != unicode.ToLower(rune(str[right])) {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	str1 := "race a car"
	str2 := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(str1))
	fmt.Println(isPalindrome(str2))
}
