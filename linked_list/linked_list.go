package linked_list

import (
	"cmp"
	"fmt"
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

func (list *LinkedList[T]) DumpToIO() {
	fmt.Printf("[")
	list.ForEach(func(value T) {
		fmt.Printf("%v ", value)
	})
	fmt.Printf("]")
}

//func (list *LinkedList[T]) IsLast(item T) bool {
//	if list.size > 0 && list.Equal(list.tail.Value, item) {
//		return true
//	}
//
//	return false
//}

func (list *LinkedList[T]) ForEach(fn func(item T)) {
	prt := list.head

	for prt != nil {
		fn(prt.Value)
		prt = prt.Next
	}
}

func (list *LinkedList[T]) ForEachWithIndex(fn func(item T, index int)) {
	var counter = 0

	list.ForEach(func(item T) {
		fn(item, counter)
		counter++
	})
}

func (list *LinkedList[T]) First() (T, bool) {
	if list.head == nil {
		var zero T
		return zero, false
	}

	return list.head.Value, true
}

// Last returns the tail item directly in O(1)
func (list *LinkedList[T]) Last() (T, bool) {
	if list.tail == nil {
		var zero T
		return zero, false
	}

	return list.tail.Value, true
}

// Map Applies fn to every element in the list and returns a new list
func (list *LinkedList[T]) Map(fn func(item T) T) *LinkedList[T] {
	newList := NewLinkedList[T](list.Compare)

	list.ForEach(func(item T) {
		newItem := fn(item)
		newList.Add(newItem)
	})

	return newList
}

func (list *LinkedList[T]) Filter(predicateFn func(item T) bool) *LinkedList[T] {
	newList := NewLinkedList[T](list.Compare)

	list.ForEach(func(item T) {
		if predicateFn(item) {
			newList.Add(item)
		}
	})

	return newList
}

func (list *LinkedList[T]) Reduce(acc T, fn func(acc T, item T) T) T {
	newValue := acc

	list.ForEach(func(item T) {
		newValue = fn(newValue, item)
	})

	return newValue
}

func (list *LinkedList[T]) IsPresent(item T) bool {
	isPresent := false

	list.iterator(func(itemHere T) bool {
		if list.Compare(item, itemHere) == 0 {
			isPresent = true
			return true // Break outer loop and return from iterator
		}

		return false
	})

	return isPresent
}

// iterator Trying to make an efficient iterator, which is breakable
func (list *LinkedList[T]) iterator(fn func(item T) bool) {
	if list.head == nil {
		return
	}

	current := list.head

	for current != nil {
		shouldBreak := fn(current.Value)

		if shouldBreak {
			return
		}

		current = current.Next
	}
}
