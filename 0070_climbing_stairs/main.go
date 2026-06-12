package main

import "fmt"

func climbStairs(n int) int {
	if n <= 1 {
		return 1
	}
	prev, curr := 1, 1
	for i := 2; i <= n; i++ {
		prev, curr = curr, prev+curr
	}
	return curr
}

func main() {
	fmt.Println(climbStairs(1))
	fmt.Println(climbStairs(2))
	fmt.Println(climbStairs(5))
	fmt.Println(climbStairs(12))
	fmt.Println(climbStairs(55))
	fmt.Println(climbStairs(100))
}

/* Рекурсия не выдержала 44 элемент
func climbStairs(n int) int {
    if n <= 1 {
        return 1
    }
	return climbStairs(n-1) + climbStairs(n-2)
}
*/

/* Глобальный кэш - беда
var cache = map[int]int{}
func climbStairs(n int) int {
    if n <= 1 {
        return 1
    }
    if val, ok := cache[n]; ok {
        return val
    }
    result := climbStairs(n-1) + climbStairs(n-2)
    cache[n] = result
    return result
}
*/