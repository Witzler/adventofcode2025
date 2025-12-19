package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func FindInvalidIDS(s1 string, s2 string) (int, int) {

	var output int = 0
	var output2 int = 0

	num1, err := strconv.Atoi(s1)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
		return 0, 0 // Return 0 on error
	}

	num2, err := strconv.Atoi(s2)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
		return 0, 0 // Return 0 on error
	}

	var index int = num2 - num1
	var temp int = num1

	for i := 0; i <= index; i++ {
		var isInvalid bool = true
		temptS := strconv.Itoa(temp)

		// checking every number for part 2; accumulate matches instead of overwriting
		if res := FindInvalidIDS_partTwo(temptS, temp); res != 0 {
			output2 += res
		}

		if len(temptS)%2 == 0 {
			tempLen := len(temptS)
			var x int = 0
			var y int = (tempLen / 2)
			for y < tempLen {
				if temptS[x] != temptS[y] {
					isInvalid = false
					break
				}
				x++
				y++
			}
			if isInvalid {
				output = output + temp
			}
		}
		temp++

	}

	return output, output2
}

/*
	TODO : Fix function to return the number if it is invalid

It is invalid if it contains the same digits consecutively like 11 1212  111 1414 824824824
*/
func FindInvalidIDS_partTwo(s1 string, s1Int int) int {
	n := len(s1)
	// Only consider numbers invalid when the entire string is a repetition
	// of a substring p (p repeated >= 2 times). Examples: "111", "1010", "824824824".
	for l := 1; l <= n/2; l++ {
		if n%l != 0 {
			continue
		}
		if strings.Repeat(s1[0:l], n/l) == s1 {
			return s1Int
		}
	}

	return 0
}

func ExampleFunction() {
	// Placeholder function for day 2 utilities
	fmt.Println("This is a utility function for Day 2")
}
