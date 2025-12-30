package day06

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/nlduy0310/aoc-2025/day06/op"
	"github.com/nlduy0310/aoc-2025/day06/problem"
	"github.com/nlduy0310/aoc-2025/valueutils"
)

func splitNonEmpty(str, sep string) []string {
	tokens := strings.Split(str, sep)
	ret := make([]string, 0)
	for _, token := range tokens {
		if len(token) > 0 {
			ret = append(ret, token)
		}
	}
	return ret
}

func parseIntsRow(str, sep string) ([]int, error) {
	tokens := splitNonEmpty(str, sep)
	ret := make([]int, len(tokens))
	for tokenIdx, token := range tokens {
		val, err := strconv.Atoi(token)
		if err != nil {
			return nil, fmt.Errorf("%q is not a valid int", token)
		}

		ret[tokenIdx] = val
	}
	return ret, nil
}

func parsePartOne(inp string) ([]problem.Problem, error) {
	lines := strings.Split(inp, "\n")
	if len(lines) > 0 && len(lines[len(lines)-1]) == 0 {
		lines = lines[:len(lines)-1]
	}
	if len(lines) < 2 {
		return nil, fmt.Errorf("expected at least 2 non-empty lines, got %d", len(lines))
	}

	ints, err := parseIntsRow(lines[0], " ")
	if err != nil {
		return nil, fmt.Errorf("failed to parse first line: %w", err)
	}
	numProblems := len(ints)
	numArgs := len(lines) - 1

	argsMat := make([][]int, numProblems)
	for p := range numProblems {
		argsMat[p] = make([]int, numArgs)
	}
	for argIdx := range numArgs {
		line := lines[argIdx]
		ints, err := parseIntsRow(line, " ")
		if err != nil {
			return nil, fmt.Errorf("failed to parse line %d: %w", argIdx+1, err)
		}
		if len(ints) != numProblems {
			return nil, fmt.Errorf("number of arguments on line %d (%d) doesn't match the expected number of problems (%d)", argIdx+1, len(ints), numProblems)
		}
		for problemIndex, arg := range ints {
			argsMat[problemIndex][argIdx] = arg
		}
	}

	ops := make([]op.Op, numProblems)
	opStrings := splitNonEmpty(lines[len(lines)-1], " ")
	if len(opStrings) != numProblems {
		return nil, fmt.Errorf("expected %d operators on line %d, got %d", numProblems, len(lines), len(opStrings))
	}
	for idx, opString := range opStrings {
		o, err := op.FromString(opString)
		if err != nil {
			return nil, fmt.Errorf("failed to parse op string %q on line %d: %w", opString, len(lines), err)
		}
		ops[idx] = o
	}

	problems := make([]problem.Problem, 0, numProblems)
	for problemIdx := range numProblems {
		operator := ops[problemIdx]
		args := make([]int, numArgs)
		for argIdx := range numArgs {
			args[argIdx] = argsMat[problemIdx][argIdx]
		}
		problem, err := problem.New(args, operator)
		if err != nil {
			return nil, fmt.Errorf("failed to initialize problem %d: %w", problemIdx+1, err)
		}
		problems = append(problems, *problem)
	}

	return problems, nil
}

func splitProblemsByColums(lines []string) ([]struct{from, to int}) {
	cols := len(lines[0])
	rows := len(lines)
	emptyCols := make([]int, 0)
	for col := range cols {
		empty := true
		for row := range rows {
			if lines[row][col] != ' ' {
				empty = false
				break
			}
		}
		if empty {
			emptyCols = append(emptyCols, col)
		}
	}

	ret := make([]struct{from, to int}, len(emptyCols)+1)
	for idx, emptyCol := range emptyCols {
		var from int
		if idx == 0 {
			from = 0
		} else {
			from = emptyCols[idx-1] + 1
		}
		ret[idx] = struct{from, to int}{from, emptyCol}
	}
	if len(emptyCols) > 0 {
		ret[len(emptyCols)] = struct{from, to int}{emptyCols[len(emptyCols)-1]+1, cols}
	}

	return ret
}

func parseProblemByColRange(lines []string, colRange struct{from, to int}) (problem.Problem, error) {
	rows := len(lines)
	args := make([]int, colRange.to-colRange.from)
	for col := colRange.to-1; col >= colRange.from; col-- {
		arg := 0

		for row := range rows-1 {
			ch := lines[row][col]
			if ch == ' ' {
				continue
			} else if ch >= '0' && ch <= '9' {
				arg = arg * 10 + int(ch - '0')
			} else {
				return valueutils.ZeroValue[problem.Problem](), fmt.Errorf("unknown character %q at row %d, col %d", ch, row, col)
			}
		}

		args[col-colRange.from] = arg
	}

	opString := string(lines[rows-1][colRange.from])
	operator, err := op.FromString(opString)
	if err != nil {
		return valueutils.ZeroValue[problem.Problem](), fmt.Errorf("unknown operator string %q at row %d, col %d", opString, rows-1, colRange.from)
	}

	prob, err := problem.New(args, operator)
	if err != nil {
		return valueutils.ZeroValue[problem.Problem](), fmt.Errorf("failed to initialize problem: %w", err)
	}

	return *prob, nil
}

func parsePartTwo(inp string) ([]problem.Problem, error) {
	lines := splitNonEmpty(inp, "\n")
	if len(lines) < 2 {
		return nil, fmt.Errorf("expected at least 2 non-empty lines, got %d", len(lines))
	}

	maxLineLength := math.MinInt
	for _, line := range lines {
		if len(line) > maxLineLength {
			maxLineLength = len(line)
		}
	}
	for i := range lines {
		if len(lines[i]) < maxLineLength {
			diff := maxLineLength - len(lines[i])
			lines[i] += strings.Repeat(" ", diff)
		}
	}

	cranges := splitProblemsByColums(lines)
	problems := make([]problem.Problem, 0, len(cranges))
	for _, crange := range cranges {
		problem, err := parseProblemByColRange(lines, crange)
		if err != nil {
			return nil, fmt.Errorf("failed to parse problem from col %d to %d: %w", crange.from, crange.to, err)
		}
		problems = append(problems, problem)
	}

	return problems, nil
}

func SolvePartOne(inp string) (int, error) {
	problems, err := parsePartOne(inp)
	if err != nil {
		return 0, fmt.Errorf("failed to parse input file: %w", err)
	}

	total := 0
	for _, problem := range problems {
		total += problem.Evaluate()
	}
	return total, nil
}

func SolvePartTwo(inp string) (int, error) {
	problems, err := parsePartTwo(inp)
	if err != nil {
		return 0, fmt.Errorf("failed to parse input file: %w", err)
	}

	total := 0
	for _, problem := range problems {
		total += problem.Evaluate()
	}
	return total, nil
}
