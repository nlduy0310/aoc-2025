package mathutils

func AbsInt(val int) int {
	if val < 0 {
		return -val
	}
	return val
}
