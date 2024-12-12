package BinaryHeap

import "testing"

func TestBinaryHeap(t *testing.T) {
	t.Run("insert one value", func(t *testing.T) {
		b := BinaryHeap[int]{}

		b.Insert(1)
		assertCorrectRoot(t, &b, 1)
	})
	t.Run("insert multiple values", func(t *testing.T) {
		b := BinaryHeap[int]{}

		b.Insert(3)
		b.Insert(2)
		b.Insert(1)
		b.Insert(1)
		b.Insert(-1)

		assertCorrectRoot(t, &b, -1)
	})
	t.Run("delete one value", func(t *testing.T) {
		b := BinaryHeap[int]{}

		b.Insert(3)
		b.Insert(2)
		b.Insert(1)
		b.Insert(1)
		b.Insert(-1)

		val, _ := b.Delete()

		if val != -1 {
			t.Errorf("got %d as deleted value, not %d", val, -1)
		}
		assertCorrectRoot(t, &b, 1)
	})
	t.Run("delete multiple values", func(t *testing.T) {
		b := BinaryHeap[int]{}

		for i := 3; i >= -1; i-- {
			b.Insert(i)
		}
		for i := -1; i < 2; i++ {
			val, _ := b.Delete()

			if val != i {
				t.Errorf("got %d from delete, want %d", val, i)
			}
		}
		b.Delete()
		assertCorrectRoot(t, &b, 3)
	})
	t.Run("delete from empty heap", func(t *testing.T) {
		b := BinaryHeap[int]{}

		_, err := b.Delete()
		if err == nil {
			t.Fatal("expected an error but didn't get one")
		}
	})

	/**
	TODO:
	1. heapify method to turn []T into a heap
	2. add a comparator field in the struct to let people choose min or max heap
	*/
}

func assertCorrectRoot(t testing.TB, b *BinaryHeap[int], expected int) {
	t.Helper()
	val, exists := b.Peek()
	if !exists {
		t.Fatal("could not insert values into binary heap!")
	}
	if val != expected {
		t.Errorf("got %d, want %d", val, expected)
	}
}
