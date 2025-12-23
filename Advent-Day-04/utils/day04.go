package utils

const MAXALLOWEDROLLS int = 4 // FOUR!!!!

func ForkliftAccess(grid [][]bool) (output int) {
	rows := len(grid)
	cols := len(grid[0])

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			adjacentRolls := 0

			if isRole(grid, r, c) {
				adjacentRolls = adjacentRolls - 1 // To not count itself
				// Check all eight directions
				for dr := -1; dr <= 1; dr++ {
					for dc := -1; dc <= 1; dc++ {
						if isRole(grid, r+dr, c+dc) {
							adjacentRolls++
						}
					}
				}
				if adjacentRolls < MAXALLOWEDROLLS {
					output++
				}
			}
		}
	}
	return //output
}

func ForkliftAccessPartTwo(grid [][]bool) (output int) {
	rows := len(grid)
	cols := len(grid[0])

	var changed bool = true

	for changed {
		changed = false
		for r := 0; r < rows; r++ {
			for c := 0; c < cols; c++ {
				adjacentRolls := 0

				if isRole(grid, r, c) {
					adjacentRolls = adjacentRolls - 1 // To not count itself
					// Check all eight directions
					for dr := -1; dr <= 1; dr++ {
						for dc := -1; dc <= 1; dc++ {
							if isRole(grid, r+dr, c+dc) {
								adjacentRolls++
							}
						}
					}
					if adjacentRolls < MAXALLOWEDROLLS {
						output++
						grid[r][c] = false
						changed = true
					}
				}
			}
		}
	}
	return //output
}

func isRole(grid [][]bool, row int, col int) bool {
	if row < 0 || row >= len(grid) {
		return false
	}
	if col < 0 || col >= len(grid[row]) {
		return false
	}
	return grid[row][col]
}
