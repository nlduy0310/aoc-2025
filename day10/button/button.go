package button

import (
	"fmt"
	"iter"
	"strconv"
	"strings"

	"github.com/nlduy0310/aoc-2025/day10/lightdiagram"
)

type Button struct {
	indicatorLights []int
}

func (b Button) String() string {
	var builder strings.Builder
	builder.WriteString("Button(")
	for idx, v := range b.indicatorLights {
		builder.WriteString(strconv.Itoa(v))
		if idx < len(b.indicatorLights)-1 {
			builder.WriteString(",")
		}
	}
	builder.WriteString("Button)")
	return builder.String()
}

func New(indicatorLights []int) (*Button, error) {
	if len(indicatorLights) == 0 {
		return nil, fmt.Errorf("a button must specify at least 1 indicator light")
	}

	indicatorLights_ := make([]int, len(indicatorLights))
	copy(indicatorLights_, indicatorLights)
	ret := Button{indicatorLights: indicatorLights_}
	return &ret, nil
}

func (b Button) IterLights() iter.Seq2[int, int] {
	return func(yield func(int, int) bool) {
		for i := 0; i < len(b.indicatorLights); i++ {
			if !yield(i, b.indicatorLights[i]) {
				return
			}
		}
	}
}

func (b Button) Flip(ld lightdiagram.LightDiagram) error {
	for _, lightIdx := range b.indicatorLights {
		err := ld.Toggle(lightIdx)
		if err != nil {
			return fmt.Errorf("can not toggle light %d: %w", lightIdx, err)
		}
	}

	return nil
}
