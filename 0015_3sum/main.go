package main

import (
	"cmp"
	"fmt"
)

func main() {
	arr := []int{-1, 0, 1, 2, -1, -4}
	arr2 := []int{0, 1, 1}
	arr3 := []int{1, 2, 0, 1, 0, 0, 0, 0}
	fmt.Println(threeSum(arr))
	fmt.Println(threeSum(arr2))
	fmt.Println(threeSum(arr3))
}

func threeSum(nums []int) [][]int {
	res := make([][]int, 0, len(nums)/3)
	nums = QuickSort(nums)

	for i := range len(nums) - 2 {
		left := i + 1
		right := len(nums) - 1

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for left < right {
			summa := nums[i] + nums[left] + nums[right]
			if summa == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if summa < 0 {
				left++
			} else if summa > 0 {
				right--
			}
		}
	}
	return res
}

func QuickSort[T cmp.Ordered](arr []T) []T {
	if len(arr) < 2 {
		return arr
	}
	left := make([]T, 0, len(arr)/2)
	right := make([]T, 0, len(arr)/2)
	eq := make([]T, 0, len(arr)/2)
	pivot := arr[len(arr)/2]
	for i := 0; i < len(arr); i++ {
		if arr[i] < pivot {
			left = append(left, arr[i])
		} else if arr[i] > pivot {
			right = append(right, arr[i])
		} else {
			eq = append(eq, arr[i])
		}
	}
	left = QuickSort(left)
	right = QuickSort(right)
	res := append(left, eq...)
	res = append(res, right...)
	return res
}
