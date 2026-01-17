package rectilinearpolygon

import (
	"fmt"

	"github.com/nlduy0310/aoc-2025/common/gridpos"
)

type RectilinearPolygon struct {
	vertices []gridpos.GridPos
}

func New(vertices []gridpos.GridPos) (*RectilinearPolygon, error) {
	if len(vertices) < 4 {
		return nil, fmt.Errorf("a rectilinear polygon must have at least 4 vertices, got %d", len(vertices))
	}
	if len(vertices)%2 != 0 {
		return nil, fmt.Errorf("a rectilinear polygon must have an even number of vertices, got %d", len(vertices))
	}

	for vertexIdx := range vertices {
		var nextVertexIdx int
		if vertexIdx < len(vertices)-1 {
			nextVertexIdx = vertexIdx + 1
		} else {
			nextVertexIdx = 0
		}

		if vertices[vertexIdx].Row != vertices[nextVertexIdx].Row && vertices[vertexIdx].Col != vertices[nextVertexIdx].Col {
			return nil, fmt.Errorf("edge %s -> %s is neither on the same row nor column", vertices[vertexIdx], vertices[nextVertexIdx])
		}
	}

	vertices_ := make([]gridpos.GridPos, len(vertices))
	copy(vertices_, vertices)
	ret := RectilinearPolygon{vertices: vertices_}
	return &ret, nil
}
