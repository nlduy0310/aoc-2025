package coordinates

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nlduy0310/aoc-2025/valueutils"
)

func FromString(str string) (Coordinates, error) {
	tokens := strings.Split(str, ",")
	if len(tokens) != 3 {
		return valueutils.ZeroValue[Coordinates](), fmt.Errorf("expected 3 tokens, got %d", len(tokens))
	}

	x, err := strconv.Atoi(tokens[0])
	if err != nil {
		return valueutils.ZeroValue[Coordinates](), fmt.Errorf("%q is not a valid int", tokens[0])
	}
	y, err := strconv.Atoi(tokens[1])
	if err != nil {
		return valueutils.ZeroValue[Coordinates](), fmt.Errorf("%q is not a valid int", tokens[1])
	}
	z, err := strconv.Atoi(tokens[2])
	if err != nil {
		return valueutils.ZeroValue[Coordinates](), fmt.Errorf("%q is not a valid int", tokens[2])
	}

	return New(x, y, z), nil
}
