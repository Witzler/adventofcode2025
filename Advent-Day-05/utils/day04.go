package utils

type Range struct {
	Start int
	End   int
}

type DataBase struct {
	Ranges []Range
}

func (r Range) isInRange(value int) bool { // (r Range) means this function belongs to Range type (like a method), r.isInRange(4)
	return value >= r.Start && value <= r.End
}

func (db DataBase) isInDB(value int) bool {
	for _, r := range db.Ranges {
		if r.isInRange(value) {
			return true
		}
	}
	return false
}

func (db *DataBase) addRange(r Range) {
	db.Ranges = append(db.Ranges, r)
}
