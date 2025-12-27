package bank

import (
	"fmt"

	"github.com/nlduy0310/aoc-2025/day03/battery"
)

// ParseFromString returns the corresponding power bank,
// treating each digit in the string as a battery joltage value
func ParseFromString(str string) (*Bank, error) {
	batteries := make([]battery.Battery, 0)

	for _, r := range str {
		if !(r >= '0' && r <= '9') {
			return nil, fmt.Errorf("invalid character: %q", r)
		}
		joltage := int(r - '0')

		battery, err := battery.NewBattery(joltage)
		if err != nil {
			panic("impossible battery initialization error")
		}
		batteries = append(batteries, battery)
	}

	ret, err := NewBank(batteries)
	if err != nil {
		return nil, fmt.Errorf("error initializing bank: %w", err)
	}

	return ret, nil
}
