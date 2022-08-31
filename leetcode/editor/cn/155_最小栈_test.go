package leet

import "testing"

//MinStack minStack = new MinStack();
//minStack.push(-2);
//minStack.push(0);
//minStack.push(-3);
//minStack.getMin();   --> 返回 -3.
//minStack.pop();
//minStack.top();      --> 返回 0.
//minStack.getMin();   --> 返回 -2.

func TestMinStack_Top(t *testing.T) {

	t.Run("最小栈", func(t *testing.T) {
		minStack := Constructor()
		minStack.Push(-2)
		minStack.Push(0)
		minStack.Push(-3)

		if got := minStack.GetMin(); got != -3 {
			t.Errorf("GetMin() = %v, want %v", got, -3)
		}

		minStack.Pop()
		if got := minStack.Top(); got != 0 {
			t.Errorf("Top() = %v, want %v", got, -3)
		}

		if got := minStack.GetMin(); got != -2 {
			t.Errorf("GetMin() = %v, want %v", got, -3)
		}
	})

	t.Run("最小栈", func(t *testing.T) {
		minStack := Constructor()

		if got := minStack.GetMin(); got != 0 {
			t.Errorf("GetMin() = %v, want %v", got, 0)
		}
		minStack.Push(0)
		minStack.Push(1)
		minStack.Push(0)

		if got := minStack.GetMin(); got != 0 {
			t.Errorf("GetMin() = %v, want %v", got, 0)
		}

		minStack.Pop()

		if got := minStack.GetMin(); got != 0 {
			t.Errorf("GetMin() = %v, want %v", got, 0)
		}
	})

	//["MinStack","push","push","push","push","getMin","pop","getMin","pop","getMin","pop","getMin"]
	//	[[],[2],[0],[3],[0],[],[],[],[],[],[],[]]
	t.Run("最小栈", func(t *testing.T) {
		minStack := Constructor()

		if got := minStack.GetMin(); got != 0 {
			t.Errorf("GetMin() = %v, want %v", got, 0)
		}
		minStack.Push(2)
		minStack.Push(0)
		minStack.Push(3)
		minStack.Push(0)

		if got := minStack.GetMin(); got != 0 {
			t.Errorf("GetMin() = %v, want %v", got, 0)
		}

		minStack.Pop()

		if got := minStack.GetMin(); got != 0 {
			t.Errorf("GetMin() = %v, want %v", got, 0)
		}
		minStack.Pop()

		if got := minStack.GetMin(); got != 0 {
			t.Errorf("GetMin() = %v, want %v", got, 0)
		}

		minStack.Pop()

		if got := minStack.GetMin(); got != 2 {
			t.Errorf("GetMin() = %v, want %v", got, 2)
		}

		minStack.Pop()

		if got := minStack.GetMin(); got != 0 {
			t.Errorf("GetMin() = %v, want %v", got, 0)
		}
	})

}
