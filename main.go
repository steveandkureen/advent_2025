package main

import (
	"flag"
	"fmt"
	"steveandkureen/advent_2025/internal"
	puzzels "steveandkureen/advent_2025/puzzles"
	"time"
)

func main() {
	day := flag.Int("day", 1, "which day to run")
	puzzle := flag.Int("puzzle", 0, "0=example, 1=easy, 2=hard")
	isHard := flag.Bool("hard", false, "true for hard")
	showOutput := flag.Bool("showOutput", true, "true for logging")

	flag.Parse()

	internal.LogOutput("Executing day: %d\n", *day)
	internal.LogOutput("Executing puzzle: %d\n", *puzzle)

	internal.ShowOutput = *showOutput
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
	internal.LogOutput("Loading data: %s\n", file)
	var result string
	var err error
	start := time.Now()
	switch *day {
	default:
		internal.LogOutput("Day not found: %d\n", *day)
	case 1:
		result, err = puzzels.FindCombonation(file, *isHard)
	case 2:
		result, err = puzzels.CheckProdcutIds(file, *isHard)
	case 3:
		result, err = puzzels.FindMaxVoltage(file, *isHard)
	case 4:
		result, err = puzzels.FindAdjecentRolls(file, *isHard)
	}
	internal.ShowOutput = true
	if err != nil {
		internal.LogOutput("Puzzel failed: %s\n", err)
	}
	internal.LogOutput("result: %s\n", result)
	elapsed := time.Since(start)
	internal.LogOutput("Execution time: %s\n", elapsed)
}
