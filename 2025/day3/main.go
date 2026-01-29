package main

import (
	"fmt"
	"strconv"
	"github.com/hugo-morvan/aoc/utils"
)

func main() {
	lines := utils.ReadLines(utils.InputFile())

	fmt.Println("Part 1:", solvePartOne(lines))
	fmt.Println("Part 2:", solvePartTwo(lines))
}

func solvePartOne(lines []string) int {
	total := 0
	for _, line := range lines {
		n := len(line)
		best := -1
		maxRight := -1

		for i := n-1; i >= 0; i-- {
			d := int(line[i] - '0')

			if maxRight != -1 {
				value := d*10 + maxRight
				if value > best {
					best = value
				}
			}

			if d > maxRight {
				maxRight = d
			}
		}
		if best == -1 {
			total += 0
		}
		total += best
	}


	return total
}

func solvePartTwo(lines []string) int {
	total := 0 
	// Greedy monotonic stack
	for _, line := range lines {
		k := 12
		n := len(line)

		if n < k {
			return 69
		}

		drop := n-k
		stack := make([]byte, 0, n)

		for i := 0; i<n; i++ {
			d := line[i]

			for drop > 0 && len(stack) > 0 && stack[len(stack)-1] < d {
				stack = stack[:len(stack)-1]
				drop--

			}

			stack = append(stack, d)
		}
		dig, _ := strconv.Atoi(string(stack[:k]))
		total += dig
	}


	return total
}
