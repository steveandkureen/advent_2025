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
touch "puzzles/day${DAY}.go"

# Create the text files in data/day{x} directory
touch "data/day${DAY}/easy.txt"
touch "data/day${DAY}/example.txt"

echo "Successfully created:"
echo "  - puzzles/day${DAY}.go"
echo "  - data/day${DAY}/easy.txt"
echo "  - data/day${DAY}/example.txt"
