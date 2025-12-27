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
		subranges := splitRangeByDigitsCount(idRange)
		for _, subrange := range subranges {
			totalDigits := mathutils.CountDigits(subrange.LowerVal)
			if totalDigits%2 == 1 {
				continue
			}

			halfDigits := totalDigits / 2
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

func SolvePartTwo(inp string) (int, error) {
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

	invalidIds := make(map[int]struct{})
	for _, idRange := range idRanges {
		subRanges := splitRangeByDigitsCount(idRange)
		for _, subRange := range subRanges {
			totalDigits := mathutils.CountDigits(subRange.LowerVal)
			repeats := mathutils.Divisors(int(totalDigits))
			for _, repeat := range repeats {
				if repeat == 1 {
					continue
				}

				digits := int(totalDigits) / repeat
				multiplier := 0
				for r := range repeat {
					multiplier += int(math.Pow10(r * digits))
				}
				leftBound := int(math.Ceil(float64(subRange.LowerVal) / float64(multiplier)))
				rightBound := int(math.Floor(float64(subRange.UpperVal) / float64(multiplier)))

				for m := leftBound; m <= rightBound; m++ {
					// lazy tracking, cause I don't know how to handle 222-222 vs 22-22-22
					invalidIds[multiplier*m] = struct{}{}
				}
			}
		}
	}

	sum := 0
	for invalidId := range invalidIds {
		sum += invalidId
	}

	return sum, nil
}
