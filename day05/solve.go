package day05

import (
	"fmt"
	"slices"
	"strings"

	"github.com/nlduy0310/aoc-2025/day05/ingredient"
	"github.com/nlduy0310/aoc-2025/day05/ingredientrange"
)

func parseInput(inp string) ([]ingredientrange.IngredientRange, []ingredient.Ingredient, error) {
	lines := strings.Split(inp, "\n")

	ranges := make([]ingredientrange.IngredientRange, 0)
	ingredients := make([]ingredient.Ingredient, 0)
	phase := 0 // 0 -> parse ranges; 1 -> parse ingredients
	for lineIdx, line := range lines {
		if len(line) == 0 {
			if phase == 1 {
				break
			}
			phase++
			continue
		}

		if phase == 0 {
			irange, err := ingredientrange.FromString(line)
			if err != nil {
				return nil, nil, fmt.Errorf("can not parse range %q at line %d: %w", line, lineIdx+1, err)
			}
			ranges = append(ranges, irange)
		} else if phase == 1 {
			ingredient, err := ingredient.FromString(line)
			if err != nil {
				return nil, nil, fmt.Errorf("can not parse ingredient %q at line %d: %w", line, lineIdx+1, err)
			}
			ingredients = append(ingredients, ingredient)
		} else {
			panic("impossible state")
		}
	}

	return ranges, ingredients, nil
}

func SolvePartOne(inp string) (int, error) {
	iranges, ingredients, err := parseInput(inp)
	if err != nil {
		return 0, fmt.Errorf("failed to parse input file: %w", err)
	}

	freshCount := 0
	for _, ingredient := range ingredients {
		fresh := false

		for _, irange := range iranges {
			if irange.Contains(ingredient) {
				fresh = true
				break
			}
		}

		if fresh {
			freshCount++
		}
	}

	return freshCount, nil
}

func SolvePartTwo(inp string) (int, error) {
	iranges, _, err := parseInput(inp)
	if err != nil {
		return 0, fmt.Errorf("can not parse input file: %w", err)
	}

	slices.SortFunc(iranges, func(r1, r2 ingredientrange.IngredientRange) int {
		return int(r1.LowerVal - r2.LowerVal)
	})
	for i := 0; i < len(iranges) - 1; {
		cur, next := iranges[i], iranges[i+1]
		if cur.Contains(next.LowerVal) {
			newLower := cur.LowerVal
			newUpper := max(cur.UpperVal, next.UpperVal)
			iranges = append(iranges[:i], iranges[i+1:]...)
			iranges[i], _ = ingredientrange.New(newLower, newUpper)
		} else {
			i++
		}
	}

	sum := 0
	for _, irange := range iranges {
		sum += int(irange.UpperVal) - int(irange.LowerVal) + 1
		// println(irange.LowerVal, "-", irange.UpperVal)
	}

	return sum, nil
}
