package linked_list

func (list *LinkedList[T]) Sort() *LinkedList[T] {
	newList := list.Duplicate()

	newList.head = mergeSort(list.head, list)

	return newList
}

func mergeSort[T any](node *Node[T], list *LinkedList[T]) *Node[T] {
	if node == nil || node.Next == nil {
		return node
	}

	left, right := split(node)

	left = mergeSort(left, list)
	right = mergeSort(right, list)

	return merge(left, right, list)
}

func merge[T any](left *Node[T], right *Node[T], list *LinkedList[T]) *Node[T] {
	// What do i want to do here?

	//1. Left or Right can be 0, 1 or N

	// 2. Check each list

	var zero T
	var newHead *Node[T] = newNode[T](zero)
	var current *Node[T] = newHead

	leftCurrent := left
	rightCurrent := right

	for leftCurrent != nil && rightCurrent != nil {
		switch list.Compare(leftCurrent.Value, rightCurrent.Value) {
		case -1:
			// Left is lesser
			current.Next = leftCurrent
			leftCurrent.Previous = current
			leftCurrent = leftCurrent.Next
		case 1:
			// Right is lesser
			current.Next = rightCurrent
			rightCurrent.Previous = current
			rightCurrent = rightCurrent.Next
		case 0:
			// Both are equal, append both
			current.Next = leftCurrent
			leftCurrent.Previous = current
			leftCurrent = leftCurrent.Next

			current = current.Next

			current.Next = rightCurrent
			rightCurrent.Previous = current
			rightCurrent = rightCurrent.Next
		}

		current = current.Next
	}

	// check directly append either of the remaining sub-lists
	if leftCurrent != nil {
		current.Next = leftCurrent
		leftCurrent.Previous = current
	}

	if rightCurrent != nil {
		current.Next = rightCurrent
		rightCurrent.Previous = current
	}

	newNodeHead := newHead.Next
	newNodeHead.Previous = nil

	return newNodeHead
}

func split[T any](node *Node[T]) (*Node[T], *Node[T]) {
	if node == nil || node.Next == nil {
		return node, nil
	}

	mid := findMid(node)

	left, right := node, mid

	// delink previous node
	mid.Previous.Next = nil

	return left, right
}

func findMid[T any](node *Node[T]) *Node[T] {
	if node == nil {
		return nil
	}

	slow := node
	fast := node

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}
