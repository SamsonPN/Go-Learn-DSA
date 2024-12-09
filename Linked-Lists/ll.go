package LinkedLists

import "fmt"

const RemoveErr = "could not find item in Linked List"

func NewRemoveErr(msg string, data any) error {
	return fmt.Errorf("%v: %v", msg, data)
}

type Node struct {
	Data any
	Next *Node
	Prev *Node
}

type LinkedList struct {
	head *Node
}

func (ll *LinkedList) add(node *Node) {
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
func (ll *LinkedList) remove(data any) error {
	var prev *Node
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
