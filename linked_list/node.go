package linkedlist

type LinkedListNode[T any] struct {
	value    T
	previous *LinkedListNode[T]
	next     *LinkedListNode[T]
}

func NewLinkedListNode[T any](value T) *LinkedListNode[T] {
	return &LinkedListNode[T]{value: value, previous: nil, next: nil}
}