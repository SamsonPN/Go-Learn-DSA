package BinaryHeap

import (
	"testing"
)

func TestMinBinaryHeap(t *testing.T) {
	t.Run("insert one value", func(t *testing.T) {
		b := BinaryHeap[int]{}

		b.Insert(1)
		assertCorrectRoot(t, &b, 1)
	})
	t.Run("insert multiple values", func(t *testing.T) {
		b := BinaryHeap[int]{
			comparator: func(a, b int) bool {
				return a <= b
			},
		}

		b.Insert(3)
		b.Insert(2)
		b.Insert(1)
		b.Insert(1)
		b.Insert(-1)

		assertCorrectRoot(t, &b, -1)
	})
	t.Run("delete one value", func(t *testing.T) {
		b := BinaryHeap[int]{
			comparator: func(a, b int) bool {
				return a <= b
			},
		}

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
		b := BinaryHeap[int]{
			comparator: func(a, b int) bool {
				return a <= b
			},
		}

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
		b := BinaryHeap[int]{
			comparator: func(a, b int) bool {
				return a <= b
			},
		}

		_, err := b.Delete()
		if err == nil {
			t.Fatal("expected an error but didn't get one")
		}
	})
	t.Run("heapify from a slice", func(t *testing.T) {
		b := BinaryHeap[int]{
			comparator: func(a, b int) bool {
				return a <= b
			},
		}
		input := []int{3, 6, 9, -1, 10, 2, 4, 5, 8, 7, 11, 0, 1}

		b.Heapify(input)

		for i := -1; i <= 11; i++ {
			val, _ := b.Delete()
			if val != i {
				t.Errorf("heapify failed: got %d, want %d", val, i)
			}
		}
	})
}

func TestMaxBinaryHeap(t *testing.T) {
	t.Run("inserting one value", func(t *testing.T) {
		b := BinaryHeap[int]{
			comparator: func(a, b int) bool {
				return a >= b
			},
		}

		b.Insert(1)
		assertCorrectRoot(t, &b, 1)
	})
	t.Run("inserting multiple values", func(t *testing.T) {
		b := BinaryHeap[int]{
			comparator: func(a, b int) bool {
				return a >= b
			},
		}

		b.Insert(1)
		b.Insert(3)
		b.Insert(2)
		assertCorrectRoot(t, &b, 3)
	})
	t.Run("delete one value", func(t *testing.T) {
		b := BinaryHeap[int]{
			comparator: func(a, b int) bool {
				return a >= b
			},
		}

		b.Insert(3)
		b.Insert(2)
		b.Insert(1)
		b.Insert(1)
		b.Insert(-1)

		val, _ := b.Delete()

		if val != 3 {
			t.Errorf("got %d as deleted value, not %d", val, -1)
		}
		assertCorrectRoot(t, &b, 2)
	})
	t.Run("delete multiple values", func(t *testing.T) {
		b := BinaryHeap[int]{
			comparator: func(a, b int) bool {
				return a >= b
			},
		}

		for i := 0; i <= 5; i++ {
			b.Insert(i)
		}
		for i := 5; i >= 2; i-- {
			val, _ := b.Delete()

			if val != i {
				t.Errorf("got %d from delete, want %d", val, i)
			}
		}
		b.Delete()
		assertCorrectRoot(t, &b, 0)
	})
	t.Run("delete from empty heap", func(t *testing.T) {
		b := BinaryHeap[int]{
			comparator: func(a, b int) bool {
				return a >= b
			},
		}

		_, err := b.Delete()
		if err == nil {
			t.Fatal("expected an error but didn't get one")
		}
	})
	t.Run("heapify from a slice", func(t *testing.T) {
		b := BinaryHeap[int]{
			comparator: func(a, b int) bool {
				return a >= b
			},
		}
		input := []int{3, 6, 9, -1, 10, 2, 4, 5, 8, 7, 11, 0, 1}

		b.Heapify(input)

		for i := 11; i >= -1; i-- {
			val, _ := b.Delete()
			if val != i {
				t.Errorf("heapify failed: got %d, want %d", val, i)
			}
		}
	})
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
