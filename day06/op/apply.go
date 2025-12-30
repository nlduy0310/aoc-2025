package op

import "fmt"

type applyFunc func(int, int) int

var applyFuncByOp = map[Op]applyFunc{
	Add: func(a, b int) int { return a + b },
	Mul: func(a, b int) int { return a * b },
}

func (o Op) Apply(a, b int) int {
	if apply, ok := applyFuncByOp[o]; ok {
		return apply(a, b)
	} else {
		panic(fmt.Errorf("unrecognized operation enum %v", o))
	}
}
