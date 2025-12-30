package ingredientrange

import (
	"fmt"
	"strings"

	"github.com/nlduy0310/aoc-2025/day05/ingredient"
	"github.com/nlduy0310/aoc-2025/valueutils"
)

func FromString(str string) (IngredientRange, error) {
	tokens := strings.Split(str, "-")
	if len(tokens) != 2 {
		return valueutils.ZeroValue[IngredientRange](), fmt.Errorf("expected 2 tokens separated by %q, got %d", "-", len(tokens))
	}

	lower, err := ingredient.FromString(tokens[0])
	if err != nil {
		return valueutils.ZeroValue[IngredientRange](), fmt.Errorf("can not parse ingredient %q: %w", tokens[0], err)
	}
	upper, err := ingredient.FromString(tokens[1])
	if err != nil {
		return valueutils.ZeroValue[IngredientRange](), fmt.Errorf("can not parse ingredient %q: %w", tokens[1], err)
	}

	ret, err := New(lower, upper)
	if err != nil {
		return valueutils.ZeroValue[IngredientRange](), fmt.Errorf("can not initialize range: %w", err)
	}

	return ret, nil
}
