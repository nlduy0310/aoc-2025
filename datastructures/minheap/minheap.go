package minheap

import "github.com/nlduy0310/aoc-2025/sliceutils"

type MinHeap[T any] struct {
	elements []T
	lessFunc func(T, T) bool
}

func New[T any](lessFunc func(T, T) bool) MinHeap[T] {
	return MinHeap[T]{make([]T, 0), lessFunc}
}

func FromSlice[T any](elements []T, lessFunc func(T, T) bool) MinHeap[T] {
	elements_ := sliceutils.ShallowCopy(elements)
	heapify(elements_, lessFunc)
	return MinHeap[T]{elements_, lessFunc}
}

func (mh MinHeap[T]) Size() int {
	return len(mh.elements)
}
