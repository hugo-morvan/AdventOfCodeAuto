# Advent of Code Template

Automated AoC setup and utilities.

## What's included
- `fetch.go` - Creates day folders, downloads inputs/examples
- `utils/input.go` - Input reading utilities with test flag support
- `templates/go/main.go` - Solution template with Part 1/2 structure

## Usage
```bash
# Set up session token
cp env.example .env  # add your AoC session cookie

# Fetch a problem (creates folder structure, downloads input/examples)
go run fetch.go 2025 1

# Solve
cd 2025/day1
go run main.go      # real input
go run main.go -t   # test input
```