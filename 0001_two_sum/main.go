package main

import "fmt"

func twoSum(nums []int, target int) []int {
	mapa := make(map[int]int, len(nums))
	for i, v := range nums {
		if j, ok := mapa[target-v]; ok {
			return []int{j, i}
		}
		mapa[v] = i

	}
	return []int{0, 0}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
}
