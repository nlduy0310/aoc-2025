package region

type Region struct {
	Width       int
	Height      int
	ShapeCounts []int
}

func New(width, height int, shapeCounts []int) Region {
	return Region{width, height, shapeCounts}
}

func (r Region) Area() int {
	return r.Width * r.Height
}
