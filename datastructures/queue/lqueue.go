package queue

import "fmt"

// List-based FIFO queue
type LQueue[T any] struct {
	elements []T
}

func EmptyLQueue[T any]() *LQueue[T] {
	ret := LQueue[T]{elements: make([]T, 0)}
	return &ret
}

func (q LQueue[T]) Size() int {
	return len(q.elements)
}

func (q *LQueue[T]) Enqueue(element T) {
	q.elements = append(q.elements, element)
}

func (q *LQueue[T]) Dequeue() (*T, error) {
	if q.Size() == 0 {
		return nil, fmt.Errorf("queue is empty")
	}

	ret := q.elements[0]
	q.elements = q.elements[1:]
	return &ret, nil
}

func (q LQueue[T]) Peek() (*T, error) {
	if q.Size() == 0 {
		return nil, fmt.Errorf("queue is empty")
	}
	ret := q.elements[0]
	return &ret, nil
}
