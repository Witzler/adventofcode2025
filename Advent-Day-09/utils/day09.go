package utils

import "math"

type Coordinate struct {
	x int
	y int
}

type Distance struct {
	Min int
	Max int
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func FindLargestRectangleArea(input [][]int) (maxArea int) {
	if len(input) == 0 {
		return 0
	}

	rows := len(input)
	if len(input[0]) != 2 {
		panic("Expected 2 values per row in input")
	}

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			p1 := Coordinate{x: input[i][0], y: input[i][1]}
			p2 := Coordinate{x: input[j][0], y: input[j][1]}

			width := abs(p2.x-p1.x) + 1
			height := abs(p2.y-p1.y) + 1
			area := width * height

			if area > maxArea {
				maxArea = area
			}
		}
	}
	return
}

func FindLargestRectangleAreaPartTwo(input [][]int) (maxArea int) {
	if len(input) == 0 {
		return 0
	}

	redPoints := map[Coordinate]bool{}
	redPointsInX := map[int][]Coordinate{}
	redPointsInY := map[int][]Coordinate{}

	// Einlesen der Punkte
	for _, row := range input {
		p := Coordinate{x: row[0], y: row[1]}
		redPoints[p] = true
		redPointsInX[p.x] = append(redPointsInX[p.x], p)
		redPointsInY[p.y] = append(redPointsInY[p.y], p)
	}

	// Alle horizontalen und vertikalen Linien markieren
	markedPoints := map[Coordinate]bool{}
	markedPointsInY := map[int][]Coordinate{}

	for p1 := range redPoints {

		// Vertikale Linien
		for _, p2 := range redPointsInX[p1.x] {
			if p1 == p2 {
				continue
			}
			yMin := min(p1.y, p2.y)
			yMax := max(p1.y, p2.y)
			for y := yMin; y <= yMax; y++ {
				p := Coordinate{x: p1.x, y: y}
				markedPoints[p] = true
				markedPointsInY[y] = append(markedPointsInY[y], p)
			}
		}

		// Horizontale Linien
		for _, p2 := range redPointsInY[p1.y] {
			if p1 == p2 {
				continue
			}
			xMin := min(p1.x, p2.x)
			xMax := max(p1.x, p2.x)
			for x := xMin; x <= xMax; x++ {
				p := Coordinate{x: x, y: p1.y}
				markedPoints[p] = true
				markedPointsInY[p1.y] = append(markedPointsInY[p1.y], p)
			}
		}
	}

	// Für jede Y-Zeile min/max X bestimmen
	distances := map[int]Distance{}
	for y, points := range markedPointsInY {
		if len(points) < 2 {
			continue
		}
		minX := math.MaxInt
		maxX := math.MinInt
		for _, p := range points {
			if p.x < minX {
				minX = p.x
			}
			if p.x > maxX {
				maxX = p.x
			}
		}
		distances[y] = Distance{Min: minX, Max: maxX}
	}

	// Rechtecke prüfen
	for p1 := range redPoints {
		for p2 := range redPoints {
			if p1 == p2 {
				continue
			}

			xMin := min(p1.x, p2.x)
			xMax := max(p1.x, p2.x)
			yMin := min(p1.y, p2.y)
			yMax := max(p1.y, p2.y)

			valid := true
			for y := yMin; y <= yMax; y++ {
				d, ok := distances[y]
				if !ok || xMin < d.Min || xMax > d.Max {
					valid = false
					break
				}
			}

			if valid {
				area := (xMax - xMin + 1) * (yMax - yMin + 1)
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}

	return
}
