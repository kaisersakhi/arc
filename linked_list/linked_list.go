package linked_list

import (
	"cmp"
)

type Node[T any] struct {
	Value    T
	Previous *Node[T]
	Next     *Node[T]
}

type LinkedList[T any] struct {
	head    *Node[T]
	tail    *Node[T]
	size    int
	Compare func(a T, b T) int
}

func NewComparableLinkedList[T cmp.Ordered]() *LinkedList[T] {
	return &LinkedList[T]{
		head: nil,
		tail: nil,
		size: 0,
		Compare: func(a, b T) int {
			if a < b {
				return -1
			} else if a > b {
				return 1
			}

			return 0
		},
	}
}

func NewLinkedList[T any](cmp func(a, b T) int) *LinkedList[T] {
	return &LinkedList[T]{
		head:    nil,
		tail:    nil,
		size:    0,
		Compare: cmp,
	}
}

func newNode[T any](value T) *Node[T] {
	return &Node[T]{
		Value:    value,
		Previous: nil,
		Next:     nil,
	}
}

func (list *LinkedList[T]) GetAt(index int) (T, bool) {
	if list.size == 0 || index < 0 || index > list.size {
		var zero T
		return zero, false
	}

	// this is inefficient best and worst case O(N)
	//list.ForEachWithIndex(func(val T, idx int) {
	//	if idx == index {
	//		found = true
	//		value = val
	//		return
	//	}
	//})

	current := list.head

	for i := 0; i < index; i++ {
		current = current.Next
	}

	return current.Value, true
}

func (list *LinkedList[T]) Add(item T) bool {
	node := newNode(item)

	if list.head == nil {
		list.head = node
		list.tail = node
		list.size++
		return true
	}

	tail := list.tail
	list.tail.Next = node
	node.Previous = tail
	list.tail = node
	list.size++

	return true
}
