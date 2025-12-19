package main

/* Answer for part1: 12599655151 */
/* Answer for part2: 20942028255*/

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
var counter2 int = 0

func main() {
	// Read input file
	data, err := os.ReadFile("assets/input.txt") // test or input
	if err != nil {
		panic(err)
	}

	line := string(data)
	idRange := strings.Split(line, ",") // Split by comma to get ID pairs

	for _, pair := range idRange {
		pair = strings.TrimSpace(pair)
		numbers := strings.Split(pair, "-") // Split by hyphen to get individual numbers per pair
		fmt.Println("First number:", numbers[0], "Second number:", numbers[1])

		part1, part2 := day02.FindInvalidIDS(numbers[0], numbers[1])
		counter += part1
		counter2 += part2
	}

	fmt.Printf("Total invalid IDs addition: %d\n", counter)

	fmt.Printf("Total invalid IDs addition (part 2): %d\n", counter2)

}
