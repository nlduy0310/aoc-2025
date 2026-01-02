package coordinates

import "math"

func EuclideanDistance(c1, c2 Coordinates) float64 {
	return math.Sqrt(
		math.Pow(float64(c1.X - c2.X), 2) +
		math.Pow(float64(c1.Y - c2.Y), 2) +
		math.Pow(float64(c1.Z - c2.Z), 2),
	)
}
