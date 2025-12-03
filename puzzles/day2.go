// Package puzzles contains the daily puzzels to be executed
package puzzles

import (
	"errors"
	"fmt"
	"steveandkureen/advent_2025/internal"
	"strconv"
	"strings"
)

func CheckProdcutIds(filename string, isHard bool) (string, error) {
	result := 0
	data, err := internal.ReadFile(filename)
	if err != nil {
		fmt.Printf("error reading file: %s\n", err)
		return "", err
	}
	fmt.Printf("Input: %s\n", string(data))

	for _, part := range strings.Split(data, ",") {
		start, end, err := ParseRange(part)
		if err != nil {
			fmt.Printf("Error Parsing Range: %s %s\n", part, err)
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
		fmt.Printf("Found Invalid: %s\n", value)
		return test
	}
	return 0
}

func checkFullPatternMatch(test int) int {
	value := strconv.FormatInt(int64(test), 10)
	// fmt.Printf("Testing: %s\n", value)
	for i := 1; i <= len(value)/2; i++ {
		current := value[0:i]
		// fmt.Printf("Current Sub: %d %s\n", i, current)
		if len(value)%i != 0 {
			// fmt.Printf("Skipping %d doesnt mod\n", i)
			continue
		}
		match := true
		for x := 1; x < len(value)/i; x++ {
			startIndex := (i * x)
			// fmt.Printf("Test: %d %s\n", x, value[startIndex:startIndex+i])
			if current != value[startIndex:startIndex+i] {
				// fmt.Printf("No sub Match: %s\n", value)
				match = false
				break
			}
		}
		if match {
			// fmt.Printf("***Match: %s\n", value)
			return test
		}
	}
	// fmt.Printf("No Match: %s\n", value)
	return 0
}

func ParseRange(input string) (int, int, error) {
	parts := strings.Split(strings.Trim(input, "\n"), "-")

	if len(parts) != 2 {
		fmt.Printf("input not 2 parts: %d\n", len(parts))
		return -1, -1, errors.New("invalid input")
	}

	start, err := strconv.Atoi(parts[0])
	if err != nil {
		fmt.Printf("error parsing input: %s\n", err)
		return -1, -1, err
	}
	end, err := strconv.Atoi(parts[1])
	if err != nil {
		fmt.Printf("error parsing input: %s\n", err)
		return -1, -1, err
	}
	return start, end, nil
}
