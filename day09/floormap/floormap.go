package floormap

import (
	"github.com/nlduy0310/aoc-2025/common/gridpos"
)

type FloorMap struct {
	redTiles []gridpos.GridPos
}

func New(redTiles []gridpos.GridPos) (*FloorMap, error) {
	redTiles_ := make([]gridpos.GridPos, len(redTiles))
	copy(redTiles_, redTiles)
	ret := FloorMap{redTiles: redTiles_}
	return &ret, nil
}

func (fm FloorMap) RedTiles() []gridpos.GridPos {
	ret := make([]gridpos.GridPos, len(fm.redTiles))
	copy(ret, fm.redTiles)
	return ret
}
