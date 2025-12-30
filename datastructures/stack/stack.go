package stack

import "fmt"

type singlyLLNode[T any] struct {
	data T
	next *singlyLLNode[T]
}

func newSinglyLLNode[T any](data T) *singlyLLNode[T] {
	ret := singlyLLNode[T]{data: data, next: nil}
	return &ret
}

type LinkedListStack[T any] struct {
	bottom *singlyLLNode[T]
	top *singlyLLNode[T]
}
// top -> node 1 -> node 2 -> bottom -> nil

func NewEmptyLinkedListStack[T any]() LinkedListStack[T] {
	return LinkedListStack[T]{bottom: nil, top: nil}
}

func (s LinkedListStack[T]) assertTopBottomInvariant() {
	if !((s.top == nil && s.bottom == nil) || (s.top != nil && s.bottom != nil)) {
		panic("top and bottom pointers consistency violated")
	}
}

func (s LinkedListStack[T]) Size() int {
	s.assertTopBottomInvariant()
	if s.top == nil {
		return 0
	}

	size := 1
	curNode := s.top
	for {
		if curNode.next != nil {
			curNode = curNode.next
			size++
		} else {
			break
		}
	}

	return size
}

func (s *LinkedListStack[T]) Push(element T) {
	s.assertTopBottomInvariant()
	if s.top == nil {
		s.top = newSinglyLLNode(element)
		s.bottom = s.top
	} else {
		tmp := s.top
		s.top = newSinglyLLNode(element)
		s.top.next = tmp
	}
}

func (s *LinkedListStack[T]) Pop() (T, error) {
	s.assertTopBottomInvariant()
	if s.top == nil {
		var zero T
		return zero, fmt.Errorf("no elements left in stack")
	}

	popRef := s.top
	if popRef.next != nil {
		s.top = popRef.next
	} else {
		s.top = nil
		s.bottom = nil
	}

	return popRef.data, nil
}

func (s *LinkedListStack[T]) Peek() (T, error) {
	s.assertTopBottomInvariant()
	if s.top == nil {
		var zero T
		return zero, fmt.Errorf("no elements left in stack")
	}

	return s.top.data, nil
}

// GetAt get the element at the provided index, starting from top of the stack
func (s LinkedListStack[T]) GetAt(index int) (T, error) {
	s.assertTopBottomInvariant()

	size := s.Size()
	var zero T
	if size == 0 {
		return zero, fmt.Errorf("stack is empty")
	} else if index < 0 || index >= size {
		return zero, fmt.Errorf("index out of bound")
	}

	curIdx := 0
	curNode := s.top
	for {
		if curIdx == index {
			return curNode.data, nil
		}
		curIdx++
		curNode = curNode.next
	}
}
