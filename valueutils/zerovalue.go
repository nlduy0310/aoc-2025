package valueutils

func ZeroValue[T any]() T {
	var ret T
	return ret
}
