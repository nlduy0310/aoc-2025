package sliceutils

func Map[S any, R any](source []S, mapper func(S) R) []R {
	ret := make([]R, len(source))
	for idx := range source {
		ret[idx] = mapper(source[idx])
	}
	return ret
}
