package main

import (
	"fmt"
	"log"

	"github.com/nlduy0310/aoc-2025/cli"
	"github.com/nlduy0310/aoc-2025/day11"
)

func main() {
	part, inp, err := cli.InitDefaults()
	if err != nil {
		log.Fatalf("failed to initialize arguments: %s", err.Error())
	}

	switch *part {
	case cli.PartOne:
		res, err := day11.SolvePartOne(*inp)
		if err != nil {
			log.Fatalf("can not solve part one: %s", err.Error())
		}
		fmt.Printf("Part one result: %d\n", res)
	case cli.PartTwo:
		res, err := day11.SolvePartTwo(*inp)
		if err != nil {
			log.Fatalf("can not solve part two: %s", err.Error())
		}
		fmt.Printf("Part two result: %d\n", res)
	default:
		panic("impossible part option")
	}
}
