package puzzles

import (
	"fmt"
	"steveandkureen/advent_2025/internal"
	"strings"
)

func FindMaxVoltage(filename string, isHard bool) (string, error) {
	data, err := internal.ReadFile(filename)
	if err != nil {
		internal.LogOutput("error reading file: %s\n", err)
	}
	// internal.LogOutput("Data: %s\n", data)
	banks := strings.Split(data, "\n")

	result := 0
	for _, bank := range banks {
		voltage := 0
		if !isHard {
			voltage = findBankMax(bank)
		} else {
			voltage = findBankMaxHard(bank, 12)
		}
		result += voltage
		// internal.LogOutput("bank: %s, voltage: %d\n", bank, voltage)
	}

	return fmt.Sprintf("%d", result), nil
}

func findBankMax(bank string) int {
	max1 := 0
	max2 := 0

	for i, v := range bank {
		voltage := int(v - '0')
		if voltage > max1 && i < len(bank)-1 {
			max2 = int(bank[i+1] - '0')
			max1 = voltage
		} else if voltage > max2 {
			max2 = voltage
		}
	}
	return (max1 * 10) + max2
}

func findBankMaxHard(bank string, size int) int {
	result := make([]int, size)
	for x := range result {
		result[x] = 0
	}
NEXTVOLTAGE:
	for i, v := range bank {
		voltage := int(v - '0')
		xStart := 0
		if len(bank)-i < size {
			xStart = size - (len(bank) - i)
		}
		for x := xStart; x < len(result); x++ {
			if voltage > result[x] {
				result[x] = voltage
				for y := x + 1; y < size; y++ {
					result[y] = 0
				}
				continue NEXTVOLTAGE
			}
		}
	}
	r := 0
	for _, digit := range result {
		r = r*10 + digit
	}
	return r
}
