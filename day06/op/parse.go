package op

import (
	"fmt"

	"github.com/nlduy0310/aoc-2025/valueutils"
)

var opByString = map[string]Op {
	"+": Add,
	"*": Mul,
}

func FromString(str string) (Op, error) {
	if op, ok := opByString[str]; ok {
		return op, nil
	} else {
		return valueutils.ZeroValue[Op](), fmt.Errorf("invalid operator string")
	}
}
