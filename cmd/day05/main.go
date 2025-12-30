package main

import (
	"fmt"
	"log"

	"github.com/nlduy0310/aoc-2025/cli"
	"github.com/nlduy0310/aoc-2025/day05"
)

func main() {
	part, inp, err := cli.InitDefaults()
	if err != nil {
		log.Fatalf("failed to initialize arguments: %s", err)
	}

	switch *part {
		case cli.PartOne:
		{
			res, err := day05.SolvePartOne(*inp)
			if err != nil {
				log.Fatalf("failed to solve part one: %s", err)
			}
			fmt.Printf("Part one result: %d\n", res)
		}
		case cli.PartTwo:
		{
			res, err := day05.SolvePartTwo(*inp)
			if err != nil {
				log.Fatalf("failed to solve part two: %s", err)
			}
			fmt.Printf("Part two result: %d\n", res)
		}
		default:
			panic("impossible state")
	}
}
