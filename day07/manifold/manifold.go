package manifold

import "github.com/nlduy0310/aoc-2025/day07/location"

type Manifold struct {
	Width int
	Height int
	StartLocation location.Location
	splitterMap map[location.Location]struct{}
}

func (m Manifold) Contains(loc location.Location) bool {
	return loc.Row >= 0 && loc.Row < m.Height && loc.Col >= 0 && loc.Col < m.Width
}

func (m Manifold) HasSplitter(loc location.Location) bool {
	if !m.Contains(loc) {
		return false
	}

	_, ok := m.splitterMap[loc]
	return ok
}
