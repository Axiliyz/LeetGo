package main

import "fmt"

func main() {
	arr := []int{1, 3, 5, 7, 9}
	l := NewLinkedList[int]()
	for _, v := range arr {
		l.Append(v)
	}
	fmt.Println(l.Reorder())
	l.Print()
}
