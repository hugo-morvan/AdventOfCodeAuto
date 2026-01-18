package main

import (
	"fmt"
	"strings"
	"github.com/hugo-morvan/aoc/utils"
	"strconv"
)

func main() {
	lines := utils.ReadLines(utils.InputFile())

	fmt.Println("Part 1:", solvePart1(lines))
	fmt.Println("Part 2:", solvePart2(lines))
}

func solvePart1(lines []string) int {
	splittedInput := strings.Split(lines[0], ",") 
	count := 0
	for idx, idRanges := range splittedInput {
		fmt.Printf("%v: %v\n", idx, idRanges)

		beginNend := strings.Split(idRanges, "-")
		
		begin, _ := strconv.Atoi(beginNend[0])
		end, _ := strconv.Atoi(beginNend[1])
		
		for i := begin; i <= end; i += 1 {
			// fmt.Println(i)
			str := strconv.Itoa(i)
			if len(str) % 2 == 1 {
				continue
			}
			mid := len(str)/2 
			first, second := str[:mid], str[mid:]
			if first == second {
				count += i
			}
		} 
	}
	return count
}

func solvePart2(lines []string) int {
	count := 0
	splittedInput := strings.Split(lines[0], ",")

	for idx , idRanges := range splittedInput {
		fmt.Println(idx)
		x := strings.Split(idRanges, "-")
		begin, _ := strconv.Atoi(x[0])
		end, _ := strconv.Atoi(x[1])
		
		for i := begin; i <= end; i+=1 {
			// for each id ranges[begin, end]
			s := strconv.Itoa(i)
			maxSub := len(s) / 2
			for subLen := 1; subLen <= maxSub; subLen += 1 {
				//fmt.Println("-----")
				//fmt.Println(s)
				//fmt.Println(subLen)
				// For each subtrings length possible
				if len(s) % subLen != 0 {
					// if not perfectly divisible, skip this substring length
					continue
				}
				allEquals := true
				numSubs := len(s)/subLen
				// check if subsequent substrings are equal.
				prev := s[0:subLen]
				for j := 1; j < numSubs; j +=1{
					next := s[j*subLen : j*subLen+subLen] // start and end of j-th substring  
					//fmt.Printf("prev: %v, next: %v\n", prev, next)

					if prev != next{
						// not all substrings are equal, go to next sub size
						allEquals = false
						break
					}
					prev = next

				}
				if allEquals{ // no need to check the rest (avoid double counts)
					//fmt.Println("INVALID!")
					count += i
					break
				}
			}
		}
	}
	return count
}
