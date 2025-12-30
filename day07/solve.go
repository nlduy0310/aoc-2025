package day07

import (
	"fmt"

	"github.com/nlduy0310/aoc-2025/datastructures/set"
	"github.com/nlduy0310/aoc-2025/day07/location"
	"github.com/nlduy0310/aoc-2025/day07/manifold"
	"github.com/nlduy0310/aoc-2025/stringutils"
)

func SolvePartOne(inp string) (int, error) {
	lines := stringutils.SplitNonEmpty(inp, "\n")
	m, err := manifold.ParseLines(lines)
	if err != nil {
		return 0, fmt.Errorf("failed to parse manifold from file: %w", err)
	}

	tachyons := []location.Location{m.StartLocation}
	splitLocations := set.New[location.Location]()
	for len(tachyons) > 0 {
		nextTachyons := set.New[location.Location]()
		for _, tachyon := range tachyons {
			if !m.Contains(tachyon) {
				continue
			}
			if m.HasSplitter(tachyon) {
				splitLocations.Add(tachyon)
				nextTachyons.Add(tachyon.Left())
				nextTachyons.Add(tachyon.Right())
			} else {
				nextTachyons.Add(tachyon.Bottom())
			}
		}
		tachyons = nextTachyons.List()
	}

	return splitLocations.Size(), nil
}

func SolvePartTwo(inp string) (int, error) {
	lines := stringutils.SplitNonEmpty(inp, "\n")
	m, err := manifold.ParseLines(lines)
	if err != nil {
		return 0, fmt.Errorf("failed to parse manifold from file: %w", err)
	}

	totalTimelines := 0
	tachyons := map[location.Location]int{m.StartLocation: 1}
	for len(tachyons) > 0 {
		nextTachyons := make(map[location.Location]int)
		for tachyonLoc, count := range tachyons {
			if !m.Contains(tachyonLoc) {
				totalTimelines += count
			} else if m.HasSplitter(tachyonLoc) {
				nextTachyons[tachyonLoc.Left()] += count
				nextTachyons[tachyonLoc.Right()] += count
			} else {
				nextTachyons[tachyonLoc.Bottom()] += count
			}
		}
		tachyons = nextTachyons
	}

	return totalTimelines, nil
}
