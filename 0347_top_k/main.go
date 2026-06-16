package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
	fmt.Println(topKFrequent([]int{1}, 1))
	fmt.Println(topKFrequent([]int{1, 2}, 2))
}

func topKFrequent(nums []int, k int) []int {
	counts := make(map[int]int)
	for _, num := range nums {
		counts[num]++
	}
	pairs := makePairs(counts)
	sorted := sortPairs(pairs)
	return printKPairs(sorted, k)
}

type pair struct {
	k int
	v int
}

func makePairs(counts map[int]int) []pair {
	pairs := make([]pair, 0, len(counts))
	for k, v := range counts {
		pairs = append(pairs, pair{k, v})
	}
	return pairs
}

func sortPairs(pairs []pair) []pair {
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].v > pairs[j].v
	})
	return pairs
}

func printKPairs(pairs []pair, k int) []int {
	if k > len(pairs) {
		fmt.Println("k is greater than the number of unique elements in the input array")
		return nil
	}
	res := make([]int, 0, k)
	for i := 0; i < k; i++ {
		res = append(res, pairs[i].k)
	}
	return res
}
