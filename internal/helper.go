// Package internal contains helper functions need by the puzzles
package internal

import (
	"os"
)

func ReadFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
