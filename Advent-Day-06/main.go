package main

// Answer Test part1: 4277556
// Answer Test part2: 3263827
// Answer for part1: 5524274308182
// Answer for part2: 8843673199391

import (
	"adventofcode2025day05/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	worksheetSolution := 0
	worksheetSolutionPartTwo := 0

	var allFields [][]string
	var charTable [][]string

	scanner := bufio.NewScanner(os.Stdin) // cat assets/input.txt | go run main.go

	for scanner.Scan() {
		row := scanner.Text()

		fields := strings.Fields(row)
		allFields = append(allFields, fields)

		line := make([]string, len(row))
		for i, r := range row {
			line[i] = string(r)
		}
		charTable = append(charTable, line)

	}
	//fmt.Println(charTable)

	worksheetSolution = utils.SolveWorksheetPartOne(allFields)
	fmt.Printf("Worksheet solution part one: %d\n", worksheetSolution)

	worksheetSolutionPartTwo = utils.SolveWorksheetPartTwo(charTable)
	fmt.Printf("Worksheet solution part two: %d\n", worksheetSolutionPartTwo)
}
