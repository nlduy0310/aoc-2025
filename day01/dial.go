package day01

import "fmt"

type dialRange struct {
	start int
	end   int
}

func newDialRange(start, end int) (*dialRange, error) {
	if start > end {
		return nil, fmt.Errorf("invalid range: start (%d) > end(%d)", start, end)
	}

	return &dialRange{start, end}, nil
}

func (r dialRange) size() int {
	return r.end - r.start + 1
}

type dial struct {
	currentValue int
	dialRange
}

func newDial(value, dialStart, dialEnd int) (*dial, error) {
	dRange, err := newDialRange(dialStart, dialEnd)
	if err != nil {
		return nil, fmt.Errorf("invalid dial range: %w", err)
	}

	if value < dialStart || value > dialEnd {
		return nil, fmt.Errorf("initial value %d out of provided range [%d, %d]", value, dialStart, dialEnd)
	}

	return &dial{value, *dRange}, nil
}

func (d *dial) rotate(r rotation) {
	var sign int
	if r.direction == Left {
		sign = -1
	} else {
		sign = 1
	}

	d.currentValue += sign * r.distance

	rsize := d.dialRange.size()
	for d.currentValue > d.dialRange.end {
		d.currentValue -= rsize
	}
	for d.currentValue < d.dialRange.start {
		d.currentValue += rsize
	}
}
