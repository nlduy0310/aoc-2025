package gridpos

import "fmt"

type GridPos struct {
	Row int
	Col int
}

func New(row, col int) GridPos {
	return GridPos{Row: row, Col: col}
}

func (gp GridPos) String() string {
	return fmt.Sprintf("GridPos(Row=%d, Col=%d)", gp.Row, gp.Col)
}
