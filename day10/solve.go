package day10

import (
	"encoding/binary"
	"errors"
	"fmt"
	"hash/maphash"
	"math"
	"math/bits"
	"slices"
	"strings"
	"sync"

	"github.com/nlduy0310/aoc-2025/datastructures/queue"
	"github.com/nlduy0310/aoc-2025/day10/button"
	"github.com/nlduy0310/aoc-2025/day10/cache"
	"github.com/nlduy0310/aoc-2025/day10/lightdiagram"
	"github.com/nlduy0310/aoc-2025/day10/vector"
	"github.com/nlduy0310/aoc-2025/sliceutils"
)

type inputLine struct {
	targetLightsState   lightdiagram.LightDiagram
	buttons             []button.Button
	targetCountersState vector.Vector
}

type partoneSearchState struct {
	ld               *lightdiagram.LightDiagram
	pressedButtonIds *[]int
}

var hashSeed maphash.Seed = maphash.MakeSeed()
var partTwoCache cache.Cache[int] = cache.New[int]()

func (s partoneSearchState) Pressed(buttonId int, b button.Button) (*partoneSearchState, error) {
	ld := s.ld.Clone()
	err := b.Flip(*ld)
	if err != nil {
		return nil, fmt.Errorf("can not press button: %w", err)
	}

	pressedButtonIds := make([]int, len(*s.pressedButtonIds), len(*s.pressedButtonIds)+1)
	copy(pressedButtonIds, *s.pressedButtonIds)
	pressedButtonIds = append(pressedButtonIds, buttonId)
	ret := partoneSearchState{ld: ld, pressedButtonIds: &pressedButtonIds}
	return &ret, nil
}

func parseInput(input string) ([]inputLine, error) {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")

	ret := make([]inputLine, 0, len(lines))
	for lineIdx, line := range lines {
		tokens := strings.Split(line, " ")
		if len(tokens) < 3 {
			return nil, fmt.Errorf("expect line %d to have at least 3 tokens, got %d", lineIdx, len(tokens))
		}

		ldToken, buttonTokens, countersToken := tokens[0], tokens[1:len(tokens)-1], tokens[len(tokens)-1]
		ld, err := lightdiagram.FromString(ldToken)
		if err != nil {
			return nil, fmt.Errorf("failed to parse light diagram from %q on line %d: %w", ldToken, lineIdx, err)
		}

		buttons := make([]button.Button, 0, len(buttonTokens))
		for _, buttonToken := range buttonTokens {
			b, err := button.FromString(buttonToken)
			if err != nil {
				return nil, fmt.Errorf("failed to parse button from %q on line %d: %w", buttonToken, lineIdx, err)
			}
			buttons = append(buttons, *b)
		}

		trimmedCountersToken := strings.TrimPrefix(countersToken, "{")
		trimmedCountersToken = strings.TrimSuffix(trimmedCountersToken, "}")
		countersVec, err := vector.FromCSVLine(trimmedCountersToken)
		if err != nil {
			return nil, fmt.Errorf("failed to parse counters from %q: %w", countersToken, err)
		}

		lineData := inputLine{*ld, buttons, *countersVec}
		ret = append(ret, lineData)
	}

	return ret, nil
}

func optimalPresses(targetState lightdiagram.LightDiagram, buttons []button.Button) (int, error) {
	bfsQueue := queue.EmptyLQueue[partoneSearchState]()
	initialLd, _ := lightdiagram.New(targetState.Size())
	bfsQueue.Enqueue(partoneSearchState{ld: initialLd, pressedButtonIds: &[]int{}})

	for {
		curState, err := bfsQueue.Dequeue()
		if err != nil {
			break
		}

		if lightdiagram.Equal(*curState.ld, targetState) {
			return len(*curState.pressedButtonIds), nil
		}

		for buttonId, button := range buttons {
			if !slices.Contains(*curState.pressedButtonIds, buttonId) {
				newState, err := curState.Pressed(buttonId, button)
				if err != nil {
					return 0, err
				}
				bfsQueue.Enqueue(*newState)
			}
		}

	}

	return -1, fmt.Errorf("unable to find buttons combination")
}

func SolvePartOne(input string) (int, error) {
	linesData, err := parseInput(input)
	if err != nil {
		return 0, fmt.Errorf("can not parse input: %w", err)
	}

	type result struct {
		data int
		err  error
	}
	var wg sync.WaitGroup
	resChan := make(chan result)
	for lineIdx, lineData := range linesData {
		wg.Add(1)
		go func(lineIdx int, data inputLine) {
			defer wg.Done()
			presses, err := optimalPresses(data.targetLightsState, data.buttons)
			resChan <- result{data: presses, err: err}
		}(lineIdx, lineData)
	}

	go func() {
		wg.Wait()
		close(resChan)
	}()

	ret := 0
	errs := make([]error, 0)
	for res := range resChan {
		if res.err != nil {
			errs = append(errs, err)
		} else {
			ret += res.data
		}
	}

	if len(errs) > 0 {
		return 0, errors.Join(errs...)
	}
	return ret, nil
}

func parityVector(source vector.Vector) vector.Vector {
	elements := make([]int, source.Size())
	for idx := range source.Size() {
		if val, _ := source.At(idx); val%2 != 0 {
			elements[idx] = 1
		}
	}
	return vector.New(elements)
}

func isZero(v vector.Vector) bool {
	return v.IsAll(func(val int) bool {
		return val == 0
	})
}

func xor(v1, v2 vector.Vector) vector.Vector {
	ret, _ := vector.ExecuteElementWise(v1, v2, func(a, b int) int {
		return a ^ b
	})
	return *ret
}

func halve(v vector.Vector) vector.Vector {
	return v.Map(func(a int) int {
		return a / 2
	})
}

func addMany(vecs []vector.Vector) vector.Vector {
	v0 := vecs[0].Clone()
	ret := &v0
	for idx := 1; idx < len(vecs); idx++ {
		ret, _ = vector.Add(*ret, vecs[idx])
	}
	return *ret
}

func isNegative(v vector.Vector) bool {
	return !v.IsAll(func(a int) bool {
		return a >= 0
	})
}

func findAllCombs(buttons []vector.Vector, parityMask vector.Vector) [][]vector.Vector {
	if len(buttons) > 64 {
		panic("function expects at most 64 buttons")
	}

	type searchState struct {
		start             int
		current           vector.Vector
		buttonIndexesMask uint64
	}

	translateIndexesMask := func(mask uint64) []vector.Vector {
		ret := make([]vector.Vector, 0, bits.OnesCount64(mask))
		for idx := range buttons {
			if bit := mask & (uint64(1) << idx); bit != 0 {
				ret = append(ret, buttons[idx])
			}
		}
		return ret
	}

	ret := make([][]vector.Vector, 0)
	q := queue.EmptyLQueue[searchState]()

	initialState := searchState{
		start:             0,
		current:           vector.Fill(parityMask.Size(), 0),
		buttonIndexesMask: 0,
	}
	q.Enqueue(initialState)
	for {
		curState, err := q.Dequeue()
		if err != nil {
			break
		}

		if equal, _ := vector.Equal(curState.current, parityMask); equal {
			ret = append(ret, translateIndexesMask(curState.buttonIndexesMask))
		}

		for idx := curState.start; idx < len(buttons); idx++ {
			q.Enqueue(searchState{
				start:             idx + 1,
				current:           xor(curState.current, buttons[idx]),
				buttonIndexesMask: curState.buttonIndexesMask | (uint64(1) << idx),
			})
		}
	}

	return ret
}

func partTwoCacheKey(buttons []vector.Vector, target vector.Vector) uint64 {
	var buf [8]byte
	var hash maphash.Hash
	hash.SetSeed(hashSeed)

	writeHash := func(val uint64) {
		binary.LittleEndian.PutUint64(buf[:], val)
		hash.Write(buf[:])
	}

	writeHash(uint64(len(buttons)))
	for _, button := range buttons {
		writeHash(button.Hash())
	}
	writeHash(target.Hash())

	return hash.Sum64()
}

// solution by u/tenthmascot
func solveLinePartTwo(buttons []vector.Vector, target vector.Vector) (int, error) {
	if isZero(target) {
		return 0, nil
	}

	key := partTwoCacheKey(buttons, target)
	if cachedVal, ok := partTwoCache.Get(key); ok {
		return cachedVal, nil
	}

	ret := math.MaxInt
	found := false
	parityMask := parityVector(target)
	parityCombs := findAllCombs(buttons, parityMask)
	if len(parityCombs) == 0 {
		return 0, fmt.Errorf("can not find any combs for parity mask %s", parityMask.String())
	}
	for _, parityComb := range parityCombs {
		var remainder *vector.Vector
		if len(parityComb) == 0 {
			tmp := target.Clone()
			remainder = &tmp
		} else {
			remainder, _ = vector.Sub(target, addMany(parityComb))
		}
		if isNegative(*remainder) {
			continue
		}
		halfRemainder := halve(*remainder)
		n, err := solveLinePartTwo(buttons, halfRemainder)
		if err != nil {
			continue
		}
		if presses := len(parityComb) + 2*n; presses < ret {
			ret = presses
			found = true
		}
	}

	if !found {
		return 0, fmt.Errorf("no solutions found")
	}
	partTwoCache.Set(key, ret)
	return ret, nil
}

func SolvePartTwo(input string) (int, error) {
	linesData, err := parseInput(input)
	if err != nil {
		return 0, fmt.Errorf("failed to parse input: %w", err)
	}

	type lineResult struct {
		lineIdx int
		pushes  int
		err     error
	}
	var wg sync.WaitGroup
	resChan := make(chan lineResult)
	for lineIdx, lineData := range linesData {
		buttonVecs := make([]vector.Vector, len(lineData.buttons))
		for buttonIdx, button := range lineData.buttons {
			buttonVecs[buttonIdx] = button.Vector(lineData.targetCountersState.Size())
		}
		targetCounters := lineData.targetCountersState
		lineIdx_ := lineIdx

		wg.Go(func() {
			pushes, err := solveLinePartTwo(buttonVecs, targetCounters)
			resChan <- lineResult{lineIdx_, pushes, err}
		})
	}

	go func() {
		wg.Wait()
		close(resChan)
	}()

	type lineError struct {
		lineIdx int
		err     error
	}
	errs := make([]lineError, 0)
	ret := 0
	for result := range resChan {
		if result.err != nil {
			errs = append(errs, lineError{result.lineIdx, result.err})
		} else {
			ret += result.pushes
		}
	}

	if len(errs) > 0 {
		return 0, errors.Join(
			sliceutils.Map(errs, func(le lineError) error {
				return fmt.Errorf("unable to solve line %d: %w", le.lineIdx, le.err)
			})...,
		)
	}

	return ret, nil
}
