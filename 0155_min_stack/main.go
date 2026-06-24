package main

import (
	"cmp"
	"errors"
	"fmt"
)

var ErrEmptyStack = errors.New("Stack is empty")

type MinStack[T cmp.Ordered] struct {
	data []T
	mins []T
}

func NewMinStack[T cmp.Ordered]() *MinStack[T] {
	return &MinStack[T]{}
}

func (s *MinStack[T]) Push(val T) {
	s.data = append(s.data, val)

	if len(s.mins) == 0 || val <= s.mins[len(s.mins)-1] {
		s.mins = append(s.mins, val)
	}
}

func (s *MinStack[T]) Pop() error {
	if s.IsEmpty() {
		return ErrEmptyStack
	}

	topId := len(s.data) - 1
	val := s.data[topId]

	if val == s.mins[len(s.mins)-1] {
		s.mins = s.mins[:len(s.mins)-1]
	}

	var zero T
	s.data[topId] = zero
	s.data = s.data[:topId]
	return nil
}

func (s *MinStack[T]) Top() (T, error) {
	var zero T
	if s.IsEmpty() {
		return zero, ErrEmptyStack
	}
	return s.data[len(s.data)-1], nil
}

func (s *MinStack[T]) GetMin() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, ErrEmptyStack
	}
	return s.mins[len(s.mins)-1], nil
}

func (s *MinStack[T]) IsEmpty() bool {
	return len(s.data) < 1
}

func main() {
	s := NewMinStack[int]()
	fmt.Println(s.GetMin())
	fmt.Println(s.Pop())
	fmt.Println(s.IsEmpty())
	fmt.Println(s.Top())
	s.Push(8)
	s.Push(1)
	s.Push(2)
	fmt.Println(s.GetMin())
	fmt.Println(s.Pop())
	fmt.Println(s.IsEmpty())
	fmt.Println(s.Top())
}
