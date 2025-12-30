package location

func (l Location) Top() Location {
	return New(l.Row - 1, l.Col)
}

func (l Location) Bottom() Location {
	return New(l.Row + 1, l.Col)
}

func (l Location) Left() Location {
	return New(l.Row, l.Col - 1)
}

func (l Location) Right() Location {
	return New(l.Row, l.Col + 1)
}

func (l Location) TopLeft() Location {
	return l.Top().Left()
}

func (l Location) TopRight() Location {
	return l.Top().Right()
}

func (l Location) BottomLeft() Location {
	return l.Bottom().Left()
}

func (l Location) BottomRight() Location {
	return l.Bottom().Right()
}

func (l Location) AdjacentLocations() []Location {
	return []Location{
		l.Top(), l.Right(), l.Bottom(), l.Left(),
		l.TopLeft(), l.TopRight(), l.BottomRight(), l.BottomLeft(),
	}
}
