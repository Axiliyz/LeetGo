package main

import (
	"cmp"
)

func mergeTwoLists[T cmp.Ordered](list1 *LinkedList[T], list2 *LinkedList[T]) *LinkedList[T] {
	res := NewLinkedList[T]()
	i1 := list1.Head
	i2 := list2.Head
	for i1 != nil && i2 != nil {
		if i1.Val < i2.Val {
			res.Append(i1.Val)
			i1 = i1.Next
		} else {
			res.Append(i2.Val)
			i2 = i2.Next
		}
	}

	if i1 != nil {
		res.Tail.Next = i1
		res.Tail = list1.Tail 
	} else if i2 != nil {
		res.Tail.Next = i2
		res.Tail = list2.Tail 
	}
	res.Length = list1.Length + list2.Length
	return res
}

func main() {
	arr1 := []int{1, 3, 4}
	arr2 := []int{1, 2, 4, 6, 7, 8, 11, 13, 14, 16, 19, 22}
	l1 := NewLinkedList[int]()
	l2 := NewLinkedList[int]()
	for _, v := range arr1 {
		l1.Append(v)
	}

	for _, v := range arr2 {
		l2.Append(v)
	}

	res := mergeTwoLists(l1, l2)
	res.String()
}