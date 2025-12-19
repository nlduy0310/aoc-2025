package day01

import (
	"fmt"
)

func SolvePartOne(input string) (int, error) {
	rotations, err := parseRotations(input)
	if err != nil {
		return 0, fmt.Errorf("can not parse rotations from input file: %w", err)
	}

	d, err := newDial(defaultStartingValue, defaultDialStart, defaultDialEnd)
	if err != nil {
		panic(fmt.Sprintf("can not construct dial from default values: %s", err))
	}

	ret := 0
	for _, rotation := range rotations {
		d.rotate(rotation)
		if d.currentValue == 0 {
			ret++
		}
	}

	return ret, nil
}
