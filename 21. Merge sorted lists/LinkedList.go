package main

import (
	"cmp"
	"errors"
	"fmt"
)

type LinkedList[T cmp.Ordered] struct {
	Head   *ListNode[T]
	Tail   *ListNode[T]
	Length int
}

type ListNode[T cmp.Ordered] struct {
	Val  T
	Next *ListNode[T]
}

func NewListNode[T cmp.Ordered](val T) *ListNode[T] {
	return &ListNode[T]{
		val,
		nil,
	}
}

func NewLinkedList[T cmp.Ordered]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (l *LinkedList[T]) Append(val T) {
	defer func() {l.Length++}()
	newNode := NewListNode(val)
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		l.Tail.Next = newNode
		l.Tail = newNode
	}
}

func (l *LinkedList[T]) Remove() error {
	defer func() {l.Length--}()
	if l.Length < 1 {
		return errors.New("List is empty")
	} else if l.Length == 1 {
		l.Head = nil
		l.Tail = nil
	} else {
		cur := l.Head
		for cur.Next != l.Tail {
			cur = cur.Next
		}
		cur.Next = nil
		l.Tail = cur
	}
	return nil
}

func (l *LinkedList[T]) Reverse() error{
	if l.Length < 2 {
		return nil
	}

	var prev *ListNode[T]
	cur := l.Head
	l.Tail = l.Head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	l.Head = prev
	return nil
}

func (l *LinkedList[T]) String() error {
	if l.Length < 1 {
		return fmt.Errorf("List is empty")
	}
	item := l.Head
	for item.Next != nil {
		fmt.Print(item.Val, "->")
		item = item.Next
	}
	fmt.Print(item.Val)
	return nil
}