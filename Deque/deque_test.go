package Deque

import (
	"math/rand/v2"
	"testing"
)

func TestDeque(t *testing.T) {
	t.Run("adding items to front of Deque", func(t *testing.T) {
		d := Deque[any]{}
		initializeDeque(d.AddFront)
		want := d.items[0]
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
		want := d.items[d.Size()-1]
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
		want := d.items[0]
		initialSize := d.Size()
		got, _ := d.RemoveFront()

		if d.Size() != (initialSize - 1) {
			t.Fatal("RemoveFront operation failed")
		}
		assertEqual(t, got, want)

		// issue: how do we know that our deletion actually worked?
		// how can we solve this issue if we have duplicates?
	})
	t.Run("removing from the rear of deque", func(t *testing.T) {
		d := Deque[any]{}
		initializeDeque(d.AddRear)
		want := d.items[d.Size()-1]
		initialSize := d.Size()
		got, _ := d.RemoveRear()

		if d.Size() != (initialSize - 1) {
			t.Fatal("RemoveRear operation failed")
		}
		assertEqual(t, got, want)
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
