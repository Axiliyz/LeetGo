package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(3))
}

func generateParenthesis(n int) []string {
	res := make([]string, 0, 2*n)
	var backtrack func(cur string, open, close int)

	backtrack = func(cur string, open, close int) {
		if len(cur) == 2*n {
			res = append(res, cur)
			return
		}
		if open < n {
			backtrack(cur+"(", open+1, close)
		}
		if close < open {
			backtrack(cur+")", open, close+1)
		}
	}

	backtrack("", 0, 0)
	return res
}
