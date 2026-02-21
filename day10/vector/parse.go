package vector

import (
	"fmt"
	"strconv"
	"strings"
)

func FromCSVLine(str string) (*Vector, error) {
	str = strings.TrimSuffix(str, "\n")
	tokens := strings.Split(str, ",")
	elements := make([]int, len(tokens))
	for idx, token := range tokens {
		if val, err := strconv.Atoi(token); err != nil {
			return nil, fmt.Errorf("invalid int token %q", token)
		} else {
			elements[idx] = val
		}
	}
	ret := New(elements)
	return &ret, nil
}
