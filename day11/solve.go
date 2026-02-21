package day11

import (
	"fmt"
	"strings"

	"github.com/nlduy0310/aoc-2025/day11/device"
)

func parseInput(input string) ([]device.Device, error) {
	lines := strings.Split(input, "\n")
	ret := make([]device.Device, 0, len(lines))
	for lineIdx, line := range lines {
		if len(line) == 0 {
			continue
		}
		if d, err := device.FromLine(line); err != nil {
			return nil, fmt.Errorf("unable to parse line %d: %w", lineIdx, err)
		} else {
			ret = append(ret, *d)
		}
	}
	return ret, nil
}

func countAllPaths(from, to string, devicesMap map[string]device.Device, memo map[string]int, exitDevice string) (int, error) {
	var countRecursively func(string) (int, error)
	countRecursively = func(initial string) (int, error) {
		if initial == to {
			return 1, nil
		} else if memoized, ok := memo[initial]; ok {
			return memoized, nil
		} else if initial == exitDevice {
			return 0, nil
		}

		ret := 0
		d, ok := devicesMap[initial]
		if !ok {
			return 0, fmt.Errorf("device %q not found", initial)
		}
		for _, output := range d.Outputs {
			c, err := countRecursively(output)
			if err != nil {
				return 0, err
			}
			ret += c
		}
		memo[initial] = ret
		return ret, nil
	}

	ret, err := countRecursively(from)
	if err != nil {
		return 0, err
	}
	return ret, nil
}

func makeDevicesMap(devices []device.Device) map[string]device.Device {
	ret := make(map[string]device.Device)
	for _, d := range devices {
		ret[d.Name] = d
	}
	return ret
}

func SolvePartOne(input string) (int, error) {
	devices, err := parseInput(input)
	if err != nil {
		return 0, fmt.Errorf("failed to parse input: %w", err)
	}

	countsMap := make(map[string]int)
	devicesMap := makeDevicesMap(devices)

	ret, err := countAllPaths("you", "out", devicesMap, countsMap, "out")
	if err != nil {
		return 0, err
	}
	return ret, nil
}

func SolvePartTwo(input string) (int, error) {
	devices, err := parseInput(input)
	if err != nil {
		return 0, fmt.Errorf("unable to parse input: %w", err)
	}

	// svr -> fft -> dac -> out
	// svr -> dac -> fft -> out
	devicesMap := makeDevicesMap(devices)
	fftCountsMap := make(map[string]int)
	dacCountsMap := make(map[string]int)
	outCountsMap := make(map[string]int)

	countPaths := func(from, to string, devicesMap map[string]device.Device, countsMap map[string]int) (int, error) {
		ret, err := countAllPaths(from, to, devicesMap, countsMap, "out")
		if err != nil {
			return 0, fmt.Errorf("unable to count paths from %q to %q: %w", from, to, err)
		}
		return ret, nil
	}

	svrToFft, err := countPaths("svr", "fft", devicesMap, fftCountsMap)
	if err != nil {
		return 0, err
	}
	fftToDac, err := countPaths("fft", "dac", devicesMap, dacCountsMap)
	if err != nil {
		return 0, err
	}
	dacToOut, err := countPaths("dac", "out", devicesMap, outCountsMap)
	if err != nil {
		return 0, err
	}
	svrToDac, err := countPaths("svr", "dac", devicesMap, dacCountsMap)
	if err != nil {
		return 0, err
	}
	dacToFft, err := countPaths("dac", "fft", devicesMap, fftCountsMap)
	if err != nil {
		return 0, err
	}
	fftToOut, err := countPaths("fft", "out", devicesMap, outCountsMap)
	if err != nil {
		return 0, err
	}

	return (svrToFft*fftToDac*dacToOut + svrToDac*dacToFft*fftToOut), nil
}
