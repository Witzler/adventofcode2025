package utils

import (
	"fmt"
	"strconv"
)

func FindInvalidIDS(s1 string, s2 string) int {

	var output int = 0

	num1, err := strconv.Atoi(s1)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
		return 0
	}

	num2, err := strconv.Atoi(s2)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
		return 0
	}

	var index int = num2 - num1
	var temp int = num1

	for i := 0; i <= index; i++ {
		var isInvalid bool = true
		temptS := strconv.Itoa(temp)
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

	return output
}

func ExampleFunction() {
	// Placeholder function for day 2 utilities
	fmt.Println("This is a utility function for Day 2")
}
