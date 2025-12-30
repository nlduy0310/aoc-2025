package grid

import (
	"fmt"

	"github.com/nlduy0310/aoc-2025/day04/location"
)

func ParseLines(lines []string) (*Grid, error) {
	if len(lines) == 0 {
		return nil, fmt.Errorf("empty input")
	}

	height := len(lines)
	width := len(lines[0])
	paperRolls := make(map[location.Location]struct{})
	for lineIdx, line := range lines {
		if len(line) != width {
			return nil, fmt.Errorf("expected line %d to have length %d, got %d", lineIdx+1, width, len(line))
		}
		for charIdx := 0; charIdx < len(line); charIdx++ {
			if line[charIdx] == '@' {
				loc := location.New(lineIdx, charIdx)
				paperRolls[loc] = struct{}{}
			}
		}
	}

	grid, err := New(width, height, paperRolls)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize grid: %w", err)
	}

	return grid, nil
}
