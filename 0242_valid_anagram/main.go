package main

import (
	"fmt"
	"maps"
)

func main() {
	fmt.Println(isAnagram("anagram", "anagram"))
	fmt.Println(isAnagram("rat", "car"))
	fmt.Println(isAnagram("апельсин", "спаниель"))
}

func buildMap(s string) map[rune]int {
	res := make(map[rune]int, len(s))
	for _, r := range s {
		res[r]++
	}
	return res
}

func isAnagram(s string, t string) bool {
	return maps.Equal(buildMap(s), buildMap(t))
}
