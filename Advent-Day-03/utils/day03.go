package utils

import (
	"strconv"
)

type Battery struct {
	Voltage int
	Index   int
}

const BATTERYCHOICE int = 12

func LargestJoltage(bank string) int {
	numberBatteries := len(bank)
	maxJoltage := 0
	maxJolString := ""
	temp := 0
	var battery1 string = ""
	battPos1 := 0
	var battery2 string = ""
	//digPos2 := 0

	for i := 0; i < numberBatteries-1; i++ {
		char := bank[i]
		joltage := int(char - '0') // Convert char digit to int
		if joltage > temp {
			temp = joltage
			battPos1 = i
			battery1 = string(char)
		}
	}
	temp = 0
	for i := battPos1 + 1; i < numberBatteries; i++ {
		char := bank[i]
		joltage := int(char - '0') // Convert char digit to int
		if joltage > temp {
			temp = joltage
			battery2 = string(char)
		}
	}
	maxJolString = battery1 + battery2
	maxJoltage, _ = strconv.Atoi(maxJolString)
	return maxJoltage
}

func FindInvalidIDS_partTwo(bank string) int64 { //stack-based approach to remove digits
	numBatteries := len(bank)
	var maxVoltage int64 = 0
	stack := make([]byte, 0, numBatteries) // Stack to hold the selected batteries

	remove := numBatteries - BATTERYCHOICE
	if remove < 0 {
		panic("Bank length is less than 12")
	}

	for i := 0; i < numBatteries; i++ {
		c := bank[i]
		// Remove smaller digits on the left if a larger one comes on the right
		for remove > 0 && len(stack) > 0 && stack[len(stack)-1] < c {
			stack = stack[:len(stack)-1]
			remove--
		}
		stack = append(stack, c)
	}

	// If digits still need to be removed, cut off the end
	if len(stack) > BATTERYCHOICE {
		stack = stack[:BATTERYCHOICE]
	}

	// Convert stack to number
	for _, d := range stack {
		maxVoltage = maxVoltage*10 + int64(d-'0')
	}

	return maxVoltage
}
