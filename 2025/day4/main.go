package main

import (
	"fmt"
	//"strings"
	"github.com/hugo-morvan/aoc/utils"
)

func main() {
	lines := utils.ReadLines(utils.InputFile())

	fmt.Println("Part 1:", solvePart1(lines))
	fmt.Println("Part 2:", solvePart2(lines))
}

// Function to add a dot padding around the grid
func padGrid(grid []string) []string {
	if len(grid) == 0 {
		return grid
	}

	w := len(grid[0])
	padded := make([]string, 0, len(grid)+2)

	// top border
	border := makeDots(w + 2)
	padded = append(padded, border)

	// middle rows
	for _, row := range grid {
		padded = append(padded, "."+row+".")
	}

	// bottom border
	padded = append(padded, border)

	return padded
}

func makeDots(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = '.'
	}
	return string(b)
}

func convolve3x3(grid []string, threshold int) []string {
	h := len(grid)
	if h == 0 {
		return grid
	}
	w := len(grid[0])

	// Copy grid into mutable form
	out := make([][]byte, h)
	for i := range grid {
		out[i] = []byte(grid[i])
	}

	totalRows := h - 2
	if totalRows <= 0 {
		return grid
	}

	for i := 1; i < h-1; i++ {
		for j := 1; j < w-1; j++ {
			count := 0

			// 3x3 window
			for di := -1; di <= 1; di++ {
				for dj := -1; dj <= 1; dj++ {
					c := grid[i+di][j+dj]
					if c == '@' || c == 'x' {
						count++
					}
				}
			}

			if count <= threshold && out[i][j]=='@' {
				out[i][j] = 'x'
			}
		}

	}

	// Convert back to []string
	result := make([]string, h)
	for i := range out {
		result[i] = string(out[i])
	}

	return result
}

func countX(grid []string) int {
	count := 0
	for _, row := range grid {
		for i := 0; i < len(row); i++ {
			if row[i] == 'x' {
				count++
			}
		}
	}
	return count
}


func replaceXWithDot(grid []string) []string {
	out := make([]string, len(grid))

	for i, row := range grid {
		b := []byte(row)
		for j := range b {
			if b[j] == 'x' {
				b[j] = '.'
			}
		}
		out[i] = string(b)
	}

	return out
}

func solvePart1(lines []string) int {
	total := 0
	paddedGrid := padGrid(lines)
	xGrid := convolve3x3(paddedGrid, 4)
	total += countX(xGrid)
	return total
}

func solvePart2(lines []string) int {
	total := 0
	prev_total := -1
	paddedGrid	:= padGrid(lines)
	for total != prev_total {
		xGrid := convolve3x3(paddedGrid, 4)
		prev_total = total
		total += countX(xGrid)
		fmt.Println(total)
		paddedGrid = replaceXWithDot(xGrid)
	}
	return total
}
