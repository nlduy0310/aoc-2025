package day04

import (
	"fmt"
	"strings"

	"github.com/nlduy0310/aoc-2025/datastructures/set"
	"github.com/nlduy0310/aoc-2025/day04/grid"
	"github.com/nlduy0310/aoc-2025/day04/location"
)

const MAX_ADJACENT_PAPER_ROLLS int = 3

func parseInput(inp string) (*grid.Grid, error) {
	lines := strings.Split(inp, "\n")
	if len(lines) > 0 && len(lines[len(lines)-1]) == 0 {
		lines = lines[:len(lines)-1]
	}

	grid, err := grid.ParseLines(lines)
	if err != nil {
		return nil, fmt.Errorf("failed to parse input lines: %w", err)
	}

	return grid, nil
}

func findRemovables(g *grid.Grid, locs []location.Location) []location.Location {
	ret := make([]location.Location, 0)

	for _, loc := range locs {
		if !g.Contains(loc) || g.EmptyAt(loc) {
			continue
		}

		adjPaperRolls := 0
		for _, adjLoc := range loc.AdjacentLocations() {
			if !g.Contains(adjLoc) || g.EmptyAt(adjLoc) {
				continue
			}
			adjPaperRolls++
		}

		if adjPaperRolls <= MAX_ADJACENT_PAPER_ROLLS {
			ret = append(ret, loc)
		}
	}

	return ret
}

func SolvePartOne(inp string) (int, error) {
	grid, err := parseInput(inp)
	if err != nil {
		return 0, fmt.Errorf("failed to parse input file: %w", err)
	}

	removables := findRemovables(grid, grid.PaperRolls())

	return len(removables), nil
}

func SolvePartTwo(inp string) (int, error) {
	grid, err := parseInput(inp)
	if err != nil {
		return 0, fmt.Errorf("failed to parse input file: %w", err)
	}

	removeCount := 0
	var removables []location.Location
	var candidateLocations set.Set[location.Location]
	candidateLocations = set.FromList(grid.PaperRolls())
	for candidateLocations.Size() > 0 {
		removables = findRemovables(grid, candidateLocations.List())
		candidateLocations = set.New[location.Location]()
		for _, removableLoc := range removables {
			grid.RemoveAt(removableLoc)
			for _, adjLoc := range removableLoc.AdjacentLocations() {
				candidateLocations.Add(adjLoc)
			}
		}
		removeCount += len(removables)
	}
	// 1 get initial possible locations
	// 2 inspect possible locations -> removable location
	// 3 remove removable locations -> possible locations
	// 4 if len(possible locations) > 0 -> goes back to 2
	return removeCount, nil
}
