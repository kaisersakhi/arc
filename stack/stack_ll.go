package stack

import (
	"cmp"

	"github.com/kaisersakhi/arc/linked_list"
)

type StackLL[T any] struct {
	list *linked_list.LinkedList[T]
}

func NewForOrderType[T cmp.Ordered]() *StackLL[T] {
	list := linked_list.NewComparableLinkedList[T]()

	return &StackLL[T]{
		list: list,
	}
}

func (stack *StackLL[T]) Push(item T) bool {
	return stack.list.Append(item)
}

func (stack *StackLL[T]) Pop() (T, error) {
	return stack.list.PopEnd()
}

func (stack *StackLL[T]) GetAt(idx int) (T, bool) {
	return stack.list.GetAt(idx)
}

//func (stack *StackLL[T]) Peek() (T, error) {
//	return
//}
