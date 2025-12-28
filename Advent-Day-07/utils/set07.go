package utils

type IntSetInterface interface { // Idk
	Add(value int)
	Has(value int) bool
	Length() int
	Values() []int
	Remove(value int)
}

type IntSet struct {
	positionBeams map[int]struct{}
}

func NewIntSet() *IntSet {
	return &IntSet{positionBeams: make(map[int]struct{})}
}

func (s *IntSet) Add(value int) {
	s.positionBeams[value] = struct{}{}
}

func (s *IntSet) Remove(value int) {
	delete(s.positionBeams, value)
}

func (s *IntSet) Length() int {
	return len(s.positionBeams)
}
func (s *IntSet) Has(value int) bool {
	_, exists := s.positionBeams[value]
	return exists
}
func (s *IntSet) Values() (values []int) {
	values = make([]int, 0, len(s.positionBeams))
	for k := range s.positionBeams {
		values = append(values, k)
	}
	return
}
