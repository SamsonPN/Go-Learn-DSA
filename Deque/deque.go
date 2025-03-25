package Deque

import Utility "github.com/SamsonPN/Go-Learn-DSA/utility"

type Deque[T any] struct {
	items []T
}

func (d *Deque[T]) AddFront(val T) {
	newItem := make([]T, 0, len(d.items)+1)
	newItem = append(newItem, val)
	d.items = append(newItem, d.items...)
}

func (d *Deque[T]) PeekFront() (T, bool) {
	if d.IsEmpty() {
		return Utility.GenerateZeroValue[T](), true
	}

	return d.items[0], false
}

func (d *Deque[T]) AddRear(val T) {
	d.items = append(d.items, val)
}

func (d *Deque[T]) PeekRear() (T, bool) {
	if d.IsEmpty() {
		return Utility.GenerateZeroValue[T](), true
	}

	return d.items[d.Size()-1], false
}

func (d *Deque[T]) RemoveFront() (T, bool) {
	if d.IsEmpty() {
		return Utility.GenerateZeroValue[T](), true
	}
	returnVal := d.items[0]
	d.items = d.items[1:]
	return returnVal, false
}

func (d *Deque[T]) RemoveRear() (T, bool) {
	if d.IsEmpty() {
		return Utility.GenerateZeroValue[T](), true
	}
	returnVal := d.items[d.Size()-1]
	d.items = d.items[:d.Size()-1]
	return returnVal, false
}

func (d *Deque[T]) Items() []T {
	return d.items
}

func (d *Deque[T]) Size() int {
	return len(d.items)
}

func (d *Deque[T]) IsEmpty() bool {
	return d.Size() == 0
}
