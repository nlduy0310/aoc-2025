package shape

import (
	"fmt"
	"strconv"
	"strings"
)

func FromLines(lines []string) (*Shape, error) {
	if len(lines) < 2 {
		return nil, fmt.Errorf("require at least 2 lines")
	}

	indexTokens := strings.Split(lines[0], ":")
	if len(indexTokens) == 0 {
		return nil, fmt.Errorf("empty index line")
	}
	indexToken := indexTokens[0]
	index, err := strconv.Atoi(indexToken)
	if err != nil {
		return nil, fmt.Errorf("invalid index %q", indexToken)
	}

	width := len(lines[1])
	height := len(lines) - 1
	presentsCount := 0
	for _, line := range lines[1:] {
		for chIdx := range line {
			if ch := line[chIdx]; ch == '#' {
				presentsCount++
			}
		}
	}

	ret, err := New(index, width, height, presentsCount)
	if err != nil {
		return nil, fmt.Errorf("unable to initialize shape: %w", err)
	}
	return ret, nil
}
