package coordinates

import "fmt"

type Coordinates struct {
	X, Y, Z int
}

func New(x, y, z int) Coordinates {
	return Coordinates{X: x, Y: y, Z: z}
}

func (c Coordinates) String() string {
	return fmt.Sprintf("(%d, %d, %d)", c.X, c.Y, c.Z)
}
