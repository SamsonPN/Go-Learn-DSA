package Queue

import "testing"

func TestQueue(t *testing.T) {
	t.Run("add one value", func(t *testing.T) {
		queue := Queue{}
		queue.Add(1)

		val, isEmpty := queue.Peek()

		if isEmpty {
			t.Fatal("queue is empty even though we added")
		}
		if val != 1 {
			t.Errorf("got %d, want %d", val, 1)
		}
	})
	t.Run("add multiple values", func(t *testing.T) {
		queue := Queue{}
		queue.Add(1)
		queue.Add(2)
		queue.Add(3)

		val, isEmpty := queue.Peek()

		if isEmpty {
			t.Fatal("queue is empty even though we added")
		}
		if val != 1 {
			t.Errorf("got %d, want %d", val, 1)
		}
	})
	t.Run("delete one value", func(t *testing.T) {
		queue := Queue{}
		queue.Add(1)
		queue.Add(2)
		queue.Add(3)

		val, isEmpty := queue.Remove()

		if isEmpty {
			t.Fatal("no values to remove")
		}
		if val != 1 {
			t.Errorf("got %d, want %d", val, 1)
		}

	})
	t.Run("delete multiple values", func(t *testing.T) {
		queue := Queue{}
		queue.Add(1)
		queue.Add(2)
		queue.Add(3)

		for i := 1; i <= 3; i += 1 {
			val, isEmpty := queue.Remove()

			if isEmpty {
				t.Fatal("no values to remove")
			}
			if val != i {
				t.Errorf("got %d, want %d", val, i)
			}
		}
	})
	t.Run("delete from empty queue", func(t *testing.T) {
		queue := Queue{}
		queue.Add(1)
		queue.Add(2)

		for i := 1; i <= 3; i += 1 {
			_, isEmpty := queue.Remove()

			if i == 3 && !isEmpty {
				t.Error("expecting queue to be empty")
			}
		}
	})
}
