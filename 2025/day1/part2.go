package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Instruction struct {
	Dir  string
	Dist int
}

func main() {
	instructions := readInput("input.txt")

	// Example usage
	counter := 0
	position := 50

	for _, instr := range instructions {
		dist := instr.Dist

		if instr.Dir == "R" {
			for dist > 0 {
				// distance to next 0 going right
				toZero := (100 - position) % 100
				if toZero == 0 {
					toZero = 100
				}

				if dist >= toZero {
					counter++
					dist -= toZero
					position = 0
				} else {
					position = (position + dist) % 100
					dist = 0
				}
			}
		} else { // L
			for dist > 0 {
				// distance to next 0 going left
				toZero := position
				if toZero == 0 {
					toZero = 100
				}

				if dist >= toZero {
					counter++
					dist -= toZero
					position = 0
				} else {
					position = ((position-dist)%100 + 100) % 100
					dist = 0
				}
			}
		}
	}
	fmt.Println("Answer:", counter)

}

func readInput(filename string) []Instruction {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	instructions := []Instruction{}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		dir := string(line[0])
		dist, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}

		instructions = append(instructions, Instruction{
			Dir:  dir,
			Dist: dist,
		})
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return instructions
}
