package vector

func ensureSameSize(v1, v2 Vector) error {
	if v1.Size() != v2.Size() {
		return sizeMismatchError
	}
	return nil
}
