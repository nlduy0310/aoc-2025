package day01

import (
	"bufio"
	"fmt"
	"strings"
)

func parseRotations(input string) ([]rotation, error) {
	rotations := make([]rotation, 0)
	scanner := bufio.NewScanner(strings.NewReader(input))
	lineIdx := 0
	for scanner.Scan() {
		lineIdx++
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		rotation, err := parseRotationString(line)
		if err != nil {
			return nil, fmt.Errorf("failed to parse rotation from line %d %q: %w", lineIdx, line, err)
		}
		rotations = append(rotations, rotation)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error scanning file content: %w", err)
	}

	return rotations, nil
}
