package main

import (
	"errors"
	"fmt"
)

func isValid(s string) bool {
	stack := &Stack{}
	pairs := map[string]string {
		")": "(",
		"]": "[",
		"}": "{",
	}

	for _, ch := range s {
		char := string(ch)

		if char == "(" || char == "[" || char == "{" {
			stack.add(char)
		} else {
			top, err := stack.remove()
			if err != nil {
				return false
			}
			if pairs[char] != top {
				return false
			}
		}
	}

	return len(stack.data) == 0
}

type Stack struct {
	data []string
}

func (s *Stack) add(item string) {
	s.data = append(s.data, item)
}

func (s *Stack) remove() (string, error) {
	if len(s.data) < 1 {
		return "", errors.New("Wrong operation, size < 1")
	}
	elem := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return elem, nil
}

func (s *Stack) peek() (string, error) {
	if len(s.data) < 1 {
		return "", errors.New("Wrong operation, size < 1")
	}
	return s.data[len(s.data)-1], nil
}

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("([])"))
	fmt.Println(isValid("([)]"))
}
