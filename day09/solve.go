package day09

import (
	"fmt"
	"math"
	"strings"

	"github.com/nlduy0310/aoc-2025/common/gridpos"
	"github.com/nlduy0310/aoc-2025/day09/floormap"
	"github.com/nlduy0310/aoc-2025/day09/rectilinearpolygon"
	"github.com/nlduy0310/aoc-2025/mathutils"
)

func parseInput(input string) (*floormap.FloorMap, error) {
	input = strings.TrimSuffix(input, "\n")

	ret, err := floormap.FromString(input)
	if err != nil {
		return nil, fmt.Errorf("can not parse input: %w", err)
	}

	return ret, nil
}

func axisAligned(pos1, pos2 gridpos.GridPos) bool {
	return pos1.Row == pos2.Row || pos1.Col == pos2.Col
}

func rectArea(pos1, pos2 gridpos.GridPos) int {
	return (mathutils.AbsInt(pos1.Row-pos2.Row) + 1) *
		(mathutils.AbsInt(pos1.Col-pos2.Col) + 1)
}

func rectVerticesOrdered(corner1, corner2 gridpos.GridPos) (
	gridpos.GridPos, gridpos.GridPos,
	gridpos.GridPos, gridpos.GridPos,
) {
	minRow, maxRow := mathutils.MinMax(corner1.Row, corner2.Row)
	minCol, maxCol := mathutils.MinMax(corner1.Col, corner2.Col)
	return gridpos.New(minRow, minCol), gridpos.New(minRow, maxCol),
		gridpos.New(maxRow, maxCol), gridpos.New(maxRow, minCol)
}

func SolvePartOne(input string) (int, error) {
	floor, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	redTiles := floor.RedTiles()
	if len(redTiles) <= 1 {
		return 0, fmt.Errorf("require at least 2 red tiles")
	}

	maxArea := math.MinInt
	for i := 0; i < len(redTiles)-1; i++ {
		for j := i + 1; j < len(redTiles); j++ {
			if area := rectArea(redTiles[i], redTiles[j]); area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea, nil
}

func SolvePartTwo(input string) (int, error) {
	floor, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	redTiles := floor.RedTiles()
	if len(redTiles) <= 1 {
		return 0, fmt.Errorf("require at least 2 red tiles")
	}

	polygon, err := rectilinearpolygon.New(redTiles)
	if err != nil {
		return 0, fmt.Errorf("can not initialize rectilinear polygon: %w", err)
	}

	maxArea := math.MinInt
	for i := 0; i < len(redTiles)-1; i++ {
		for j := i + 1; j < len(redTiles); j++ {
			corner1, corner2 := redTiles[i], redTiles[j]
			var contained bool
			if axisAligned(corner1, corner2) {
				contained = polygon.ContainsPoint(corner1) &&
					polygon.ContainsPoint(corner2) &&
					polygon.ContainsAxisAlignedLine(corner1, corner2)
			} else {
				vert1, vert2, vert3, vert4 := rectVerticesOrdered(corner1, corner2)
				contained = polygon.ContainsPoint(vert1) &&
					polygon.ContainsPoint(vert2) &&
					polygon.ContainsPoint(vert3) &&
					polygon.ContainsPoint(vert4) &&
					polygon.ContainsAxisAlignedLine(vert1, vert2) &&
					polygon.ContainsAxisAlignedLine(vert2, vert3) &&
					polygon.ContainsAxisAlignedLine(vert3, vert4) &&
					polygon.ContainsAxisAlignedLine(vert4, vert1)
			}

			if !contained {
				continue
			}
			if area := rectArea(corner1, corner2); area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea, nil
}
