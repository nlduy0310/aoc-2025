package floormap

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nlduy0310/aoc-2025/common/gridpos"
)

func FromString(str string) (*FloorMap, error) {
	lines := strings.Split(str, "\n")
	redTiles := make([]gridpos.GridPos, 0, len(lines))
	for lineIdx, line := range lines {
		tokens := strings.Split(line, ",")
		if len(tokens) != 2 {
			return nil, fmt.Errorf("expected line %d to have 2 tokens, got %d", lineIdx, len(tokens))
		}

		colStr, rowStr := tokens[0], tokens[1]
		colIdx, err := strconv.Atoi(colStr)
		if err != nil {
			return nil, fmt.Errorf("invalid int at line %d: %q", lineIdx, colStr)
		}
		rowIdx, err := strconv.Atoi(rowStr)
		if err != nil {
			return nil, fmt.Errorf("invalid int at line %d: %q", lineIdx, rowStr)
		}
		tile := gridpos.New(rowIdx, colIdx)
		redTiles = append(redTiles, tile)
	}

	ret, err := New(redTiles)
	if err != nil {
		return nil, fmt.Errorf("can not initialize floor map: %w", err)
	}
	return ret, nil
}
