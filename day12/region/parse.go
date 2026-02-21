package region

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func FromLine(line string) (*Region, error) {
	re := regexp.MustCompile(`^(\d+)x(\d+):\s([\s\d]+)\n?$`)
	matches := re.FindStringSubmatch(line)
	if matches == nil {
		return nil, fmt.Errorf("invalid line format")
	}

	widthToken := matches[1]
	heightToken := matches[2]
	countsToken := matches[3]

	width, err := strconv.Atoi(widthToken)
	if err != nil {
		return nil, fmt.Errorf("invalid width token %q", widthToken)
	}
	height, err := strconv.Atoi(heightToken)
	if err != nil {
		return nil, fmt.Errorf("invalid height token %q", heightToken)
	}
	countTokens := strings.Split(countsToken, " ")
	counts := make([]int, 0, len(countTokens))
	for _, countToken := range countTokens {
		count, err := strconv.Atoi(countToken)
		if err != nil {
			return nil, fmt.Errorf("invalid count token: %q", countToken)
		}
		counts = append(counts, count)
	}

	ret := New(width, height, counts)
	return &ret, nil
}
