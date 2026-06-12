package main

import "fmt"

func main() {
	arr := []int{3, 2, 0, -4}
	l := NewLinkedList[int]()
	for _, v := range arr {
		l.Append(v)
	}
	l.Reverse()
	a := l.Head
	for a != nil {
		fmt.Println(a.Val)
		a = a.Next
	}
}