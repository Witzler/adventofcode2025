package main

import (
	"bufio"
	"os"
	"strconv"
)

// Answer Test part1:
// Answer Test part2:
// Answer for part1:
// Answer for part2:

func main() {
	// Your code here

	freshIngredients := 0
	freshIngredientsPartTwo := 0

	scanner := bufio.NewScanner(os.Stdin) // cat assets/input.txt | go run main.go
	for scanner.Scan() {
		row := scanner.Text()

		_ = row
		// Process the row here

	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	//t.Printf("Total largest joltage: %d\n", accessableRolls)
	//t.Printf("Total for part 2: %d\n", accessableRollsPartTwo)
	os.Stdout.WriteString("Total largest joltage: " + strconv.Itoa(accessableRolls) + "\n")   // alternative to fmt.Printf, less exec time
	os.Stdout.WriteString("Total for part 2: " + strconv.Itoa(accessableRollsPartTwo) + "\n") // alternative to fmt.Printf, less exec time

}
