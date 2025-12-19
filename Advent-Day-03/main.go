package main

// Answer Test part1: 357
// Answer Test part2: 0
// Answer for part1: 17343
// Answer for part2:

import (
	"adventofcode2025day03/utils" //d03 "adventofcode2025day03/utils"  => d03.LargestJoltage(line)
	"bufio"
	"fmt"
	"os"
)

var totalJoltage int = 0
var totalJoltage2 int = 0

func main() {
	file, err := os.Open("assets/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bank := scanner.Text()
		//fmt.Println(line)
		totalJoltage = totalJoltage + utils.LargestJoltage(bank)
		totalJoltage2 = totalJoltage2 + 0 // Placeholder for part 2
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		panic(err)
	}

	fmt.Printf("Total largest joltage: %d\n", totalJoltage)
	fmt.Printf("Total for part 2: %d\n", totalJoltage2)

}
