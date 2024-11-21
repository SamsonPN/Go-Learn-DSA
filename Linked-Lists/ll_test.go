package LinkedLists

import (
	"testing"
)

func TestNode(t *testing.T) {

	t.Run("initializing node with only data", func(t *testing.T) {
		node := Node{data: 1}

		if node.data != 1 || node.next != nil {
			t.Error("node does not have correct initialized data")
		}
	})

	t.Run("initializing node's next with another node", func(t *testing.T) {
		nextNode := Node{data: 2}
		node := Node{next: &nextNode}

		if node.next != &nextNode && node.next.data != 2 {
			t.Error("Unable to access the next node or its data")
		}
	})
}

func TestLinkedList(t *testing.T) {
	t.Run("add one node to linked list", func(t *testing.T) {
		linkedList := LinkedList{head: nil}
		node := Node{data: 31}
		linkedList.add(&node)

		if linkedList.head == nil {
			t.Fatal("the linked list did not initialize properly")
		}
		if linkedList.head.data != 31 {
			t.Error("the linked list did not add a new node properly")
		}
	})
	t.Run("add multiple nodes to linked list", func(t *testing.T) {
		linkedList := LinkedList{head: nil}
		data := []int{31, 77, 17, 93, 26, 54}

		for _, val := range data {
			node := Node{data: val}
			linkedList.add(&node)
		}

		current := linkedList.head
		count := len(data) - 1
		for current != nil {
			if current.data != data[count] {
				t.Errorf("linked list nodes are not linked correctly. we want %d, but got %d", data[count], current.data)
			}
			current = current.next
			count--
		}
	})

	t.Run("remove node from linkedlist", func(t *testing.T) {
		linkedList := LinkedList{head: nil}
		data := []int{31, 77, 17}

		for _, val := range data {
			node := Node{data: val}
			linkedList.add(&node)
		}

		table := []struct {
			test   string
			data   int
			errMsg string
			err    error
		}{
			{"nonexistent item removal ", 11, "Removed an item that was not present in the list", NewRemoveErr(string(RemoveErr), 11)},
			{"middle item removal", 77, "Could not find 77 in the list", nil},
			{"tail item removal", 31, "Could not find 31 in the list", nil},
			{"head item removal", 17, "Could not find 17 in the list", nil},
		}

		for _, testcase := range table {
			t.Run(testcase.test, func(t *testing.T) {
				err := linkedList.remove(testcase.data)

				if err != nil && err.Error() != testcase.err.Error() {
					t.Errorf("Expected error: %v, but got: %v", testcase.err, err)
				}

				if testcase.err == nil && err != nil {
					t.Errorf("Could not find item %d in Linked List to remove", testcase.data)
				}
			})
		}
	})
}
