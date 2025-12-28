package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Point struct {
	x, y, z int
}

type Edge struct {
	i, j int
	dist int64
}

/* ---------- Union-Find ---------- */

type DSU struct {
	parent []int
	size   []int
	sets   int
}

func NewDSU(n int) *DSU {
	p := make([]int, n)
	s := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i
		s[i] = 1
	}
	return &DSU{p, s, n}
}

func (d *DSU) Find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.Find(d.parent[x])
	}
	return d.parent[x]
}

func (d *DSU) Union(a, b int) bool {
	ra, rb := d.Find(a), d.Find(b)
	if ra == rb {
		return false
	}
	if d.size[ra] < d.size[rb] {
		ra, rb = rb, ra
	}
	d.parent[rb] = ra
	d.size[ra] += d.size[rb]
	d.sets--
	return true
}

/* ---------- Main ---------- */

func main() {
	points := readInput()
	if len(points) == 0 {
		panic("no input points read")
	}

	n := len(points)

	edges := make([]Edge, 0, n*(n-1)/2)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dx := int64(points[i].x - points[j].x)
			dy := int64(points[i].y - points[j].y)
			dz := int64(points[i].z - points[j].z)
			dist := dx*dx + dy*dy + dz*dz
			edges = append(edges, Edge{i, j, dist})
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].dist < edges[j].dist
	})

	part1(points, edges)
	part2(points, edges)
}

/* ---------- Part 1 ---------- */

func part1(points []Point, edges []Edge) {
	n := len(points)
	ds := NewDSU(n)

	limit := 1000
	if limit > len(edges) {
		limit = len(edges)
	}

	for i := 0; i < limit; i++ {
		ds.Union(edges[i].i, edges[i].j)
	}

	componentSizes := map[int]int{}
	for i := 0; i < n; i++ {
		root := ds.Find(i)
		componentSizes[root]++
	}

	sizes := make([]int, 0, len(componentSizes))
	for _, s := range componentSizes {
		sizes = append(sizes, s)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))

	if len(sizes) < 3 {
		panic("not enough components for part 1")
	}

	result := sizes[0] * sizes[1] * sizes[2]
	fmt.Println("Part 1:", result)
}

/* ---------- Part 2 ---------- */

func part2(points []Point, edges []Edge) {
	n := len(points)
	ds := NewDSU(n)

	var last Edge
	for _, e := range edges {
		if ds.Union(e.i, e.j) {
			last = e
			if ds.sets == 1 {
				break
			}
		}
	}

	result := points[last.i].x * points[last.j].x
	fmt.Println("Part 2:", result)
}

/* ---------- Input ---------- */

func readInput() []Point {
	scanner := bufio.NewScanner(os.Stdin)
	points := []Point{}

	for scanner.Scan() {
		var x, y, z int
		_, err := fmt.Sscanf(scanner.Text(), "%d,%d,%d", &x, &y, &z)
		if err == nil {
			points = append(points, Point{x, y, z})
		}
	}

	return points
}

// 79056
// 4639477
