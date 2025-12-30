package grid

import (
	"fmt"

	"github.com/nlduy0310/aoc-2025/day04/location"
)

type Grid struct {
	Width int
	Height int
	paperRolls map[location.Location]struct{}
}

func New(width, height int, paperRolls map[location.Location]struct{}) (*Grid, error) {
	if width <= 0 {
		return nil, fmt.Errorf("width must be positive")
	}
	if height <= 0 {
		return nil, fmt.Errorf("height must be positive")
	}

	if paperRolls == nil {
		paperRolls = make(map[location.Location]struct{})
	} else {
		for loc := range paperRolls {
			if !(loc.Col >= 0 && loc.Col < width && loc.Row >= 0 && loc.Row < height) {
				return nil, fmt.Errorf("paper roll at row %d, col %d is out of bound", loc.Row, loc.Col)
			}
		}
	}

	ret := Grid{Width: width, Height: height, paperRolls: paperRolls}
	return &ret, nil
}

func (g Grid) Contains(l location.Location) bool {
	return l.Col >= 0 && l.Col < g.Width && l.Row >= 0 && l.Row < g.Height
}

func (g Grid) EmptyAt(l location.Location) bool {
	_, ok := g.paperRolls[l]
	return !ok
}

func (g *Grid) RemoveAt(l location.Location) error {
	if _, ok := g.paperRolls[l]; ok {
		delete(g.paperRolls, l)
		return nil
	}

	return fmt.Errorf("location is not a paper roll")
}

func (g Grid) PaperRolls() []location.Location {
	ret := make([]location.Location, 0, len(g.paperRolls))
	for loc := range g.paperRolls {
		ret = append(ret, loc)
	}

	return ret
}
