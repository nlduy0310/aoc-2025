package shape

import (
	"fmt"
)

type Shape struct {
	Index         int
	Width         int
	Height        int
	PresentsCount int
}

func New(index, width, height, presentsCount int) (*Shape, error) {
	if index < 0 {
		return nil, fmt.Errorf("index must be non-negative")
	} else if width <= 0 {
		return nil, fmt.Errorf("width must be positive")
	} else if height <= 0 {
		return nil, fmt.Errorf("height must be positive")
	} else if presentsCount < 0 {
		return nil, fmt.Errorf("presents count must be non-negative")
	}

	ret := Shape{index, width, height, presentsCount}
	return &ret, nil
}

func (s Shape) Area() int {
	return s.Width * s.Height
}
