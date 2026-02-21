package minheap

func leftChildIndex(idx int) int {
	return 2*idx + 1
}

func rightChildIndex(idx int) int {
	return 2*idx + 2
}

func parentIndex(idx int) int {
	return (idx - 1) / 2
}
