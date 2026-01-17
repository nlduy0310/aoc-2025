package lightdiagram

import (
	"fmt"
	"strings"
)

const offIndicatorChar byte = '.'
const onIndicatorChar byte = '#'

// FromString parse from strings like '[.##.]'
func FromString(str string) (*LightDiagram, error) {
	str = strings.TrimPrefix(str, "[")
	str = strings.TrimSuffix(str, "]")

	size := len(str)
	ret, err := New(size)
	if err != nil {
		return nil, fmt.Errorf("can not initialize diagram: %w", err)
	}

	for idx := range size {
		ch := str[idx]
		switch ch {
		case offIndicatorChar:
		case onIndicatorChar:
			ret.Toggle(idx)
		default:
			return nil, fmt.Errorf("unknown character %q", ch)
		}
	}

	return ret, nil
}
