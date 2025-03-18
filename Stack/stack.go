package Stack

import (
	Utility "github.com/SamsonPN/Go-Learn-DSA/utility"
)

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(val T) {
	s.items = append(s.items, val)
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		return Utility.GenerateZeroValue[T](), true
	}
	val := s.items[s.Size()-1]
	s.items = s.items[:s.Size()-1]
	return val, false
}

func (s *Stack[T]) Peek() (T, bool) {
	if s.IsEmpty() {
		return Utility.GenerateZeroValue[T](), true
	}
	return s.items[s.Size()-1], false
}

func (s *Stack[T]) Size() int {
	return len(s.items)
}

func (s *Stack[T]) IsEmpty() bool {
	return s.Size() == 0
}
