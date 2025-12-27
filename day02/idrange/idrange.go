package idrange

import "fmt"

type IdRange struct {
	LowerVal int
	UpperVal int
}

func NewRange(lowerVal, upperVal int) (IdRange, error) {
	if lowerVal > upperVal {
		return IdRange{}, fmt.Errorf("invalid range: lower (%d) > upper (%d)", lowerVal, upperVal)
	}

	return IdRange{lowerVal, upperVal}, nil
}
