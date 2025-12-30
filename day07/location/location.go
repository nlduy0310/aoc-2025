package location

type Location struct {
	Row int
	Col int
}

func New(row, col int) Location {
	return Location{row, col}
}
