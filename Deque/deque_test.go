package Deque

import (
	"math/rand/v2"
	"testing"
)

func TestDeque(t *testing.T) {
	t.Run("adding items to front of Deque", func(t *testing.T) {
		d := Deque[any]{}
		initializeDeque(d.AddFront)
		want := d.Items()[0]
		got, _ := d.PeekFront()
		assertEqual(t, got, want)
	})
	t.Run("peeking front from an empty deque", func(t *testing.T) {
		d := Deque[any]{}

		_, isEmpty := d.PeekFront()

		if !isEmpty {
			t.Fatal("did not return zero value")
		}
	})
	t.Run("adding items to back of Deque", func(t *testing.T) {
		d := Deque[any]{}
		initializeDeque(d.AddRear)
		want := d.Items()[d.Size()-1]
		got, _ := d.PeekRear()
		assertEqual(t, got, want)
	})
	t.Run("peeking rear from an empty deque", func(t *testing.T) {
		d := Deque[any]{}

		_, isEmpty := d.PeekRear()

		if !isEmpty {
			t.Fatal("did not return zero value")
		}
	})
	t.Run("removing from the front of deque", func(t *testing.T) {
		d := Deque[any]{}
		initializeDeque(d.AddFront)
		original := append(make([]any, 0, len(d.Items())), d.Items()...)

		want := d.Items()[0]
		initialSize := d.Size()
		got, _ := d.RemoveFront()

		if d.Size() != (initialSize - 1) {
			t.Fatal("RemoveFront operation failed")
		}
		assertEqual(t, got, want)

		if d.Size() > 1 {
			for i := range d.Size() {
				if original[i+1] != d.Items()[i] {
					t.Errorf("RemoveFront operation failed! original is %v, and current is %v", original, d.Items())
					break
				}
			}
		}
	})
	t.Run("removing from the rear of deque", func(t *testing.T) {
		d := Deque[any]{}
		initializeDeque(d.AddRear)
		original := append(make([]any, 0, len(d.Items())), d.Items()...)

		want := d.Items()[d.Size()-1]
		initialSize := d.Size()
		got, _ := d.RemoveRear()

		if d.Size() != (initialSize - 1) {
			t.Fatal("RemoveRear operation failed")
		}
		assertEqual(t, got, want)
		if d.Size() > 1 {
			for i := range d.Size() {
				if original[i] != d.Items()[i] {
					t.Errorf("RemoveRear operation failed! original is %v, and current is %v", original, d.Items())
					break
				}
			}
		}
	})
}

func initializeDeque(dFunc func(any)) {
	loopLen := rand.IntN(10) + 1
	for range loopLen {
		curr := rand.IntN(100)
		dFunc(curr)
	}
}

func assertEqual(t testing.TB, got, want any) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
