package linked_list

func (list *LinkedList[T]) Duplicate() *LinkedList[T] {
	newList := NewLinkedList(list.Compare)

	list.ForEach(func(item T) {
		newList.Add(item)
	})

	return newList
}
