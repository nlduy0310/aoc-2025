package idrange

import (
	"fmt"
	"strconv"
	"strings"
)

// ParseFromString parse one range from a string.
// input string must be in the format 'xxx-xxx'
func ParseFromString(inp string) (IdRange, error) {
	tokens := strings.Split(inp, "-")
	if len(tokens) != 2 {
		return IdRange{}, fmt.Errorf("expected 2 tokens separated by '-', got %d", len(tokens))
	}

	lowerVal, err := strconv.Atoi(tokens[0])
	if err != nil {
		return IdRange{}, fmt.Errorf("not a valid integer: %q", tokens[0])
	}
	upperVal, err := strconv.Atoi(tokens[1])
	if err != nil {
		return IdRange{}, fmt.Errorf("not a valid integer: %q", tokens[1])
	}

	ret, err := NewRange(lowerVal, upperVal)
	if err != nil {
		return IdRange{}, fmt.Errorf("failed to initialize range: %w", err)
	}

	return ret, nil
}

// ParseManyFromString parse many ranges from a string.
// input string must be in format 'xxx-xxx,xxx-xxx,...'
func ParseManyFromString(inp string) ([]IdRange, error) {
	tokens := strings.Split(inp, ",")
	ret := make([]IdRange, 0)

	for idx, token := range tokens {
		idrange, err := ParseFromString(token)
		if err != nil {
			return nil, fmt.Errorf("error parsing token #%d, %q: %w", idx+1, token, err)
		}
		ret = append(ret, idrange)
	}

	return ret, nil
}
