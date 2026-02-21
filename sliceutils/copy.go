package sliceutils

func ShallowCopy[T any](source []T) []T {
	ret := make([]T, len(source))
	copy(ret, source)
	return ret
}
