package utils

type Range struct {
	Start int
	End   int
}

type DataBase struct {
	Ranges []Range
}

func (r Range) isInRange(value int) bool { // (r Range) means this function belongs to Range type (like a method), r.isInRange(4)
	return r.Start <= value && value <= r.End
}

func (db DataBase) isInDB(value int) bool {
	for _, r := range db.Ranges {
		if r.isInRange(value) {
			return true
		}
	}
	return false
}

func (db *DataBase) AddRange(r Range) {
	db.Ranges = append(db.Ranges, r)
}

func (r Range) CreateRange(x int, y int) Range {
	return Range{Start: x, End: y}
}

func (db *DataBase) sortRanges() {
	// Simple bubble sort implementation
	n := len(db.Ranges)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if db.Ranges[j].Start > db.Ranges[j+1].Start {
				db.Ranges[j], db.Ranges[j+1] = db.Ranges[j+1], db.Ranges[j]
			}
		}
	}

}
func (db *DataBase) mergeRanges() { // Merging overlapping ranges
	if len(db.Ranges) == 0 {
		return
	}
	merged := make([]Range, 0)
	current := db.Ranges[0]
	for i := 1; i < len(db.Ranges); i++ {
		if db.Ranges[i].Start <= current.End+1 { // Overlapping or contiguous ranges
			if db.Ranges[i].End > current.End {
				current.End = db.Ranges[i].End
			}
		} else {
			merged = append(merged, current)
			current = db.Ranges[i]
		}

	}
	merged = append(merged, current)
	db.Ranges = merged
}

func CountFreshIngredients(db DataBase, ingredients []int) (freshIngredients int) { //Part 1
	freshIngredients = 0
	for _, ingredientID := range ingredients {
		if db.isInDB(ingredientID) {
			freshIngredients++
		}
	}
	return //freshIngredients
}

func ConsideredIngredients(db DataBase) (consideredIngredientsIDs int) { //Part 2
	db.sortRanges()
	db.mergeRanges()
	for _, r := range db.Ranges {
		consideredIngredientsIDs += (r.End - r.Start + 1)

		//fmt.Printf("Considering range %d-%d: %d ingredients\n", r.Start, r.End, (r.End - r.Start + 1))
	}
	return
}
