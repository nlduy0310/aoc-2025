package day01

import "fmt"

type direction int

const (
	Left direction = iota
	Right
)

var directionFromStringMap = map[string]direction{
	"L": Left,
	"R": Right,
}

func parseDirectionString(str string) (direction, error) {
	dir, ok := directionFromStringMap[str]
	if !ok {
		return 0, fmt.Errorf("unknown direction string")
	}

	return dir, nil
}
