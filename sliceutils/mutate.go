package sliceutils

func Swap[T any](slice []T, idx1, idx2 int) {
	slice[idx1], slice[idx2] = slice[idx2], slice[idx1]
}
