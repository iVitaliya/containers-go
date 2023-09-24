package linkedlist

import "github.com/iVitaliya/containers-go/utils"

// ****************************************************************************
//
//	# LINKED LIST
//
// ****************************************************************************
type LinkedList[T any] struct {
	head *LinkedListNode[T]
	tail *LinkedListNode[T]
	size int
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{head: nil, tail: nil, size: 0}
}

//***************************************************************************
//                                  INSPECTION
//***************************************************************************

// Returns the size of the Linked List. - O(1)
func (list *LinkedList[T]) Size() int {
	return list.size
}

// Returns `true` if the Linked List has been found empty, `false` otherwise. - O(1)
func (list *LinkedList[T]) IsEmpty() bool {
	return list.size == 0
}

//***************************************************************************
//                                 INSERTION
//***************************************************************************

// Adds a Node to the head of the LinkedList. - O(1)
//
// value: the value to add to the LinkedList.
func (list *LinkedList[T]) AddFront(value T) bool {
	newNode := NewLinkedListNode[T](value)

	if list.size == 0 {
		list = &LinkedList[T]{
			head: newNode,
			tail: newNode,
			size: 1,
		}
	} else {
		// Link old head backwards.
		list.head.previous = newNode

		// Link new head forwards.
		newNode.next = list.head

		list.head = newNode
		list.size++
	}

	return true
}

// Adds a Node to the tail of the LinkedList. - O(1)
//
// value: the value to add to the LinkedList.
func (list *LinkedList[T]) AddBack(value T) bool {
	newNode := NewLinkedListNode[T](value)

	if list.size == 0 {
		list = &LinkedList[T]{
			head: newNode,
			tail: newNode,
			size: 1,
		}
	} else {
		// Link old tail forwards.
		list.tail.next = newNode

		// Link new tail backwards.
		newNode.previous = list.tail

		list.tail = newNode
		list.size++
	}

	return true
}

// Adds a Node at specified index. - O(1)
//
// index: the index on where to add specified value.
//
// value: the value to add to the LinkedList.
func (list *LinkedList[T]) AddAt(index int, value T) bool {
	if index == 0 {
		return list.AddFront(value)
	}

	if index == list.size {
		list.AddBack(value)
	}

	if index < 0 || index >= list.size {
		return false
	}

	current := list.head

	// Traverse to the index.
	for i := 0; i < index-1; i++ {
		current = current.next
	}

	newNode := NewLinkedListNode[T](value)

	// Link the next Node.
	current.next.previous = newNode
	newNode.next = current.next

	// Link the previous Node.
	newNode.previous = current
	current.next = newNode

	list.size++

	return true
}

//***************************************************************************
//                                 ACCESSING
//***************************************************************************

// Gets the value of the head of the current Node. - O(1)
func (list *LinkedList[T]) PeekFront() T {
	if list == nil || list.size == 0 {
		return *new(T)
	}

	return list.head.value
}

// Gets the value of the tail of the current Node. - O(1)
func (list *LinkedList[T]) PeekBack() T {
	if list == nil || list.size == 0 {
		return *new(T)
	}

	return list.tail.value
}

// Gets the element at the specified index. - O(1)
//
// index: the index to use for getting the Node.
func (list *LinkedList[T]) Get(index int) T {
	if index < 0 || index >= list.size || list == nil {
		return *new(T)
	}

	i := 0
	current := list.head

	for i < index {
		current = current.next

		i++
	}

	return current.value
}

//***************************************************************************
//                                 SEARCHING
//***************************************************************************

// Removes the first occurrence of the specified item in the LinkedList.
//
// value: the value to search for.
//
// Returns the index of the first occurrence of the element, and -1 if the element doesn't seem to exist in the List.
func (list *LinkedList[T]) IndexOf(value T) int {
	// List is empty.
	if list.size == 0 {
		return -1
	}

	index := 0
	current := list.head

	for !utils.DefaultEquals[T](current.value, value) {
		// current.value === null means we reached end of list without finding element.
		if current.next == nil {
			return -1
		}

		current = current.next
		index += 1
	}

	return index
}

// Checks if the specified value exists in the LinkedList.
//
// value: the value to search for.
//
// Returns whether or not the specified value exists in the LinkedList.
func (list *LinkedList[T]) Contains(value T) bool {
	index := list.IndexOf(value)

	return index != -1
}

//***************************************************************************
//                                 DELETION
//***************************************************************************

// Removes the head from the LinkedList. - O(1)
//
// Returns the value of removed head.
func (list *LinkedList[T]) RemoveFront() T {
	if list == nil || list.size == 0 {
		return *new(T)
	}

	// Extract the value of the head so we can return it later.
	value := list.head.value

	if list.head.next != nil {
		list.head.next.previous = nil

		// Move the head pointer forwards.
		list.head = list.head.next

		list.size--
	} else {
		// List is size 1, clear the list.
		list = nil
	}

	return value
}

// Removes the tail from the LinkedList. - O(1)
//
// Returns the value of removed tail.
func (list *LinkedList[T]) RemoveBack() T {
	if list == nil || list.size == 0 {
		return *new(T)
	}

	// Extract the value of the tail so we can return it later.
	value := list.tail.value

	if list.tail.next != nil {
		list.tail.previous.next = nil

		// Move the tail pointer backwards.
		list.tail = list.tail.previous

		list.size--
	} else {
		// List is size 1, clear the list.
		list = nil
	}

	return value
}

// Removes first occurence of the Node with specified value. Returns true if
// the removal was successful, and false otherwise. - O(n)
//
// value: the value to remove.
//
// Returns the value of the removed Node.
func (list *LinkedList[T]) Remove(value T) T {
	index := list.IndexOf(value) // O(n)

	if index == -1 {
		return *new(T)
	}

	return list.RemoveAt(index) // O(n)
}

func (list *LinkedList[T]) RemoveAt(index int) T {
	if list == nil || list.size == 0 {
		return *new(T)
	}

	if index == 0 {
		return list.RemoveFront()
	} else if index == list.size-1 {
		return list.RemoveBack()
	}

	if index < 0 || index >= list.size {
		return *new(T)
	}

	i := 0
	current := list.head

	for i < index {
		current = current.next

		i++
	}

	// Delete the Node.
	current.previous.next = current.next
	current.next.previous = current.previous

	list.size--

	return current.value
}

// Deletes all the registered Nodes. - O(1)
func (list *LinkedList[T]) Clear() {
	list = nil
}

//***************************************************************************
//                                 HELPERS
//***************************************************************************

// Appends values from an Array to the List. - O(k)
func (list *LinkedList[T]) FromArray(array []T) *LinkedList[T] {
	for i := 0; i < len(array); i++ {
		list.AddBack(array[i])
	}

	return list
}
