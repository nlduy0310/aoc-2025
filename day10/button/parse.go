package button

import (
	"fmt"
	"strconv"
	"strings"
)

// FromString parse from something like '(0,1,2,3,4)'
func FromString(str string) (*Button, error) {
	str = strings.TrimPrefix(str, "(")
	str = strings.TrimSuffix(str, ")")

	tokens := strings.Split(str, ",")
	lights := make([]int, 0, len(tokens))
	for _, token := range tokens {
		val, err := strconv.Atoi(token)
		if err != nil {
			return nil, fmt.Errorf("invalid int %q", token)
		}
		lights = append(lights, val)
	}

	ret, err := New(lights)
	if err != nil {
		return nil, fmt.Errorf("can not initialize button: %w", err)
	}

	return ret, nil
}
