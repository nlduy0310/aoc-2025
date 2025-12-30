package ingredient

import (
	"fmt"
	"strconv"

	"github.com/nlduy0310/aoc-2025/valueutils"
)

func FromString(str string) (Ingredient, error) {
	id, err := strconv.Atoi(str)
	if err != nil {
		return valueutils.ZeroValue[Ingredient](), fmt.Errorf("not a valid ingredient ID")
	}

	return Ingredient(id), nil
}
