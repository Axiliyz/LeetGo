package main

import "fmt"

func maxProfit(prices []int) int {
	res := 0
	min := prices[0]
	for _, price := range prices {
		if price > min {
			res = max(res, price-min)
		} else {
			min = price
		}
	}
	return res
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{4, 3, 2, 1}))
}