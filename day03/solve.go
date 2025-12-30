package day03

import (
	"fmt"
	"math"
	"strings"

	"github.com/nlduy0310/aoc-2025/datastructures/stack"
	"github.com/nlduy0310/aoc-2025/day03/bank"
	"github.com/nlduy0310/aoc-2025/day03/battery"
	"github.com/nlduy0310/aoc-2025/mathutils"
)

const PART_TWO_BATTERIES_COUNT int = 12

// getMaxBattery returns the index of the battery with the highest joltage
// in the half-open interval [fromIdx, toIdx].
// It returns an error if the range is invalid or falls outside the bank size.
func getMaxBattery(b bank.Bank, fromIdx, toIdx int) (int, error) {
	if b.Size() == 0 {
		return -1, fmt.Errorf("empty bank")
	}
	if !(fromIdx < toIdx) || !(fromIdx >= 0 && toIdx <= b.Size()) {
		return -1, fmt.Errorf("invalid search range")
	}

	maxIdx := -1
	maxJoltage := math.MinInt
	for idx := fromIdx; idx < toIdx; idx++ {
		bat, err := b.BatteryAt(idx)
		if err != nil {
			panic("impossible index error")
		}
		if bat.Joltage > maxJoltage {
			maxJoltage = bat.Joltage
			maxIdx = idx
		}
	}

	return maxIdx, nil
}

func parseInp(inp string) ([]bank.Bank, error) {
	lines := strings.Split(inp, "\n")
	banks := make([]bank.Bank, 0, len(lines))
	for lineIdx, line := range lines {
		line = strings.TrimSuffix(line, "\n")
		if len(line) == 0 {
			continue
		}

		bank, err := bank.ParseFromString(line)
		if err != nil {
			return nil, fmt.Errorf("failed to parse bank from line %d: %w", lineIdx+1, err)
		}
		banks = append(banks, *bank)
	}

	return banks, nil
}

func SolvePartOne(inp string) (int, error) {
	banks, err := parseInp(inp)
	if err != nil {
		return 0, err
	}

	sum := 0
	for bankIdx, bank := range banks {
		firstBatIdx, err := getMaxBattery(bank, 0, bank.Size()-1)
		if err != nil {
			return 0, fmt.Errorf("failed to get max battery of bank %d: %w", bankIdx+1, err)
		}

		secondBatIdx, err := getMaxBattery(bank, firstBatIdx+1, bank.Size())
		if err != nil {
			return 0, fmt.Errorf("failed to get max battery of bank %d from %d to %d: %w", bankIdx+1, firstBatIdx+1, bank.Size(), err)
		}

		firstBat, err := bank.BatteryAt(firstBatIdx)
		if err != nil {
			panic("impossible index error")
		}
		secondBat, err := bank.BatteryAt(secondBatIdx)
		if err != nil {
			panic("impossible indedx error")
		}

		sum += firstBat.Joltage*int(math.Pow10(int(mathutils.CountDigits(secondBat.Joltage)))) + secondBat.Joltage
	}

	return sum, nil
}

func SolvePartTwo(inp string) (int, error) {
	banks, err := parseInp(inp)
	if err != nil {
		return 0, err
	}

	sum := 0
	for bankIdx, bank := range banks {
		if bank.Size() < PART_TWO_BATTERIES_COUNT {
			return 0, fmt.Errorf("expected bank %d to have at least %d batteries, got %d", bankIdx+1, PART_TWO_BATTERIES_COUNT, bank.Size())
		}

		removalCount := bank.Size() - PART_TWO_BATTERIES_COUNT
		batteriesStack := stack.NewEmptyLinkedListStack[battery.Battery]()
		for batIdx := 0; batIdx < bank.Size(); batIdx++ {
			currentBat, _ := bank.BatteryAt(batIdx)
			for lastBat, err := batteriesStack.Peek(); removalCount > 0 && err == nil && lastBat.Joltage < currentBat.Joltage; lastBat, err = batteriesStack.Peek() {
				batteriesStack.Pop()
				removalCount--
			}
			batteriesStack.Push(currentBat)
		}
		if batteriesStack.Size() < PART_TWO_BATTERIES_COUNT {
			panic("impossible state")
		}


		num := 0
		stackSize := batteriesStack.Size()
		for i := stackSize - 1; i >= stackSize - PART_TWO_BATTERIES_COUNT; i-- {
			bat, _ := batteriesStack.GetAt(i)
			num = num * 10 + bat.Joltage
		}
		sum += num
	}

	return sum, nil
}
