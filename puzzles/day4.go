package puzzles

import (
	"fmt"
	"steveandkureen/advent_2025/internal"
	"strings"
)

func FindAdjecentRolls(filename string, isHard bool) (string, error) {
	data, err := internal.ReadFile(filename)
	if err != nil {
		internal.LogOutput("error reading file: %s\n", err)
	}
	internal.LogOutput("The map: \n%s", data)
	room := parseMap(data)
	room2 := parseMap(data)
	result := 0
	for {
		newResult := 0
		for r, row := range room {
			for c, spot := range row {
				if spot == '@' {
					count := checkAdjecent(room, r, c)
					if count < 4 {
						newResult++
						room[r][c] = 'x'
					}
				}
			}
		}
		result += newResult
		if newResult == 0 || !isHard {
			break
		}
	}
	outputMap(room2)
	return fmt.Sprintf("%d", result), nil
}

func checkAdjecent(room [][]rune, row int, column int) int {
	count := 0
	for r := -1; r <= 1; r++ {
		if (row+r < 0) ||
			(row+r > len(room)-1) {
			continue
		}
		for c := -1; c <= 1; c++ {
			if (r == 0 && c == 0) ||
				(column+c < 0) ||
				(column+c > len(room[row+r])-1) {
				continue
			}
			if room[row+r][column+c] == '@' {
				count++
			}
			// internal.LogOutput("Spot[%d][%d]:%s count:%d\n", row+r, column+c, string(room[row+r][column+c]), count)
		}
	}
	return count
}

func parseMap(data string) [][]rune {
	result := [][]rune{}

	for _, line := range strings.Split(data, "\n") {
		runeLine := []rune(strings.Trim(line, "\n"))
		result = append(result, runeLine)
	}
	return result
}

func outputMap(room [][]rune) {
	for r := range room {
		for c := range room[r] {
			fmt.Printf("%s", string(room[r][c]))
		}
		fmt.Printf("\n")
	}
}
