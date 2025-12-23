package main

import (
	"adventofcode2025day04/utils"
	"bufio"
	"os"
	"strconv"
)

// Answer Test part1: 13
// Answer Test part2: 43
// Answer for part1: 1553
// Answer for part2: 8442

func main() {
	// Your code here

	accessableRolls := 0
	accessableRollsPartTwo := 0
	gridSlice := make([][]bool, 0) // Placeholder for grid input // initial length 0, can change cause slices are dynamic

	scanner := bufio.NewScanner(os.Stdin) // cat assets/input.txt | go run main.go
	for scanner.Scan() {
		gridRow := scanner.Text()

		gridSlice = append(gridSlice, make([]bool, len(gridRow)))
		gridRowIndex := len(gridSlice) - 1

		for i, r := range gridRow {
			if r == '@' {
				gridSlice[gridRowIndex][i] = true
			} //else if r == '.' {gridSlice[gridRowIndex][i] = false} // unnecessary else cause false is default
		}
	}

	accessableRolls += utils.ForkliftAccess(gridSlice)
	accessableRollsPartTwo += utils.ForkliftAccessPartTwo(gridSlice)

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	//t.Printf("Total reachable rolls: %d\n", accessableRolls)
	//t.Printf("Total reachable rolls part 2:  %d\n", accessableRollsPartTwo)
	os.Stdout.WriteString("Total reachable rolls  " + strconv.Itoa(accessableRolls) + "\n")               // alternative to fmt.Printf, less exec time
	os.Stdout.WriteString("Total reachable rolls part 2: " + strconv.Itoa(accessableRollsPartTwo) + "\n") // alternative to fmt.Printf, less exec time

}
