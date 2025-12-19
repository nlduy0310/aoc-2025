package day01

import (
	"fmt"
	"strconv"
)

type rotation struct {
	direction
	distance int
}

func newRotation(dir direction, dist int) (rotation, error) {
	if dist < 0 {
		return rotation{Left, 0}, fmt.Errorf("negative distance not allowed: %d", dist)
	}

	return rotation{dir, dist}, nil
}

func parseRotationString(str string) (rotation, error) {
	if len(str) < 2 {
		return rotation{Left, 0}, fmt.Errorf("string too short")
	}

	directionString := string(str[0])
	dir, err := parseDirectionString(directionString)
	if err != nil {
		return rotation{Left, 0}, fmt.Errorf("can not parse direction string %q: %w", directionString, err)
	}

	distanceString := str[1:]
	distance, err := strconv.Atoi(distanceString)
	if err != nil {
		return rotation{Left, 0}, fmt.Errorf("can not parse distance string %q: %w", distanceString, err)
	}

	return rotation{dir, distance}, nil
}
