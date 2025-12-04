// Package puzzles contains the daily puzzels to be executed
package puzzles

import (
	"errors"
	"steveandkureen/advent_2025/internal"
	"strconv"
	"strings"
)

func CheckProdcutIds(filename string, isHard bool) (string, error) {
	result := 0
	data, err := internal.ReadFile(filename)
	if err != nil {
		internal.LogOutput("error reading file: %s\n", err)
		return "", err
	}
	internal.LogOutput("Input: %s\n", string(data))

	for _, part := range strings.Split(data, ",") {
		start, end, err := ParseRange(part)
		if err != nil {
			internal.LogOutput("Error Parsing Range: %s %s\n", part, err)
			return "", err
		}
		for test := start; test <= end; test++ {
			if test < 10 {
				continue
			}
			if !isHard {
				result += checkHalfPatternMatch(test)
			} else {
				result += checkFullPatternMatch(test)
			}

		}
	}
	return strconv.FormatInt(int64(result), 10), nil
}

func checkHalfPatternMatch(test int) int {
	value := strconv.FormatInt(int64(test), 10)
	if len(value)%2 != 0 {
		// it wont split evenly
		return 0
	}
	half1 := value[0 : len(value)/2]
	half2 := value[len(value)/2:]
	if half1 == half2 {
		internal.LogOutput("Found Invalid: %s\n", value)
		return test
	}
	return 0
}

func checkFullPatternMatch(test int) int {
	value := strconv.FormatInt(int64(test), 10)
	internal.LogOutput("Testing: %s\n", value)
	for i := 1; i <= len(value)/2; i++ {
		current := value[0:i]
		internal.LogOutput("Current Sub: %d %s\n", i, current)
		if len(value)%i != 0 {
			internal.LogOutput("Skipping %d doesnt mod\n", i)
			continue
		}
		match := true
		for x := 1; x < len(value)/i; x++ {
			startIndex := (i * x)
			internal.LogOutput("Test: %d %s\n", x, value[startIndex:startIndex+i])
			if current != value[startIndex:startIndex+i] {
				internal.LogOutput("No sub Match: %s\n", value)
				match = false
				break
			}
		}
		if match {
			internal.LogOutput("***Match: %s\n", value)
			return test
		}
	}
	internal.LogOutput("No Match: %s\n", value)
	return 0
}

func ParseRange(input string) (int, int, error) {
	parts := strings.Split(strings.Trim(input, "\n"), "-")

	if len(parts) != 2 {
		internal.LogOutput("input not 2 parts: %d\n", len(parts))
		return -1, -1, errors.New("invalid input")
	}

	start, err := strconv.Atoi(parts[0])
	if err != nil {
		internal.LogOutput("error parsing input: %s\n", err)
		return -1, -1, err
	}
	end, err := strconv.Atoi(parts[1])
	if err != nil {
		internal.LogOutput("error parsing input: %s\n", err)
		return -1, -1, err
	}
	return start, end, nil
}
