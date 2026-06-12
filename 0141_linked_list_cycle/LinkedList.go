package main

import "errors"

type LinkedList[T any] struct {
	Head   *ListNode[T]
	Tail   *ListNode[T]
	Length int
}

type ListNode[T any] struct {
	Val  T
	Next *ListNode[T]
}

func NewListNode[T any](value T) *ListNode[T] {
	return &ListNode[T]{value, nil}
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (l *LinkedList[T]) Append(val T) {
	newNode := NewListNode(val)
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		l.Tail.Next = newNode
		l.Tail = newNode
	}
	l.Length++
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

func (l *LinkedList[T]) hasCycle() bool {
	if l.Length < 2 {
		return false
	}
	slow := l.Head
	fast := l.Head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

/* Через мапу - неопт
func hasCycle(head *ListNode) bool {
    seen := map[*ListNode]bool{}
    for cur := head; cur != nil; cur = cur.Next {
        if seen[cur] {
            return true
        }
        seen[cur] = true
    }
    return false
}
*/
