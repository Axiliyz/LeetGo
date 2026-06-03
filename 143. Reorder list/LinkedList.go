package main

import (
	"errors"
	"fmt"
)

type ListNode[T any] struct {
	Val  T
	Next *ListNode[T]
	Prev *ListNode[T]
}

type LinkedList[T any] struct {
	Head   *ListNode[T]
	Tail   *ListNode[T]
	Length int
}

func NewNode[T any](val T) *ListNode[T] {
	return &ListNode[T]{Val: val}
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (l *LinkedList[T]) Append(val T) {
	defer func() { l.Length++ }()
	node := NewNode(val)
	if l.Head == nil {
		l.Head = node
		l.Tail = node
	} else {
		node.Prev = l.Tail
		l.Tail.Next = node
		l.Tail = node
	}
}

func (l *LinkedList[T]) Remove() error {
	if l.Length < 1 {
		return errors.New("List is empty")
	} else if l.Length == 1 {
		l.Tail = nil
		l.Head = nil
	} else {
		cur := l.Head
		for cur.Next != l.Tail {
			cur = cur.Next
		}
		cur.Next = nil
		l.Tail = cur
	}
	l.Length--
	return nil
}

func (l *LinkedList[T]) Print() {
	if l.Length > 0 {
		cur := l.Head
		fmt.Print(cur.Val)
		cur = cur.Next
		for cur != nil {
			fmt.Print("->", cur.Val)
			cur = cur.Next
		}
		fmt.Println()
	}
}

func (l *LinkedList[T]) Reorder() error {
	if l.Length < 1 {
		return errors.New("List is empty")
	} else if l.Length < 3 {
		return nil
	}

	// current := 0
	left, right := l.Head, l.Tail
	for left != right && left.Next != right {
		leftNext := left.Next
		rightPrev := right.Prev

		left.Next = right
		right.Next = leftNext
		left = leftNext
		right = rightPrev
	}
	left.Next = nil

	var prev *ListNode[T]
	cur := l.Head
	for cur != nil {
		cur.Prev = prev
		prev = cur
		cur = cur.Next
	}
	l.Tail = prev

	return nil
}
