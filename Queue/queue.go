package Queue

import (
	LinkedList "github.com/SamsonPN/Go-Learn-DSA/Linked-Lists"
)

type Queue struct {
	head *LinkedList.Node
	tail *LinkedList.Node
	size int
}

func (q *Queue) Add(val any) {
	newNode := &LinkedList.Node{Data: val}
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

func (q *Queue) Peek() (any, bool) {
	if q.size == 0 {
		return 0, true
	}

	return q.head.Data, false
}

func (q *Queue) Remove() (any, bool) {
	if q.size == 0 {
		return 0, true
	}

	result := q.head.Data
	q.head = q.head.Next
	if q.head != nil {
		q.head.Prev = nil
	}
	q.size--
	return result, false
}
