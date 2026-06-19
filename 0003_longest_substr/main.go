package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	seen := make(map[rune]int)
	res := 0
	cur := 0
	for i, v := range s {
		if _, ok := seen[v]; !ok {
			seen[v] = i
			cur++
		} else {
			if seen[v] < i-cur {
				cur++
				seen[v] = i
			} else {
				cur = i - seen[v]
				seen[v] = i
			}
		}
		res = max(res, cur)
	}
	return res
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring(""))
	fmt.Println(lengthOfLongestSubstring(" "))
	fmt.Println(lengthOfLongestSubstring("au"))
	fmt.Println(lengthOfLongestSubstring("dvdf"))
}
