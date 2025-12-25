package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func SolveWorksheetPartOne(entireWorksheet [][]string) (output int) {
	numRows := len(entireWorksheet)
	numCols := len(entireWorksheet[0])

	for col := 0; col < numCols; col++ { // über Spalten
		temp := 0
		tempoCount := 0
		operator := entireWorksheet[numRows-1][col]

		for row := 0; row < numRows-1; row++ { // über Zeilen
			value := entireWorksheet[row][col]
			valueInt, _ := strconv.Atoi(value)

			switch operator {
			case "+":
				temp += valueInt
			case "*":
				if tempoCount == 0 {
					temp = 1
					tempoCount++
				}
				temp *= valueInt

			}
		}
		output += temp
	}

	return
}

func SolveWorksheetPartTwo(entireWorksheet [][]string) (output int) {
	numRows := len(entireWorksheet)
	numCols := len(entireWorksheet[0])
	numArguments := numRows - 1

	emptyColsCounter := 0
	var newNumberCollection []string

	operatorCollected := 0
	var operators []string
	for _, c := range entireWorksheet[numRows-1] { // extract operators for easy access
		if c == "*" || c == "+" {
			operators = append(operators, c)
		}
	}

	for i := 0; i < numCols; i++ {
		newNumber := ""

		for j := 0; j < numArguments; j++ { // column number building
			newNumber += entireWorksheet[j][i]
		}

		s := strings.TrimSpace(newNumber)
		if s == "" {
			emptyColsCounter++
		} else {
			newNumberCollection = append(newNumberCollection, newNumber)
			emptyColsCounter = 0
		}

		if (emptyColsCounter == 1 || i == numCols-1) && !(len(newNumberCollection) == 0) {

			operation := operators[operatorCollected]
			operatorCollected++

			switch operation {
			case "+":
				va := 0
				for _, val := range newNumberCollection {
					val = strings.TrimSpace(val)
					valInt, err := strconv.Atoi(val)
					if err != nil {
						fmt.Println("Invalid value:", val)
						continue
					}
					va += valInt
				}

				output += va

				newNumberCollection = []string{} // or nil
				emptyColsCounter = 0
				newNumber = ""

			case "*":
				va := 1
				for _, val := range newNumberCollection {
					val = strings.TrimSpace(val)
					valInt, err := strconv.Atoi(val)
					if err != nil {
						fmt.Println("Invalid value:", val)
						continue
					}
					va *= valInt
				}
				output += va

				newNumberCollection = []string{} // or nil
				emptyColsCounter = 0
				newNumber = ""
			default:
				fmt.Println("Unknown operation:", operation)
				continue
			}
		}
	}
	return
}

/*
func SolveWorksheetPartTwo(entireWorksheet [][]string) (output int) {
	numRows := len(entireWorksheet)
	numCols := len(entireWorksheet[0])

	numArguments := numRows - 1

	for col := 0; col < numCols; col++ {
		var tempTask []string
		tempRows := len(tempTask)
		value := ""

		maxValueLength := 0
		operation := ""
		operation = entireWorksheet[numRows-1][col]

		for row := 0; row < numArguments; row++ {
			value = entireWorksheet[row][col]
			if len(value) > maxValueLength {
				maxValueLength = len(value)
			}
		}
		for tempRows < maxValueLength { // I don't care
			tempTask = append(tempTask, "")
			tempRows++
			tempRows = len(tempTask)
		}

		switch operation {
		case "+":
			for row := 0; row < numArguments; row++ {
				value := entireWorksheet[row][col]
				valueLength := len(value)
				for i := 0; i < valueLength; i++ {
					tempTask[i] = tempTask[i] + string(value[i])
					//fmt.Println("TempTask:", tempTask)
				}
			}

			output += helperAdd(tempTask)

		case "*":
			for row := 0; row < numArguments; row++ {
				value := entireWorksheet[row][col]
				valueLength := len(value)

				for i := valueLength - 1; i >= 0; i-- {
					tempTask[valueLength-1-i] = tempTask[valueLength-1-i] + string(value[i])
					//fmt.Println("TempTask:", tempTask)
				}

			}
			output += helperMultiply(tempTask)
		default:
			fmt.Println("Unknown operation:", operation)
			continue
		}

		//fmt.Println("Operation:", operation)

	}

	return
}

func helperAdd(input []string) (tempOutput int) {
	for _, val := range input {
		valInt, _ := strconv.Atoi(val)
		tempOutput += valInt
	}
	return
}
func helperMultiply(input []string) (tempOutput int) {
	tempOutput = 1
	for _, val := range input {
		valInt, _ := strconv.Atoi(val)
		tempOutput *= valInt
	}
	return
}
*/
