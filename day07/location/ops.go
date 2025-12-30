package location

func (l Location) Top() Location {
	return New(l.Row-1, l.Col)
}

func (l Location) Bottom() Location {
	return New(l.Row+1, l.Col)
}

func (l Location) Left() Location {
	return New(l.Row, l.Col-1)
}

func (l Location) Right() Location {
	return New(l.Row, l.Col+1)
}
