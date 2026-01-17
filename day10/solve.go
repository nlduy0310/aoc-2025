package day10

import (
	"fmt"
	"slices"
	"strings"

	"github.com/nlduy0310/aoc-2025/datastructures/queue"
	"github.com/nlduy0310/aoc-2025/day10/button"
	"github.com/nlduy0310/aoc-2025/day10/lightdiagram"
)

type inputLine struct {
	targetLightsState lightdiagram.LightDiagram
	buttons           []button.Button
}

type state struct {
	ld               *lightdiagram.LightDiagram
	pressedButtonIds *[]int
}

func (s state) Pressed(buttonId int, b button.Button) (*state, error) {
	ld := s.ld.Clone()
	err := b.Flip(*ld)
	if err != nil {
		return nil, fmt.Errorf("can not press button: %w", err)
	}

	pressedButtonIds := make([]int, len(*s.pressedButtonIds), len(*s.pressedButtonIds)+1)
	copy(pressedButtonIds, *s.pressedButtonIds)
	pressedButtonIds = append(pressedButtonIds, buttonId)
	ret := state{ld: ld, pressedButtonIds: &pressedButtonIds}
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

		ldToken, buttonTokens, _ := tokens[0], tokens[1:len(tokens)-1], tokens[len(tokens)-1]
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
		lineData := inputLine{targetLightsState: *ld, buttons: buttons}
		ret = append(ret, lineData)
	}

	return ret, nil
}

func optimalPresses(targetState lightdiagram.LightDiagram, buttons []button.Button) (int, error) {
	bfsQueue := queue.EmptyLQueue[state]()
	initialLd, _ := lightdiagram.New(targetState.Size())
	bfsQueue.Enqueue(state{ld: initialLd, pressedButtonIds: &[]int{}})

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

	ret := 0
	for lineIdx, lineData := range linesData {
		presses, err := optimalPresses(lineData.targetLightsState, lineData.buttons)
		if err != nil {
			return 0, fmt.Errorf("failed to solve line %d: %w", lineIdx+1, err)
		}
		ret += presses
	}

	return ret, nil
}

func SolvePartTwo(input string) (int, error) {
	panic("not implemented")
}
