package day01

import "fmt"

func SolvePartTwo(input string) (int, error) {
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
		clicks := d.rotate_count_clicks(rotation, defaultClickTarget)
		ret += clicks
	}

	return ret, nil
}
