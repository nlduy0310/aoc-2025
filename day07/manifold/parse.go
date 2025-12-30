package manifold

import (
	"fmt"

	"github.com/nlduy0310/aoc-2025/day07/location"
)

func ParseLines(lines []string) (*Manifold, error) {
	if len(lines) == 0 {
		return nil, fmt.Errorf("empty input")
	}

	height := len(lines)
	width := len(lines[0])
	splitters := make(map[location.Location]struct{})
	var startLocation *location.Location = nil

	for lineIdx, line := range lines {
		if len(line) != width {
			return nil, fmt.Errorf("expected line %d to have length %d, got %d", lineIdx+1, width, len(line))
		}
		for colIdx := range len(line) {
			ch := line[colIdx]
			if ch == '.' {
				continue
			} else if ch == '^' {
				splitters[location.New(lineIdx, colIdx)] = struct{}{}
			} else if ch == 'S' {
				if startLocation != nil {
					return nil, fmt.Errorf("multiple starting locations found")
				}
				loc := location.New(lineIdx, colIdx)
				startLocation = &loc
			} else {
				return nil, fmt.Errorf("found unknown character %q", ch)
			}
		}
	}

	if startLocation == nil {
		return nil, fmt.Errorf("found no starting locations")
	}

	ret := Manifold{
		Width: width, Height: height,
		StartLocation: *startLocation,
		splitterMap: splitters,
	}
	return &ret, nil
}
