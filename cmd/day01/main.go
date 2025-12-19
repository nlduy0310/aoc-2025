package main

import (
	"fmt"
	"log"

	"github.com/nlduy0310/aoc-2025/cli"
	"github.com/nlduy0310/aoc-2025/day01"
)

func main() {
	part, inp, err := cli.InitDefaults()
	if err != nil {
		log.Fatalf("failed to initialize arguments: %s", err)
	}

	switch *part {
	case cli.PartOne:
		{
			res, err := day01.SolvePartOne(*inp)
			if err != nil {
				log.Fatalf("error solving part one: %s", err)
			}
			fmt.Printf("Part one result: %d\n", res)
		}
	case cli.PartTwo:
		{
			panic("part two not implemented")
		}
	default:
		panic("impossible state")
	}
}
