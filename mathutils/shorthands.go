package mathutils

import "cmp"

func MinMax[T cmp.Ordered](a, b T) (T, T) {
	return min(a, b), max(a, b)
}
