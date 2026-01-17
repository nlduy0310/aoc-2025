package rectilinearpolygon

import (
	"github.com/nlduy0310/aoc-2025/common/gridpos"
	"github.com/nlduy0310/aoc-2025/mathutils"
)

func isHorizontalEdge(vertex1, vertex2 gridpos.GridPos) bool {
	return vertex1.Row == vertex2.Row
}

func edgeContains(edgeVertex1, edgeVertex2 gridpos.GridPos, point gridpos.GridPos) bool {
	if isHorizontalEdge(edgeVertex1, edgeVertex2) {
		return point.Row == edgeVertex1.Row &&
			mathutils.InRangeUnordered(edgeVertex1.Col, edgeVertex2.Col, point.Col, true, true)
	} else {
		return point.Col == edgeVertex1.Col &&
			mathutils.InRangeUnordered(edgeVertex1.Row, edgeVertex2.Row, point.Row, true, true)
	}
}

func (rp RectilinearPolygon) ContainsPoint(point gridpos.GridPos) bool {
	// count polygon crossings of the ray to the right
	// unless we're on an edge
	crossings := 0
	for vertexIdx, edgeStart := range rp.vertices {
		edgeEnd := rp.vertices[(vertexIdx+1)%len(rp.vertices)]

		if edgeContains(edgeStart, edgeEnd, point) {
			return true
		}

		if isHorizontalEdge(edgeStart, edgeEnd) {
			continue
		}

		minRow, maxRow := mathutils.MinMax(edgeStart.Row, edgeEnd.Row)
		if point.Col < edgeStart.Col &&
			mathutils.InRange(minRow, maxRow, point.Row, true, false) {
			crossings++
		}
	}

	return crossings%2 == 1
}

func (rp RectilinearPolygon) ContainsAxisAlignedLine(lineStart, lineEnd gridpos.GridPos) bool {
	if !(rp.ContainsPoint(lineStart) && rp.ContainsPoint(lineEnd)) {
		return false
	}

	lineHorizontal := isHorizontalEdge(lineStart, lineEnd)

	for vertexIdx, edgeStart := range rp.vertices {
		edgeEnd := rp.vertices[(vertexIdx+1)%len(rp.vertices)]

		edgeHorizontal := isHorizontalEdge(edgeStart, edgeEnd)
		if lineHorizontal == edgeHorizontal {
			continue
		}

		if lineHorizontal {
			if mathutils.InRangeUnordered(edgeStart.Row, edgeEnd.Row, lineStart.Row, false, false) &&
				mathutils.InRangeUnordered(lineStart.Col, lineEnd.Col, edgeStart.Col, false, false) {
				return false
			}
		} else {
			if mathutils.InRangeUnordered(edgeStart.Col, edgeEnd.Col, lineStart.Col, false, false) &&
				mathutils.InRangeUnordered(lineStart.Row, lineEnd.Row, edgeStart.Row, false, false) {
				return false
			}
		}
	}

	return true
}
