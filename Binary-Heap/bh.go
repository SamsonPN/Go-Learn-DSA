package BinaryHeap

import (
	"fmt"

	constraints "golang.org/x/exp/constraints"
)

/*
*
put this into a utility folder later!
*/
func generateZeroValue[T any]() T {
	var zeroValue T
	return zeroValue
}

type BinaryHeap[T constraints.Ordered] struct {
	heap []T
	size int
}

func (b *BinaryHeap[T]) percolateUp(index int) {
	child := index
	parent := (index - 1) / 2

	for child != 0 {
		if b.heap[parent] <= b.heap[child] {
			return
		}
		b.heap[parent], b.heap[child] = b.heap[child], b.heap[parent]
		child = parent
		parent = (child - 1) / 2
	}
}

func (b *BinaryHeap[T]) percolateDown() {
	parent := 0
	for parent != b.size-1 {
		leftChild := (2 * parent) + 1
		rightChild := (2 * parent) + 2

		// if left child is OoB, then right is as well
		if leftChild >= b.size {
			return
		}

		if rightChild >= b.size {
			if b.heap[parent] <= b.heap[leftChild] {
				return
			}
			b.heap[parent], b.heap[leftChild] = b.heap[leftChild], b.heap[parent]
			parent = leftChild
		}
		var minChild int
		if b.heap[leftChild] <= b.heap[rightChild] {
			minChild = leftChild
		} else {
			minChild = rightChild
		}

		if b.heap[parent] <= b.heap[minChild] {
			return
		}

		b.heap[parent], b.heap[minChild] = b.heap[minChild], b.heap[parent]
		parent = minChild
	}
}

func (b *BinaryHeap[T]) Insert(val T) {
	b.heap = append(b.heap, val)
	b.size++

	if b.size != 1 {
		b.percolateUp(b.size - 1)
	}
}

func (b *BinaryHeap[T]) Peek() (T, bool) {
	if b.size == 0 {
		return generateZeroValue[T](), false
	}
	return b.heap[0], true
}

func (b *BinaryHeap[T]) Delete() (T, error) {
	if b.size == 0 {
		return generateZeroValue[T](), fmt.Errorf("trying to delete from an empty heap")
	}

	res, _ := b.Peek()
	b.heap[0] = b.heap[b.size-1]
	b.heap = b.heap[:b.size-1]
	b.size--
	b.percolateDown()
	return res, nil
}
