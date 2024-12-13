package Queue

import (
	LinkedList "github.com/SamsonPN/Go-Learn-DSA/Linked-Lists"
	Utility "github.com/SamsonPN/Go-Learn-DSA/utility"
)

type Queue[T comparable] struct {
	head *LinkedList.Node[T]
	tail *LinkedList.Node[T]
	size int
}

func (q *Queue[T]) Add(val T) {
	newNode := &LinkedList.Node[T]{Data: val}
	if q.size == 0 {
		q.head = newNode
		q.tail = newNode
	} else {
		q.tail.Next = newNode
		newNode.Prev = q.tail
		q.tail = newNode
	}

	q.size++
}

func (q *Queue[T]) Peek() (T, bool) {
	if q.size == 0 {
		return Utility.GenerateZeroValue[T](), true
	}

	return q.head.Data, false
}

func (q *Queue[T]) Remove() (T, bool) {
	if q.size == 0 {
		return Utility.GenerateZeroValue[T](), true
	}

	result := q.head.Data
	q.head = q.head.Next
	if q.head != nil {
		q.head.Prev = nil
	}
	q.size--
	return result, false
}
