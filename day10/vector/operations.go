package vector

import (
	"fmt"
	"slices"
)

var elementAdd = func(a, b int) int {
	return a + b
}

var elementSub = func(a, b int) int {
	return a - b
}

var elementDiv = func(a, b int) int {
	return a / b
}

func executeElementWise(v1, v2 Vector, op func(int, int) int) Vector {
	elements := make([]int, v1.Size())
	for idx := range v1.Size() {
		elements[idx] = op(v1.elements[idx], v2.elements[idx])
	}
	return New(elements)
}

func ExecuteElementWise(v1, v2 Vector, op func(int, int) int) (*Vector, error) {
	if err := ensureSameSize(v1, v2); err != nil {
		return nil, err
	}
	ret := executeElementWise(v1, v2, op)
	return &ret, nil
}

func Add(v1, v2 Vector) (*Vector, error) {
	if err := ensureSameSize(v1, v2); err != nil {
		return nil, err
	}

	ret := executeElementWise(v1, v2, elementAdd)
	return &ret, nil
}

func Sub(v1, v2 Vector) (*Vector, error) {
	if err := ensureSameSize(v1, v2); err != nil {
		return nil, err
	}

	ret := executeElementWise(v1, v2, elementSub)
	return &ret, nil
}

func Div(v1, v2 Vector) (*Vector, error) {
	if err := ensureSameSize(v1, v2); err != nil {
		return nil, err
	}

	if slices.Contains(v2.elements, 0) {
		return nil, fmt.Errorf("zero division")
	}

	ret := executeElementWise(v1, v2, elementDiv)
	return &ret, nil
}

func Equal(v1, v2 Vector) (bool, error) {
	if err := ensureSameSize(v1, v2); err != nil {
		return false, err
	}

	for idx := range v1.Size() {
		if v1.elements[idx] != v2.elements[idx] {
			return false, nil
		}
	}
	return true, nil
}

func (v Vector) Map(mapper func(int) int) Vector {
	ret := Fill(v.Size(), 0)
	for idx := range v.Size() {
		ret.elements[idx] = mapper(v.elements[idx])
	}
	return ret
}
