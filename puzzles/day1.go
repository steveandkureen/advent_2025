package puzzles

import (
	"fmt"
	"steveandkureen/advent_2025/internal"
	"strconv"
	"strings"
)

func FindCombonation(filename string, isHard bool) (string, error) {
	data, err := internal.ReadFile(filename)
	if err != nil {
		internal.LogOutput("error reading file: %s\n", err)
	}
	internal.LogOutput("Input: %s\n", string(data))
	currentPos := 50
	steps := strings.Split(data, "\n")
	result := 0
	for _, line := range steps {
		if len(line) == 0 {
			continue
		}
		direction := line[0]
		num, err := strconv.Atoi(line[1:])
		if err != nil {
			internal.LogOutput("Bad data found %s\n", line)
		}

		if isHard {
			result, currentPos = doHardCombo(direction, num, currentPos, result)
		} else {
			result, currentPos = doEasyCombo(direction, num, currentPos, result)
		}
	}
	internal.LogOutput("Answer: %d\n", result)
	return fmt.Sprintf("%d", result), nil
}

func doEasyCombo(direction byte, num int, currentPos int, result int) (int, int) {
	if direction == 'L' {
		currentPos -= num
	} else {
		currentPos += num
	}

	for currentPos < 0 {
		currentPos += 100
	}
	for currentPos > 99 {
		currentPos -= 100
	}

	if currentPos == 0 {
		result++
	}
	return result, currentPos
}

func doHardCombo(direction byte, num int, currentPos int, result int) (int, int) {
	startPos := currentPos
	result += num / 100

	num = num % 100

	if direction == 'L' {
		currentPos -= num
	} else {
		currentPos += num
	}

	if currentPos < 0 {
		currentPos += 100
		if currentPos != 0 && startPos != 0 {
			result++
		}
	} else if currentPos > 99 {
		currentPos -= 100
		if currentPos != 0 {
			result++
		}
	}

	if currentPos == 0 {
		result++
	}

	return result, currentPos
}
