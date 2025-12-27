package mathutils

import "math"

func CountDigits(i int) uint {
	if i == 0 {
		return 1
	}

	if i < 0 {
		i = -i
	}

	return uint(math.Floor(math.Log10(float64(i))) + 1)
}
