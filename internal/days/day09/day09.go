package day09

import (
	"fmt"
	"strings"
)

// const maxGrid = 98477
const maxGrid = 12

type Day09 struct{}

type point struct {
	x int
	y int
}

func parse(input string) []point {
	lines := strings.Split(strings.Trim(input, "\n "), "\n")
	points := make([]point, len(lines))

	for i, line := range lines {
		fmt.Sscanf(line, "%d,%d", &points[i].x, &points[i].y)
	}
	return points
}

func area(p1, p2 point) int {
	dx := p1.x - p2.x
	dy := p1.y - p2.y
	if dx < 0 {
		dx *= -1
	}
	if dy < 0 {
		dy *= -1
	}

	return (dx + 1) * (dy + 1)
}

func (d Day09) Part1(input string) (string, error) {
	points := parse(input)
	var max int = -1
	for i, p := range points[:len(points)-1] {
		for j := i + 1; j < len(points); j++ {
			a := area(p, points[j])
			if a > max {
				max = a
			}
		}
	}
	return fmt.Sprintf("%d", max), nil
}

func (d Day09) Part2(input string) (string, error) {
	points := parse(input)

	// grid := [maxGrid][maxGrid]byte{}
	grid := make([][]byte, maxGrid)
	for i := range maxGrid {
		grid[i] = make([]byte, maxGrid)
	}
	for i := range maxGrid {
		for j := range maxGrid {
			grid[j][i] = '.'
		}
	}

	for i, pt := range points {
		grid[pt.y][pt.x] = '#'

		var next point
		if i == len(points)-1 {
			next = points[0]
		} else {
			next = points[i+1]
		}

		dx := next.x - pt.x
		dy := next.y - pt.y

		if dx != 0 {
			fmt.Printf("going from %d,%d\n", pt.x, next.x)
			if dx < 0 {
				dx = -1
			} else {
				dx = 1
			}
			x := pt.x + dx
			for x != next.x {
				grid[next.y][x] = 'X'
				x += dx
			}
		}
		if dy != 0 {
			fmt.Printf("going from %d,%d\n", pt.y, next.y)
			if dy < 0 {
				dy = -1
			} else {
				dy = 1
			}
			y := pt.y + dy
			for y != next.y {
				grid[y][next.x] = 'X'
				y += dy
			}
		}

		grid[next.y][next.x] = '#'
	}

	for _, line := range grid {
		fmt.Println(string(line))
	}

	return "", nil
}
