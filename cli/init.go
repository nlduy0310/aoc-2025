package cli

import (
	"fmt"
)

func InitDefaults() (*PartOption, *string, error) {
	partArg, err := GetArgument(0)
	if err != nil {
		return nil, nil, fmt.Errorf("can not read part option: %w", err)
	}
	partOption := PartOption(*partArg)
	if partOption != PartOne && partOption != PartTwo {
		return nil, nil, fmt.Errorf("invalid part option: %q", partOption)
	}

	input, err := ReadFileArgument(1)
	if err != nil {
		return nil, nil, fmt.Errorf("can not read input file: %w", err)
	}

	return &partOption, input, nil
}
