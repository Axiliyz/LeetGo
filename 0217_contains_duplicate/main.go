package main

import "fmt"

func containsDuplicate(nums []int) bool {
	seen := make(map[int]bool, len(nums))
	for _, val := range nums {
		if _, ok := seen[val]; ok {
			return true
		}
		seen[val] = true
	}

	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1,2,3,1}))
	fmt.Println(containsDuplicate([]int{1,2,3,4}))
	fmt.Println(containsDuplicate([]int{1,1,3,1,3,4,3,2,4,2}))
}