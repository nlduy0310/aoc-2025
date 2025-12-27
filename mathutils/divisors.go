package mathutils

// Divisors returns all positive divisors of an int.
//
// This function assumes the input is non-zero.
func Divisors(i int) []int {
	if i == 0 {
		return []int{}
	} else if i < 0 {
		i = -i
	}

	small := make([]int, 0)
	large := make([]int, 0)
	for d := 1; d*d <= i; d++ {
		if i%d == 0 {
			small = append(small, d)
			if d*d != i {
				large = append(large, i/d)
			}
		}
	}

	ret := make([]int, 0, len(small)+len(large))
	for _, d := range small {
		ret = append(ret, d)
	}
	for idx := len(large) - 1; idx >= 0; idx-- {
		ret = append(ret, large[idx])
	}

	return ret
}
