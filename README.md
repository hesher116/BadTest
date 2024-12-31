# Longest Fragment Sequence Finder

A program that finds the longest possible sequence of text fragments where each subsequent fragment begins with the last two characters of the previous one.

## Requirements

- Go 1.20 or newer (https://golang.org/dl/)

## Installation

1. Clone the repository:

git clone git@github.com:hesher116/BadTest.git

2. Navigate to the project directory:

cd [directory-name]

## Usage

1. Place your fragments in the `fragments.txt` file (one fragment per line)

2. Run the program:

go run main.go

## Input Format

The `fragments.txt` file should contain text fragments, one per line.

## Output Format

The program will display:
- The longest sequence found
- Number of fragments used

## License

MIT License