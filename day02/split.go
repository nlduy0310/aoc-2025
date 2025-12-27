package day02

import (
	"math"

	"github.com/nlduy0310/aoc-2025/day02/idrange"
	"github.com/nlduy0310/aoc-2025/mathutils"
)

// maxIntWithDigits returns the maximum int that has the provided number of digits.
//
// This function assumes the number of digits are positive
func maxIntWithDigits(digits uint) int {
	return int(math.Pow10(int(digits)) - 1)
}

// minIntWithDigits returns the minimum int that has the provided number of digits.
//
// This function assumes the number of digits are positive
func minIntWithDigits(digits uint) int {
	if digits == 1 {
		return 0
	}

	return int(math.Pow10(int(digits - 1)))
}

// splitRangeByDigitsCount splits a range into sub-ranges where
// all elements in each resulting sub-range have the same digits count.
//
// This function assumes the range is non-negative.
func splitRangeByDigitsCount(r idrange.IdRange) []idrange.IdRange {
	lowerDigitsCount := mathutils.CountDigits(r.LowerVal)
	upperDigitsCount := mathutils.CountDigits(r.UpperVal)
	if lowerDigitsCount == upperDigitsCount {
		return []idrange.IdRange{r}
	}

	ret := make([]idrange.IdRange, 0, upperDigitsCount-lowerDigitsCount+1)
	for count := lowerDigitsCount; count <= upperDigitsCount; count++ {
		var subrange idrange.IdRange
		var err error
		switch count {
		case lowerDigitsCount:
			subrange, err = idrange.NewRange(r.LowerVal, maxIntWithDigits(count))
		case upperDigitsCount:
			subrange, err = idrange.NewRange(minIntWithDigits(count), r.UpperVal)
		default:
			subrange, err = idrange.NewRange(minIntWithDigits(count), maxIntWithDigits(count))
		}

		if err != nil {
			panic("impossible range initialization error")
		}
		ret = append(ret, subrange)
	}

	return ret
}
