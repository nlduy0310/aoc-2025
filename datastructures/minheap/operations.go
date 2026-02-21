package minheap

import "github.com/nlduy0310/aoc-2025/sliceutils"

func siftDown[T any](idx int, elements []T, lessFunc func(T, T) bool) {
	minIdx := idx
	leftIdx, rightIdx := leftChildIndex(idx), rightChildIndex(idx)
	if leftIdx >= 0 && leftIdx < len(elements) &&
		lessFunc(elements[leftIdx], elements[minIdx]) {
		minIdx = leftIdx
	}
	if rightIdx >= 0 && rightIdx < len(elements) &&
		lessFunc(elements[rightIdx], elements[minIdx]) {
		minIdx = rightIdx
	}

	if minIdx != idx {
		sliceutils.Swap(elements, idx, minIdx)
		siftDown(minIdx, elements, lessFunc)
	}
}

func siftUp[T any](idx int, elements []T, lessFunc func(T, T) bool) {
	parentIdx := parentIndex(idx)
	if parentIdx >= 0 && parentIdx < len(elements) &&
		lessFunc(elements[idx], elements[parentIdx]) {
		sliceutils.Swap(elements, idx, parentIdx)
		siftUp(parentIdx, elements, lessFunc)
	}
}

func heapify[T any](elements []T, lessFunc func(T, T) bool) {
	for idx := (len(elements) - 2) / 2; idx >= 0; idx-- {
		siftDown(idx, elements, lessFunc)
	}
}

func (mh *MinHeap[T]) Push(val T) {
	mh.elements = append(mh.elements, val)
	siftUp(len(mh.elements)-1, mh.elements, mh.lessFunc)
}

func (mh *MinHeap[T]) Pop() (*T, error) {
	n := mh.Size()
	if n == 0 {
		return nil, emptyHeapError
	}

	ret := mh.elements[0]
	sliceutils.Swap(mh.elements, 0, n-1)
	mh.elements = mh.elements[:n-1]
	if len(mh.elements) > 0 {
		siftDown(0, mh.elements, mh.lessFunc)
	}
	return &ret, nil
}

func (mh MinHeap[T]) Peek() (*T, error) {
	if mh.Size() == 0 {
		return nil, emptyHeapError
	}

	ret := mh.elements[0]
	return &ret, nil
}
