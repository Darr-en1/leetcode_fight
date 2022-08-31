package leet

import "testing"

func TestMyQueue_Push(t *testing.T) {
	t.Run("用栈实现队列", func(t *testing.T) {
		queue := &MyQueue{}
		queue.Push(1)
		queue.Push(2)
		if got := queue.Peek(); got != 1 {
			t.Errorf("Peek() = %v, want %v", got, 1)
		}
		queue.Push(3)
		if got := queue.Pop(); got != 1 {
			t.Errorf("Pop() = %v, want %v", got, 1)
		}
		if got := queue.Pop(); got != 2 {
			t.Errorf("Pop() = %v, want %v", got, 2)
		}
		if got := queue.Pop(); got != 3 {
			t.Errorf("Pop() = %v, want %v", got, 3)
		}
		if got := queue.Empty(); got != true {
			t.Errorf("Empty() = %v, want %v", got, true)
		}

	})
}
