package device

import (
	"strings"
)

// FromLine parse from something like 'ccc: ddd eee fff'
func FromLine(line string) (*Device, error) {
	parts := strings.Split(line, ": ")
	if len(parts) != 2 {
		return nil, invalidLineFormatError
	}

	name := parts[0]
	outputs := strings.Split(parts[1], " ")
	ret := Device{Name: name, Outputs: outputs}
	return &ret, nil
}
