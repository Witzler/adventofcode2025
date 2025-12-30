package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Present = []string

func parse(input string) (presents []Present, puzzles []puzzle) {
	input = strings.ReplaceAll(input, "\r\n", "\n")
	input = strings.TrimSpace(input)
	var parts = strings.Split(input, "\n\n")
	if len(parts) < 2 {
		return
	}
	for _, part := range parts[:len(parts)-1] {
		var lines = strings.Split(part, "\n")
		lines = lines[1:]
		// each present is represented by its lines (after the header)
		presents = append(presents, append([]string{}, lines...))
	}

	for _, line := range strings.Split(parts[len(parts)-1], "\n") {
		if strings.TrimSpace(line) == "" {
			continue
		}
		var x, y, a, b, c, d, e, f int
		fmt.Sscanf(line, "%dx%d: %d %d %d %d %d %d", &x, &y, &a, &b, &c, &d, &e, &f)
		puzzles = append(puzzles, puzzle{x, y, []int{a, b, c, d, e, f}})
	}
	return
}

type puzzle struct {
	sizeX, sizeY int
	quantity     []int
}

// rotations/flip
// 0 init
// 1, 2, 3 left
// 4 flip
// 5, 6, 7 flip, left

type board struct {
	sizeX, sizeY int
}

func nbPixel(p Present) int {
	var res int
	for _, row := range p {
		for _, c := range row {
			if c == '#' {
				res++
			}
		}
	}
	return res
}
func filterPuzzle(puzzles []puzzle, presents []Present) []puzzle {
	var res []puzzle
	for _, p := range puzzles {
		var nb int
		for i, qt := range p.quantity {
			nb += nbPixel(presents[i]) * qt
		}
		if nb <= p.sizeX*p.sizeY {
			res = append(res, p)
			fmt.Printf("diff = %d\n", p.sizeX*p.sizeY-nb)
		}
	}
	return res
}

func Part1(input string) int {
	var presents, puzzles = parse(input)
	fmt.Printf("#puzzles: %v\n", len(puzzles))
	puzzles = filterPuzzle(puzzles, presents)
	fmt.Printf("#puzzles: %v\n", len(puzzles))
	return len(puzzles)
}

func Part2(input string) int {
	presents, puzzles := parse(input)
	puzzles = filterPuzzle(puzzles, presents)
	if len(puzzles) == 0 {
		return 0
	}
	minDiff := -1
	for _, p := range puzzles {
		var nb int
		for i, qt := range p.quantity {
			nb += nbPixel(presents[i]) * qt
		}
		diff := p.sizeX*p.sizeY - nb
		if minDiff == -1 || diff < minDiff {
			minDiff = diff
		}
	}
	if minDiff < 0 {
		return 0
	}
	return minDiff
}

func main() {
	fmt.Println("--2025 day 12 solution--")
	data, err := os.ReadFile("assets/input.txt")
	if err != nil {
		panic(err)
	}
	var inputDay = string(data)

	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
