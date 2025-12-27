package day02

import (
	"fmt"
	"math"
	"strings"

	"github.com/nlduy0310/aoc-2025/day02/idrange"
	"github.com/nlduy0310/aoc-2025/mathutils"
)

func SolvePartOne(inp string) (int, error) {
	inp = strings.TrimSuffix(inp, "\n")

	idRanges, err := idrange.ParseManyFromString(inp)
	if err != nil {
		return 0, fmt.Errorf("failed to parse input: %w", err)
	}

	// assuming all ranges are positive
	for _, idRange := range idRanges {
		if idRange.LowerVal < 0 || idRange.UpperVal < 0 {
			panic("negative range found")
		}
	}

	sum := 0
	for _, idRange := range idRanges {
		subranges := SplitRangeByDigitsCount(idRange)
		for _, subrange := range subranges {
			digits := mathutils.CountDigits(subrange.LowerVal)
			if digits%2 == 1 {
				continue
			}

			halfDigits := digits / 2
			multiplier := int(math.Pow10(int(halfDigits))) + 1
			leftBound := int(math.Ceil(float64(subrange.LowerVal) / float64(multiplier)))
			rightBound := int(math.Floor(float64(subrange.UpperVal) / float64(multiplier)))
			// for i := leftBound; i <= rightBound; i++ {
			// 	sum += i * multiplier
			// }
			leftSum := leftBound * (leftBound + 1) / 2
			rightSum := rightBound * (rightBound + 1) / 2
			sum += multiplier * (rightSum - leftSum + leftBound)
		}
	}

	return sum, nil
}
