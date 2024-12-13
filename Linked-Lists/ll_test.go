package LinkedLists

import (
	"testing"
)

func TestNode(t *testing.T) {

	t.Run("initializing node with only Data", func(t *testing.T) {
		node := Node[int]{Data: 1}

		if node.Data != 1 || node.Next != nil {
			t.Error("node does not have correct initialized Data")
		}
	})

	t.Run("initializing node's Next with another node", func(t *testing.T) {
		NextNode := Node[int]{Data: 2}
		node := Node[int]{Next: &NextNode}

		if node.Next != &NextNode && node.Next.Data != 2 {
			t.Error("Unable to access the Next node or its Data")
		}
	})
}

func TestLinkedList(t *testing.T) {
	t.Run("add one node to linked list", func(t *testing.T) {
		linkedList := LinkedList[int]{head: nil}
		node := Node[int]{Data: 31}
		linkedList.add(&node)

		if linkedList.head == nil {
			t.Fatal("the linked list did not initialize properly")
		}
		if linkedList.head.Data != 31 {
			t.Error("the linked list did not add a new node properly")
		}
	})
	t.Run("add multiple nodes to linked list", func(t *testing.T) {
		linkedList := LinkedList[int]{head: nil}
		Data := []int{31, 77, 17, 93, 26, 54}

		for _, val := range Data {
			node := Node[int]{Data: val}
			linkedList.add(&node)
		}

		current := linkedList.head
		count := len(Data) - 1
		for current != nil {
			if current.Data != Data[count] {
				t.Errorf("linked list nodes are not linked correctly. we want %d, but got %d", Data[count], current.Data)
			}
			current = current.Next
			count--
		}
	})

	t.Run("remove node from linkedlist", func(t *testing.T) {
		linkedList := LinkedList[int]{head: nil}
		Data := []int{31, 77, 17}

		for _, val := range Data {
			node := Node[int]{Data: val}
			linkedList.add(&node)
		}

		table := []struct {
			test   string
			Data   int
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
				err := linkedList.remove(testcase.Data)

				if err != nil && err.Error() != testcase.err.Error() {
					t.Errorf("Expected error: %v, but got: %v", testcase.err, err)
				}

				if testcase.err == nil && err != nil {
					t.Errorf("Could not find item %d in Linked List to remove", testcase.Data)
				}
			})
		}
	})
}
