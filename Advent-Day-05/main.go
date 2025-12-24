package main

import (
	"adventofcode2025day05/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Answer Test part1: 3
// Answer Test part2: 14
// Answer for part1: 601
// Answer for part2: 367899984917516

func main() {
	freshIngredients := 0
	freshIngredientsPartTwo := 0

	availableIngredients := false
	db := utils.DataBase{Ranges: make([]utils.Range, 0)}
	availableIngredientsList := make([]int, 0)

	scanner := bufio.NewScanner(os.Stdin) // cat assets/input.txt | go run main.go
	for scanner.Scan() {
		row := scanner.Text()

		if row == "" { //is empty line
			availableIngredients = true
		}

		if !availableIngredients {
			start := 0
			end := 0
			_, err := fmt.Sscanf(row, "%d-%d", &start, &end)
			if err != nil {
				fmt.Println("Error parsing range:", err)
				continue
			}
			newRange := utils.Range{}.CreateRange(start, end) // Using the CreateRange method // // Alternative: newRange :=utils.Range{Start: start, End: end}
			db.AddRange(newRange)
		} else {
			ingredient := 0
			_, err := fmt.Sscanf(row, "%d", &ingredient)
			if err != nil {
				continue
			}
			availableIngredientsList = append(availableIngredientsList, ingredient)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	freshIngredients = utils.CountFreshIngredients(db, availableIngredientsList)
	freshIngredientsPartTwo = utils.ConsideredIngredients(db)

	//t.Printf()
	os.Stdout.WriteString("Fresh ingredients: " + strconv.Itoa(freshIngredients) + "\n")                   // alternative to fmt.Printf, less exec time
	os.Stdout.WriteString("Considered fresh ingredients: " + strconv.Itoa(freshIngredientsPartTwo) + "\n") // alternative to fmt.Printf, less exec time

}
