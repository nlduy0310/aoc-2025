package problem

import (
	"fmt"

	"github.com/nlduy0310/aoc-2025/day06/op"
)

type Problem struct {
	args []int
	op.Op
}

func New(args []int, op op.Op) (*Problem, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("can not initialize with empty arguments")
	}

	ret := Problem{args, op}
	return &ret, nil
}

func (p Problem) Evaluate() int {
	if len(p.args) == 1 {
		return p.args[0]
	}

	ret := p.Op.Apply(p.args[0], p.args[1])
	for i := 2; i < len(p.args); i++ {
		ret = p.Op.Apply(ret, p.args[i])
	}
	return ret
}
