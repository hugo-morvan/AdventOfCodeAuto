package utils

import (
	"bufio"
	"flag"
	"os"
	"fmt"
)

var useTest = flag.Bool("t", false, "use test input")

func InputFile() string {
	flag.Parse()

	if *useTest {
		fmt.Println("Running with test inputs...")
		return "test.txt"
	}
	fmt.Println("Running with full inputs...")
	return "input.txt"
}

func ReadLines(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines
}
