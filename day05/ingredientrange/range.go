package ingredientrange

import (
	"fmt"

	"github.com/nlduy0310/aoc-2025/day05/ingredient"
	"github.com/nlduy0310/aoc-2025/valueutils"
)

type IngredientRange struct {
	LowerVal ingredient.Ingredient
	UpperVal ingredient.Ingredient
}

func New(lower, upper ingredient.Ingredient) (IngredientRange, error) {
	if lower > upper {
		return valueutils.ZeroValue[IngredientRange](), fmt.Errorf("invalid range: lower (%d) > upper (%d)", lower, upper)
	}

	return IngredientRange{lower, upper}, nil
}

func (r IngredientRange) Contains(i ingredient.Ingredient) bool {
	return i >= r.LowerVal && i <= r.UpperVal
}
