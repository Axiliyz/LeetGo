package main

import "fmt"

func main() {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
	fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
}

func longestConsecutive(nums []int) int {
	set := make(map[int]bool, len(nums))
	for _, n := range nums {
		set[n] = true
	}

	res := 0
	for n := range set {
		if set[n-1] {
			continue
		}
		len := 1
		for set[n+len] {
			len++
		}
		res = max(res, len)
	}
	return res
}
