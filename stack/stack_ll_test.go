package stack

import "testing"

func TestNewForOrderType(t *testing.T) {
	t.Run("it returns new stack", func(t *testing.T) {
		stack := NewForOrderType[int]()

		stack.Push(1)
		stack.Push(2)
		stack.Push(3)

		el1, _ := stack.GetAt(0) // Top
		el2, _ := stack.GetAt(1)
		el3, _ := stack.GetAt(2)

		if el1 != 3 && el2 != 2 && el3 != 1 {
			t.Errorf("elements should be in reverse order")
		}
	})
}
