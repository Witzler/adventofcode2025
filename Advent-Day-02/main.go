package main

//day02 "adventofcode2025day02/utils"
import (
	day02 "adventofcode2025day02/utils"
	"fmt"
	"os"
	"strings"
)

// #define
// global varaibles
var counter int = 0

func main() {

	var counter int

	// Read input file
	data, err := os.ReadFile("assets/input.txt") //test or input
	if err != nil {
		panic(err)
	}

	line := string(data)
	idRange := strings.Split(line, ",") // Split by comma to get ID pairs

	for _, pair := range idRange {
		pair = strings.TrimSpace(pair)
		numbers := strings.Split(pair, "-")
		fmt.Println("First number:", numbers[0], "Second number:", numbers[1])

		counter = counter + day02.FindInvalidIDS(numbers[0], numbers[1])
	}

	fmt.Printf("Total invalid IDs addition: %d\n", counter)
}
