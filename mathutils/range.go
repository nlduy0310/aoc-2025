package mathutils

import "cmp"

func InRange[T cmp.Ordered](leftBound, rightBound T, target T, includeLower, includeUpper bool) bool {
	var lowerCond, upperCond bool
	if includeLower {
		lowerCond = target >= leftBound
	} else {
		lowerCond = target > leftBound
	}
	if includeUpper {
		upperCond = target <= rightBound
	} else {
		upperCond = target < rightBound
	}
	return lowerCond && upperCond
}

func InRangeUnordered[T cmp.Ordered](bound1, bound2 T, target T, includeLower, includeUpper bool) bool {
	minBound, maxBound := MinMax(bound1, bound2)
	return InRange(minBound, maxBound, target, includeLower, includeUpper)
}
