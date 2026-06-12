package main

import "fmt"

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))
}

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	maxArea := 0
	for left < right {
		area := (right - left) * min(height[left], height[right])
		maxArea = max(area, maxArea)
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return maxArea
}
