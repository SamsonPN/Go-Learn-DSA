package Stack

import (
	"math/rand/v2"
	"testing"
)

func TestStack(t *testing.T) {
	t.Run("push one item", func(t *testing.T) {
		s := &Stack[int]{}
		want := rand.IntN(100)
		s.Push(want)
		got, isEmpty := s.Peek()

		if isEmpty != false {
			t.Fatal("push operation failed")
		}
		if got != want {
			t.Fatalf("got %v, want %v", got, want)
		}
	})
	t.Run("peek on an empty stack", func(t *testing.T) {
		s := Stack[any]{}
		_, isEmpty := s.Peek()

		if isEmpty != true {
			t.Fatal("expected an error")
		}
	})
	t.Run("push multiple items", func(t *testing.T) {
		s := Stack[any]{}
		want := rand.IntN(100)
		s.Push(1)
		s.Push("one")
		s.Push(want)

		got, _ := s.Peek()

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("pop one item", func(t *testing.T) {
		s := Stack[any]{}
		want := rand.IntN(100)
		s.Push(want)
		got, _ := s.Pop()
		size := s.Size()

		if size != 0 {
			t.Fatalf("pop operation failed")
		}
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("pop from an empty stack", func(t *testing.T) {
		s := Stack[any]{}
		_, isEmpty := s.Pop()

		if !isEmpty {
			t.Error("expecting queue to be empty")
		}
	})
}
