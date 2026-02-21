package sliceutils

func Reduce[T, R any](source []T, init R, fn func(R, T) R) R {
	ret := init
	for _, element := range source {
		ret = fn(ret, element)
	}
	return ret
}
