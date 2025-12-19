package cli

import (
	"fmt"
	"os"
)

// ReadFileArgument reads from the file specified by the CLI argument at the provided index
func ReadFileArgument(idx int) (*string, error) {
	path, err := GetArgument(idx)
	if err != nil {
		return nil, fmt.Errorf("can not get CLI argument at index %d: %w", idx, err)
	}

	content, err := os.ReadFile(*path)
	if err != nil {
		return nil, fmt.Errorf("can not read file %q: %w", *path, err)
	}

	ret := string(content)
	return &ret, nil
}
