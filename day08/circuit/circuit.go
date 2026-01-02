package circuit

import (
	"fmt"

	"github.com/nlduy0310/aoc-2025/day08/coordinates"
)

type Circuit struct {
	boxesMap map[coordinates.Coordinates]struct{}
}

func New() Circuit {
	return Circuit{boxesMap: make(map[coordinates.Coordinates]struct{})}
}

func FromBoxes(boxes ...coordinates.Coordinates) Circuit {
	ret := New()
	for _, box := range boxes {
		ret.Add(box)
	}
	return ret
}

func (c Circuit) Has(boxCoords coordinates.Coordinates) bool {
	_, ok := c.boxesMap[boxCoords]
	return ok
}

func (c *Circuit) Add(boxCoords coordinates.Coordinates) error {
	if c.Has(boxCoords) {
		return fmt.Errorf("circuit already has this box")
	}

	c.boxesMap[boxCoords] = struct{}{}
	return nil
}

func (c Circuit) Size() int {
	return len(c.boxesMap)
}

func (c Circuit) Boxes() []coordinates.Coordinates {
	ret := make([]coordinates.Coordinates, 0, len(c.boxesMap))
	for box := range c.boxesMap {
		ret = append(ret, box)
	}
	return ret
}

func Join(c1, c2 Circuit) Circuit {
	ret := New()
	for box := range c1.boxesMap {
		ret.Add(box)
	}
	for box := range c2.boxesMap {
		ret.Add(box)
	}
	return ret
}
