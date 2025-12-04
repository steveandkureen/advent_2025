// Package internal contains helper functions need by the puzzles
package internal

import (
	"fmt"
	"os"
)

var ShowOutput bool = true

func ReadFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func LogOutput(format string, args ...any) {
	if !ShowOutput {
		return
	}
	fmt.Printf(format, args...)
}
