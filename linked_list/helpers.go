package linked_list

import "fmt"

func (list *LinkedList[T]) Duplicate() *LinkedList[T] {
	newList := NewLinkedList(list.Compare)

	list.ForEach(func(item T) {
		newList.Add(item)
	})

	return newList
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
