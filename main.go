package main

import (
	"flag"
	"fmt"
	puzzels "steveandkureen/advent_2025/puzzles"
	"time"
)

func main() {
	day := flag.Int("day", 1, "which day to run")
	puzzle := flag.Int("puzzle", 0, "0=example, 1=easy, 2=hard")
	isHard := flag.Bool("hard", false, "true for hard")

	flag.Parse()

	fmt.Printf("Executing day: %d\n", *day)
	fmt.Printf("Executing puzzle: %d\n", *puzzle)

	var file string
	switch *puzzle {
	default:
		file = "example.txt"
	case 1:
		file = "easy.txt"
	case 2:
		file = "hard.txt"
	}
	file = fmt.Sprintf("./data/day%d/%s", *day, file)
	fmt.Printf("Loading data: %s\n", file)
	var result string
	var err error
	start := time.Now()
	switch *day {
	default:
		fmt.Printf("Day not found: %d\n", *day)
	case 1:
		result, err = puzzels.FindCombonation(file, *isHard)
	case 2:
		result, err = puzzels.CheckProdcutIds(file, *isHard)
	}
	if err != nil {
		fmt.Printf("Puzzel failed: %s\n", err)
	}
	fmt.Printf("result: %s\n", result)
	elapsed := time.Since(start)
	fmt.Printf("Execution time: %s\n", elapsed)
}
