package day01

import (
	"fmt"

	"github.com/nlduy0310/aoc-2025/mathutils"
)

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

func (d *dial) rotate_count_clicks(r rotation, val int) int {
	rsize := d.dialRange.size()
	absDist := mathutils.AbsInt(r.distance)
	var sign int
	if r.direction == Left {
		sign = -1
	} else {
		sign = 1
	}

	clicks := absDist / rsize
	remainingDist := absDist % rsize
	next := d.currentValue + sign*remainingDist
	if remainingDist != 0 && d.currentValue != val {
		before := val - d.currentValue
		after := val - next
		prod := before * after
		if prod <= 0 || next <= val-rsize || next >= val+rsize {
			clicks++
		}
	}

	if next > d.dialRange.end {
		next -= rsize
	} else if next < d.dialRange.start {
		next += rsize
	}
	d.currentValue = next

	return clicks
}
