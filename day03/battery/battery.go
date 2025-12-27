package battery

import "fmt"

type Battery struct {
	Joltage int
}

func NewBattery(joltage int) (Battery, error) {
	if !(joltage >= 0 && joltage <= 9) {
		return Battery{}, fmt.Errorf("joltage out of [0, 9] range: %d", joltage)
	}

	return Battery{joltage}, nil
}
