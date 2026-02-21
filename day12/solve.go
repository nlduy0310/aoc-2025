package day12

import (
	"fmt"
	"math"
	"slices"
	"strings"

	"github.com/nlduy0310/aoc-2025/day12/region"
	"github.com/nlduy0310/aoc-2025/day12/shape"
)

func parseInput(input string) ([]shape.Shape, []region.Region, error) {
	input = strings.TrimSuffix(input, "\n")
	parts := strings.Split(input, "\n\n")
	if len(parts) < 2 {
		return nil, nil, fmt.Errorf("expect at least 2 parts separated by 2 newlines")
	}

	shapeParts := parts[:len(parts)-1]
	shapes := make([]shape.Shape, 0, len(shapeParts))
	for shapePartIdx, shapePart := range shapeParts {
		s, err := shape.FromLines(strings.Split(shapePart, "\n"))
		if err != nil {
			return nil, nil, fmt.Errorf("failed to parse shape part #%d: %w", shapePartIdx, err)
		}
		shapes = append(shapes, *s)
	}

	regionLines := strings.Split(parts[len(parts)-1], "\n")
	regions := make([]region.Region, 0, len(regionLines))
	for regionLineIdx, regionLine := range regionLines {
		r, err := region.FromLine(regionLine)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to parse region line #%d: %w", regionLineIdx, err)
		}
		regions = append(regions, *r)
	}

	return shapes, regions, nil
}

func validateInput(shapes []shape.Shape, regions []region.Region) error {
	maxIdx := math.MinInt
	for shapeIdx, s := range shapes {
		if s.Height != s.Width || s.Height != 3 {
			return fmt.Errorf("shape #%d is not bound by a square of 3", shapeIdx)
		}
		if s.Index > maxIdx {
			maxIdx = s.Index
		}
	}

	for regionIdx, r := range regions {
		if maxIdx >= len(r.ShapeCounts) {
			return fmt.Errorf("region #%d uses %d shapes, got %d", regionIdx, len(r.ShapeCounts), maxIdx+1)
		}
	}

	return nil
}

func SolvePartOne(input string) (int, error) {
	shapes, regions, err := parseInput(input)
	if err != nil {
		return 0, fmt.Errorf("unable to parse input: %w", err)
	} else if err = validateInput(shapes, regions); err != nil {
		return 0, fmt.Errorf("unable to validate input: %w", err)
	}

	slices.SortFunc(shapes, func(s1, s2 shape.Shape) int {
		return s1.Index - s2.Index
	})

	ret := 0
	for _, r := range regions {
		maxArea := r.Area()
		area := 0
		for shapeIdx, count := range r.ShapeCounts {
			area += count * shapes[shapeIdx].Area()
		}
		if area <= maxArea {
			ret++ // ?
		}
	}
	return ret, nil
}

func SolvePartTwo(input string) (int, error) {
	panic("merry christmas")
}
