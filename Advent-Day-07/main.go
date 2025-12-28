package main

// Solution for Advent of Code 2025 Day 07
// Solution part 1 test: 21
// Solution part 1 input: 1640
// Solution part 2 test: <not solved yet>
// Solution part 2 input: <not solved yet>

import (
	"adventofcode2025day07/utils"
	"bufio"
	"fmt"
	"os"
)

func main() {
	timesBeamWasSplit := 0
	timesBeamWasSplit2 := 0

	var diagram [][]rune
	scanner := bufio.NewScanner(os.Stdin) // cat assets/input.txt | go run main.go
	for scanner.Scan() {
		line := scanner.Text()
		diagram = append(diagram, []rune(line))
	}

	timesBeamWasSplit = utils.SolveDiagramPartOne(diagram)
	timesBeamWasSplit2 = utils.SolveDiagramPartTwo(diagram)

	fmt.Printf("Times beam was split: %d\n", timesBeamWasSplit)
	fmt.Printf("Beam's timelines: %d\n", timesBeamWasSplit2)

}
