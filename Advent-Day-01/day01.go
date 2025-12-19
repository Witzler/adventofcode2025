package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const DIALMAX int = 100  // max degrees on dial //like #define in C
const DIALSTART int = 50 // starting position of the dial

var currentDialPosition int = DIALSTART // current position of the dial

// part one variables
var passwordCountEndedOnZero int = 0 // final password counter //when dial hits 0
// part two variables
var passwordCountPassedZero int = 0 // final password counter //when dial passes 0

func main() {

	// read file line by line
	file, err := os.Open("input.txt") // input or test.txt
	if err != nil {
		log.Fatalf("Error when opening file: %v", err) // log fatal and exit
	}
	defer file.Close() // close file at the end

	// Scanner to read file line by line
	scanner := bufio.NewScanner(file)
	// var line string
	for scanner.Scan() {
		// line = scanner.Text()
		line := scanner.Text() // line as string
		//fmt.Println(line)

		if len(line) < 2 {
			log.Fatalf("Invalid line: %v", line)
		}

		direction := string(line[0]) // first character}
		var degrees int
		degreesS := line[1:] // rest of the string
		degrees, err = strconv.Atoi(degreesS)
		if err != nil {
			log.Fatalf("Invalid degrees: %v", degreesS)
		}

		turnDialPartTwo(direction, degrees) //Goes first cause doesn't change variable outside of function
		turnDialPartOne(direction, degrees)

	}
	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error when reading file: %v", err)
	}

	// Output final password counter
	//log.Printf("Puzzle input: %v", passwordCount)
	//log.Printf("Puzzle input: %d	", passwordCount)
	//fmt.Printf("Puzzle input: %v\n", passwordCount)
	fmt.Printf("Puzzle 1 input: %d\n", passwordCountEndedOnZero)
	data := []byte(strconv.Itoa(passwordCountEndedOnZero))
	os.WriteFile("inputSolution.txt", data, 0644)

	fmt.Printf("Puzzle 2 input: %d\n", passwordCountPassedZero)
	data2 := []byte(strconv.Itoa(passwordCountPassedZero))
	os.WriteFile("inputSolution2.txt", data2, 0644)

}

func turnDialPartOne(direction string, degrees int) int { //part one turn of the dial
	var outputDialPosition int

	switch direction {
	case "L":
		outputDialPosition = ((currentDialPosition-degrees)%DIALMAX + DIALMAX) % DIALMAX
	case "R":
		outputDialPosition = (currentDialPosition + degrees) % DIALMAX
	default:
		log.Fatalf("Invalid direction: %v", direction)
		fmt.Println("Invalid direction:", direction)
	}

	currentDialPosition = outputDialPosition
	if outputDialPosition == 0 {
		passwordCountEndedOnZero++
	}

	return outputDialPosition
}

func turnDialPartTwo(direction string, degrees int) {
	var amountPassedZero int = 0
	var outputDialPosition int = 0

	amountPassedZero = degrees / DIALMAX
	remainderDegrees := degrees % DIALMAX

	switch direction {
	case "L":
		outputDialPosition = ((currentDialPosition-remainderDegrees)%DIALMAX + DIALMAX) % DIALMAX
		if ((outputDialPosition > currentDialPosition) && (currentDialPosition != 0)) || outputDialPosition == 0 {
			amountPassedZero++
		}

	case "R":
		outputDialPosition = (currentDialPosition + remainderDegrees) % DIALMAX
		if (currentDialPosition+remainderDegrees >= DIALMAX) || (outputDialPosition == 0) {
			amountPassedZero++
		}

	default:
		log.Fatalf("Invalid direction: %v", direction)
		fmt.Println("Invalid direction:", direction)
	}

	//fmt.Printf("Dir: %s Deg: %d From: %d To: %d PassedZero: %d\n", direction, degrees, currentDialPosition2, outputDialPosition, amountPassedZero)

	passwordCountPassedZero += amountPassedZero
}
