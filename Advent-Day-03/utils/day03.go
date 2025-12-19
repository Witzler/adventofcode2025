package utils

import "strconv"

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

func FindInvalidIDS_partTwo(bank string) int {}
