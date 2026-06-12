package main

import "fmt"

func main() {
	arr := []int{3, 2, 0, -4}
	l := NewLinkedList[int]()
	for _, v := range arr {
		l.Append(v)
	}
	fmt.Println(l.hasCycle())
	l.Tail.Next = l.Head
	fmt.Println(l.hasCycle())
}
