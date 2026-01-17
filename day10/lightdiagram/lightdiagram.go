package lightdiagram

import "fmt"

type LightDiagram struct {
	lights []bool
}

func New(size int) (*LightDiagram, error) {
	if size <= 0 {
		return nil, fmt.Errorf("non-positive diagram size not allowed")
	}

	lights_ := make([]bool, size)
	ret := LightDiagram{lights: lights_}
	return &ret, nil
}

func (ld LightDiagram) Size() int {
	return len(ld.lights)
}

func Equal(ld1, ld2 LightDiagram) bool {
	if ld1.Size() != ld2.Size() {
		return false
	}

	for idx := range ld1.Size() {
		if ld1.lights[idx] != ld2.lights[idx] {
			return false
		}
	}

	return true
}

func (ld LightDiagram) Clone() *LightDiagram {
	lights_ := make([]bool, len(ld.lights))
	copy(lights_, ld.lights)
	ret := LightDiagram{lights: lights_}
	return &ret
}

func (ld *LightDiagram) Toggle(lightIdx int) error {
	if lightIdx < 0 || lightIdx >= ld.Size() {
		return fmt.Errorf("diagram does not contain light %d", lightIdx)
	}

	ld.lights[lightIdx] = !ld.lights[lightIdx]
	return nil
}

func (ld LightDiagram) Toggled(lightIdx int) (*LightDiagram, error) {
	clone := ld.Clone()
	err := clone.Toggle(lightIdx)
	return clone, err
}
