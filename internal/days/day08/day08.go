package day08

import (
	"cmp"
	"fmt"
	"math"
	"slices"
	"strings"
)

type Day08 struct{}

type Box struct {
	x int
	y int
	z int
}

func parse(input string) []Box {
	lines := strings.Split(strings.Trim(input, "\n "), "\n")
	boxes := make([]Box, 0, len(lines))
	var (
		x int
		y int
		z int
	)

	for _, line := range lines {
		_, _ = fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)
		boxes = append(boxes, Box{x, y, z})
	}
	return boxes
}

func dist(b1, b2 Box) float64 {
	dx := b2.x - b1.x
	dy := b2.y - b1.y
	dz := b2.z - b1.z
	return math.Sqrt(float64(dx*dx + dy*dy + dz*dz))
}

func (d Day08) Part1(input string) (string, error) {
	boxes := parse(input)

	type Connection struct {
		src      int
		dst      int
		distance float64
	}
	connections := []Connection{}

	for i, b := range boxes {
		for j := i + 1; j < len(boxes); j++ {
			connections = append(connections, Connection{i, j, dist(b, boxes[j])})
		}
	}

	slices.SortFunc(connections, func(c1, c2 Connection) int {
		return cmp.Compare(c1.distance, c2.distance)
	})

	// unconnected := map[int]struct{}{}
	// for i := range len(boxes) {
	// 	unconnected[i] = struct{}{}
	// }

	g := map[int][]int{}
	var total int
	for _, c := range connections {
		if total == 10 {
			break
		}

		_, sok := g[c.src]
		_, dok := g[c.dst]
		if sok && dok {
			continue
		}

		total++
		g[c.src] = append(g[c.src], c.dst)
		g[c.dst] = append(g[c.dst], c.src)

		// delete(unconnected, c.dst)
		// delete(unconnected, c.src)
		// if len(unconnected) == 0 {
		// 	break
		// }
	}

	fmt.Println(g)

	return "", nil
}

func (d Day08) Part2(input string) (string, error) { return "", nil }
