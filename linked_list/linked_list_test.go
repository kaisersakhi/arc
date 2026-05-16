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
}
