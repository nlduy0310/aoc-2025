package bank

import (
	"fmt"

	"github.com/nlduy0310/aoc-2025/day03/battery"
)

type Bank struct {
	batteries []battery.Battery
}

func NewBank(batteries []battery.Battery) (*Bank, error) {
	ret := Bank{batteries}
	return &ret, nil
}

func (b Bank) Size() int {
	return len(b.batteries)
}

func (b Bank) BatteryAt(idx int) (battery.Battery, error) {
	if idx < 0 || idx >= b.Size() {
		bat, err := battery.NewBattery(0)
		if err != nil {
			panic("impossible initialization error")
		}
		return bat, fmt.Errorf("invalid index")
	}

	return b.batteries[idx], nil
}
