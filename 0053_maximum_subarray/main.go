package main

import "fmt"

func maxSubArray(nums []int) int {
	currentSum := nums[0]
	maxSum := nums[0]
	for _, v := range nums[1:] {
		currentSum = max(v, currentSum+v)
		maxSum = max(currentSum, maxSum)
	}
	return maxSum
}

func main() {
	fmt.Println(maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4}))
	fmt.Println(maxSubArray([]int{1}))
	fmt.Println(maxSubArray([]int{5,4,-1,7,8}))
}
