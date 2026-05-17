package linked_list

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
