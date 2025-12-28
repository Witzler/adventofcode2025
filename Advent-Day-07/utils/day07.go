package utils

import (
	"fmt"
)

const manifoldEntrance = 'S'
const emptySpace = '.'
const splitter = '^'
const tachyonBeam = '|'

func SolveDiagramPartOne(diagram [][]rune) (beamSplits int) {
	diagramRows := len(diagram)
	diagramCols := len(diagram[0])

	currentBeams := NewIntSet()

	for i := 0; i < diagramCols; i++ {
		if diagram[0][i] == manifoldEntrance {
			currentBeams.Add(i)
			break
		}
	}

	for row := 1; row < diagramRows; row++ {
		nextBeams := NewIntSet()

		values := currentBeams.Values()
		//fmt.Println("Current beams at row", row, ":", values)

		for j := 0; j < len(values); j++ {
			switch diagram[row][values[j]] {
			case emptySpace:
				nextBeams.Add(values[j])
			case splitter:
				if currentBeams.Has(values[j]) {
					beamSplits++
					nextBeams.Add(values[j] - 1)
					nextBeams.Add(values[j] + 1)
				}
			default:
				fmt.Println("Unexpected character in diagram")
			}

		}

		currentBeams = nextBeams
		nextBeams = nil
	}

	return
}

func SolveDiagramPartTwo(diagram [][]rune) (splitTimelines int) {

	currentBeam := 0
	// Not solved yet
	for i := 0; i < len(diagram[0]); i++ {
		if diagram[0][i] == manifoldEntrance {
			currentBeam = i
			break
		}
	}
	fmt.Println("Current beam at start: ", currentBeam)
	memo := make(map[State]int)
	splitTimelines = findTimelines(diagram, currentBeam, 1, memo)

	fmt.Println("SplitTimelines: ", splitTimelines)
	return

}

type State struct {
	row int
	col int
}

func findTimelines(
	diagram [][]rune,
	currentBeam int,
	row int,
	memo map[State]int,
) int {

	if row >= len(diagram) {
		return 1
	}

	if currentBeam < 0 || currentBeam >= len(diagram[row]) {
		return 0
	}

	state := State{row, currentBeam}
	if val, ok := memo[state]; ok {
		return val
	}

	timelines := 0

	switch diagram[row][currentBeam] {
	case emptySpace:
		timelines = findTimelines(diagram, currentBeam, row+1, memo)

	case splitter:
		if currentBeam-1 >= 0 {
			timelines += findTimelines(diagram, currentBeam-1, row+1, memo)
		}
		if currentBeam+1 < len(diagram[row]) {
			timelines += findTimelines(diagram, currentBeam+1, row+1, memo)
		}

	default:
		panic("unexpected character")
	}

	memo[state] = timelines
	return timelines
}
