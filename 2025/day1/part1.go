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
		fmt.Printf("Dir: %s, Dist: %d\n", instr.Dir, instr.Dist)
		if instr.Dir == "L" {
			position -= instr.Dist
		} else {
			position += instr.Dist
		}
		position = ((position % 100) + 100) % 100

		if position == 0 {
			counter++
		}
		println("Position:", position)
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
