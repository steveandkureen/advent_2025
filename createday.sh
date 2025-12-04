#!/bin/bash

# Check if argument is provided
if [ -z "$1" ]; then
  echo "Error: Please provide a day number"
  echo "Usage: ./createday.sh <day_number>"
  exit 1
fi

DAY=$1

# Create the puzzles directory if it doesn't exist
mkdir -p puzzles

# Create the data/day{x} directory
mkdir -p "data/day${DAY}"

# Create the Go file in puzzles directory

cat >puzzles/day${DAY}.go <<'EOF'
  package puzzles
  import (
 	"steveandkureen/advent_2025/internal"
)
  func Startfunc(filename string, isHard bool) int {
    data, err := internal.ReadFile(filename)
    if err != nil {
      internal.LogOutput("error reading file: %s\n", err)
    }
  return 0
}
EOF

cat >puzzles/day${DAY}_test.go <<'EOF'
  package puzzles
  
  import "testing"

  func startfunc() int {
    data, err := internal.ReadFile(filename)
    if err != nil {
      internal.LogOutput("error reading file: %s\n", err)
    }
}
EOF

# Create the text files in data/day{x} directory
touch "data/day${DAY}/easy.txt"
touch "data/day${DAY}/example.txt"

echo "Successfully created:"
echo "  - puzzles/day${DAY}.go"
echo "  - data/day${DAY}/easy.txt"
echo "  - data/day${DAY}/example.txt"
