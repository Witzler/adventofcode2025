package main

// Solution part 1 test: 50
// Solution part 1 input:
// Solution part 2 test:
// Solution part 2 input:
import (
	"adventofcode2025day09/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	var grid [][]int
	scanner := bufio.NewScanner(os.Stdin) // cat assets/input.txt | go run main.go

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ",")
		intLine := make([]int, len(parts))

		for i, p := range parts {
			n, err := strconv.Atoi(p)
			if err != nil {
				log.Fatal(err)
			}
			intLine[i] = n
		}

		grid = append(grid, intLine)
	}

	//fmt.Println("Grid:", grid)

	largestRectangleArea := 0
	largestRectangleArea = utils.FindLargestRectangleArea(grid)
	largestRectangleArea2 := utils.FindLargestRectangleAreaPartTwo(grid)

	fmt.Printf("Largest rectangle area: %d\n", largestRectangleArea)
	fmt.Printf("Largest area part two area part two: %d\n", largestRectangleArea2)

}
