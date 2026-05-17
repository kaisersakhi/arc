package linked_list

import (
	"slices"
	"testing"
)

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
			if list.Compare == nil {
				t.Error("Equal should not be nil")
			}

			if list.Compare(1, 1) != 0 {
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

func TestNewLinkedList(t *testing.T) {
	type User struct {
		Name string
		ID   int
	}

	t.Run("When type is non-primitive", func(t *testing.T) {
		t.Run("it should take Equals for comparing", func(t *testing.T) {
			var userList = NewLinkedList[User](func(a, b User) int {
				if a.ID < b.ID {
					return -1
				} else if a.ID > b.ID {
					return 1
				}

				return 0
			})

			userList.Add(User{
				Name: "Kaiser",
				ID:   22,
			})

			userList.Add(User{
				Name: "Sakhi",
				ID:   23,
			})

			if !userList.IsPresent(User{"", 23}) {
				t.Error("user should be present in the list")
			}
		})
	})
}

// For testing, I am not going to expose FindMid
//func TestLinkedList_FindMid(t *testing.T) {
//	t.Run("when there are many items in the list", func(t *testing.T) {
//		t.Run("it returns mid", func(t *testing.T) {
//			list := NewComparableLinkedList[int]()
//
//			list.Add(1)
//			list.Add(2)
//			list.Add(3)
//
//			node := list.FindMid()
//
//			if node.Value != 2 {
//				t.Error("mid value should be 2")
//			}
//		})
//	})
//}

func TestLinkedList_Sort(t *testing.T) {
	t.Run("when list is empty", func(t *testing.T) {

	})

	t.Run("when list is not empty", func(t *testing.T) {
		t.Run("it sorts the list in asscending order", func(t *testing.T) {
			list := NewComparableLinkedList[int]()

			list.Add(3)
			list.Add(55)
			list.Add(2)
			list.Add(0)

			sortedList := list.Sort()

			if sortedList.size != 4 {
				t.Error("list size should be same as the orignal")
			}

			val1, _ := sortedList.GetAt(0)
			val2, _ := sortedList.GetAt(1)
			val3, _ := sortedList.GetAt(2)
			val4, _ := sortedList.GetAt(3)

			if val1 != 0 || val2 != 2 || val3 != 3 || val4 != 55 {
				t.Error("list should be sorted")
			}
		})
	})
}

func TestLinkedList_ReverseForEach(t *testing.T) {
	t.Run("when reversed is directly called", func(t *testing.T) {
		list := NewComparableLinkedList[int]()

		list.Add(3)
		list.Add(55)
		list.Add(2)
		list.Add(0)

		collected := make([]int, 0, 4)

		list.ReverseForEach(func(item int) {
			collected = append(collected, item)
		})

		if !slices.Equal(collected, []int{0, 2, 55, 3}) {
			t.Error("the collection should be reversed")
		}
	})

	t.Run("when reverse for each is done after sorting", func(t *testing.T) {
		t.Run("it should return the list in reverse sorted form", func(t *testing.T) {
			list := NewComparableLinkedList[int]()

			list.Add(3)
			list.Add(55)
			list.Add(2)
			list.Add(0)

			sortedList := list.Sort()

			if sortedList.size != 4 {
				t.Error("list size should be same as the orignal")
			}

			val1, _ := sortedList.GetAt(0)
			val2, _ := sortedList.GetAt(1)
			val3, _ := sortedList.GetAt(2)
			val4, _ := sortedList.GetAt(3)

			if val1 != 0 || val2 != 2 || val3 != 3 || val4 != 55 {
				t.Error("list should be sorted")
			}

			collected := make([]int, 0, 4)

			sortedList.ReverseForEach(func(item int) {
				collected = append(collected, item)
			})

			if !slices.Equal(collected, []int{55, 3, 2, 0}) {
				t.Errorf("%v", collected)
				t.Error("the list should be sorted and reversed")
			}
		})
	})
}
