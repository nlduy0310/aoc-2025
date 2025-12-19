package cli

import (
	"fmt"
	"os"
)

// GetArgument gets the CLI argument at the provided index
func GetArgument(idx int) (*string, error) {
	if idx < 0 {
		return nil, fmt.Errorf("invalid index")
	}

	args := os.Args[1:]
	if idx > len(args)-1 {
		return nil, fmt.Errorf("argument not provided")
	}

	return &args[idx], nil
}
