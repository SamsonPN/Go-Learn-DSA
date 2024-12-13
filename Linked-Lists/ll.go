package LinkedLists

import "fmt"

const RemoveErr = "could not find item in Linked List"

func NewRemoveErr(msg string, data any) error {
	return fmt.Errorf("%v: %v", msg, data)
}

type Node[T comparable] struct {
	Data T
	Next *Node[T]
	Prev *Node[T]
}

type LinkedList[T comparable] struct {
	head *Node[T]
}

func (ll *LinkedList[T]) add(node *Node[T]) {
	if ll.head == nil {
		ll.head = node
		return
	}

	temp := ll.head
	ll.head = node
	node.Next = temp
}

/*
https://runestone.academy/ns/books/published/pythonds3/BasicDS/TheUnorderedListAbstractDataType.html
*/
func (ll *LinkedList[T]) remove(data T) error {
	var prev *Node[T]
	current := ll.head

	for current != nil && current.Data != data {
		prev = current
		current = current.Next
	}

	if current == nil {
		return NewRemoveErr(string(RemoveErr), data)
	}

	if prev == nil {
		ll.head = ll.head.Next
	} else {
		prev.Next = current.Next
	}
	return nil
}
