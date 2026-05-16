package linked_list

import "testing"

func TestNewComparableLinkedList(t *testing.T) {
	t.Run("when newly created", func(t *testing.T) {
		t.Run("head should be nil", func(t *testing.T) {
			var list *LinkedList[int] = NewComparableLinkedList[int]()
			if list.head != nil {
				t.Error("head should be nil")
			}
		})

		t.Run("tail should be nil", func(t *testing.T) {
			var list *LinkedList[int] = NewComparableLinkedList[int]()
			if list.tail != nil {
				t.Error("tail should be nil")
			}
		})

		t.Run("size should be 0", func(t *testing.T) {
			var list *LinkedList[int] = NewComparableLinkedList[int]()
			if list.size != 0 {
				t.Error("size should be 0")
			}
		})

		t.Run("Equal should be set to default", func(t *testing.T) {
			var list *LinkedList[int] = NewComparableLinkedList[int]()
			if list.Equal == nil {
				t.Error("Equal should not be nil")
			}

			if list.Equal(1, 1) == false {
				t.Error("Equal should equal true")
			}
		})
	})

	t.Run("when list has one element", func(t *testing.T) {
		t.Run("head should have point to it", func(t *testing.T) {
			var list *LinkedList[int] = NewComparableLinkedList[int]()

			list.Add(1)

			if list.head == nil {
				t.Error("head should not be nil")
			}

			if list.head.Value != 1 {
				t.Error("head value should be 1")
			}
		})

		t.Run("tail should have point to it", func(t *testing.T) {
			var list *LinkedList[int] = NewComparableLinkedList[int]()
			list.Add(1)
			if list.tail == nil {
				t.Error("tail should not be nil")
			}
		})
	})
}

func TestLinkedList_Add(t *testing.T) {
	var list *LinkedList[int] = NewComparableLinkedList[int]()

	list.Add(1)
	list.Add(2)

	if list.head == nil {
		t.Error("head should not be nil")
	}

	if list.head.Value != 1 {
		t.Error("head value should be 1")
	}

	if list.tail.Value != 2 {
		t.Error("tail value should be 2")
	}

	if list.head.Next.Value != 2 {
		t.Error("head next should be 2")
	}
}

func TestLinkedList_GetAt(t *testing.T) {
	t.Run("when list is empty", func(t *testing.T) {
		list := NewComparableLinkedList[int]()

		var zero int

		item, found := list.GetAt(0)

		if found == true {
			t.Error("found should be false")
		}

		if zero != item {
			t.Error("value should be zero")
		}
	})

	t.Run("when there are items in the list", func(t *testing.T) {
		list := NewComparableLinkedList[int]()

		list.Add(3)
		list.Add(33)
		list.Add(55)
		list.Add(378)

		if val, found := list.GetAt(0); val != 3 || found == false {
			t.Error("value should be 3 and found should be true")
		}

		if val, found := list.GetAt(3); val != 378 || found == false {
			t.Errorf("Value = %v & found = %v", val, found)
			t.Error("value should be 378 and found should be true")
		}
	})
}
