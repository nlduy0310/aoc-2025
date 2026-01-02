package day08

import (
	"cmp"
	"fmt"
	"slices"

	"github.com/nlduy0310/aoc-2025/day08/circuit"
	"github.com/nlduy0310/aoc-2025/day08/coordinates"
	"github.com/nlduy0310/aoc-2025/stringutils"
)

type coordsPair struct {
	first  coordinates.Coordinates
	second coordinates.Coordinates
}

const NUM_PAIRS int = 1000
const NUM_MUL int = 3

func parseInput(input string) ([]coordinates.Coordinates, error) {
	lines := stringutils.SplitNonEmpty(input, "\n")
	ret := make([]coordinates.Coordinates, 0, len(lines))
	for lineIdx, line := range lines {
		coords, err := coordinates.FromString(line)
		if err != nil {
			return nil, fmt.Errorf("failed to parse coordinates on line %d: %w", lineIdx+1, err)
		}
		ret = append(ret, coords)
	}
	return ret, nil
}

func numPairs(n int) int {
	return n * (n - 1) / 2
}

func findKClosestPair(list []coordinates.Coordinates, k int) ([]coordsPair, error) {
	nPairs := numPairs(len(list))
	if k <= 0 {
		return nil, fmt.Errorf("k must be positive")
	} else if k > nPairs {
		return nil, fmt.Errorf("maximum number of pairs is %d", nPairs)
	}

	type pairDistance struct {
		firstIdx  int
		secondIdx int
		dist      float64
	}
	pairDistances := make([]pairDistance, 0, nPairs)
	for i := 0; i < len(list)-1; i++ {
		for j := i + 1; j < len(list); j++ {
			pairDistances = append(pairDistances, pairDistance{
				firstIdx:  i,
				secondIdx: j,
				dist:      coordinates.EuclideanDistance(list[i], list[j]),
			})
		}
	}

	slices.SortFunc(pairDistances, func(a, b pairDistance) int {
		return cmp.Compare(a.dist, b.dist)
	})
	ret := make([]coordsPair, 0, k)
	for i := range k {
		pairInfo := pairDistances[i]
		ret = append(ret, coordsPair{
			first:  list[pairInfo.firstIdx],
			second: list[pairInfo.secondIdx],
		})
	}
	return ret, nil
}

func findCircuitIdx(circuitsList []circuit.Circuit, box coordinates.Coordinates) int {
	for idx, circ := range circuitsList {
		if circ.Has(box) {
			return idx
		}
	}

	return -1
}

func SolvePartOne(inp string) (int, error) {
	coordsList, err := parseInput(inp)
	if err != nil {
		return 0, fmt.Errorf("failed to parse input file: %w", err)
	}

	closestPairs, err := findKClosestPair(coordsList, NUM_PAIRS)
	if err != nil {
		return 0, fmt.Errorf("failed to find %d closest pairs: %w", NUM_PAIRS, err)
	}

	circuitsList := make([]circuit.Circuit, 0)
	for _, pair := range closestPairs {
		firstCircuitIdx := findCircuitIdx(circuitsList, pair.first)
		secondCircuitIdx := findCircuitIdx(circuitsList, pair.second)
		if firstCircuitIdx < 0 && secondCircuitIdx < 0 {
			circuitsList = append(circuitsList, circuit.FromBoxes(pair.first, pair.second))
		} else if firstCircuitIdx < 0 && secondCircuitIdx >= 0 {
			circuitsList[secondCircuitIdx].Add(pair.first)
		} else if firstCircuitIdx >= 0 && secondCircuitIdx < 0 {
			circuitsList[firstCircuitIdx].Add(pair.second)
		} else if firstCircuitIdx != secondCircuitIdx {
			newCircuit := circuit.Join(circuitsList[firstCircuitIdx], circuitsList[secondCircuitIdx])
			leftIdx, rightIdx := min(firstCircuitIdx, secondCircuitIdx), max(firstCircuitIdx, secondCircuitIdx)
			circuitsList = append(circuitsList[:rightIdx], circuitsList[rightIdx+1:]...)
			circuitsList = append(circuitsList[:leftIdx], circuitsList[leftIdx+1:]...)
			circuitsList = append(circuitsList, newCircuit)
		}
	}

	slices.SortFunc(circuitsList, func(c1, c2 circuit.Circuit) int {
		return cmp.Compare(c2.Size(), c1.Size())
	})
	res := 1
	for i := range NUM_MUL {
		res *= circuitsList[i].Size()
	}
	return res, nil
}

func SolvePartTwo(inp string) (int, error) {
	coordsList, err := parseInput(inp)
	if err != nil {
		return 0, fmt.Errorf("failed to parse input file: %w", err)
	}

	nPairs := numPairs(len(coordsList))
	closestPairs, err := findKClosestPair(coordsList, nPairs)
	if err != nil {
		return 0, fmt.Errorf("failed to find %d closest pair: %w", nPairs, err)
	}

	var lastPair coordsPair
	circuitsList := make([]circuit.Circuit, 0)
	for _, pair := range closestPairs {
		firstCircuitIdx := findCircuitIdx(circuitsList, pair.first)
		secondCircuitIdx := findCircuitIdx(circuitsList, pair.second)
		connected := true
		if firstCircuitIdx < 0 && secondCircuitIdx < 0 {
			circuitsList = append(circuitsList, circuit.FromBoxes(pair.first, pair.second))
		} else if firstCircuitIdx < 0 && secondCircuitIdx >= 0 {
			circuitsList[secondCircuitIdx].Add(pair.first)
		} else if firstCircuitIdx >= 0 && secondCircuitIdx < 0 {
			circuitsList[firstCircuitIdx].Add(pair.second)
		} else if firstCircuitIdx != secondCircuitIdx {
			newCircuit := circuit.Join(circuitsList[firstCircuitIdx], circuitsList[secondCircuitIdx])
			leftIdx, rightIdx := min(firstCircuitIdx, secondCircuitIdx), max(firstCircuitIdx, secondCircuitIdx)
			circuitsList = append(circuitsList[:rightIdx], circuitsList[rightIdx+1:]...)
			circuitsList = append(circuitsList[:leftIdx], circuitsList[leftIdx+1:]...)
			circuitsList = append(circuitsList, newCircuit)
		} else {
			connected = false
		}

		if connected {
			lastPair = pair
			// how the fuck does this work?????
			// if count >= NUM_PAIRS && len(circuitsList) == 1 {
			// 	break
			// }
		}
	}

	return lastPair.first.X * lastPair.second.X, nil
}
