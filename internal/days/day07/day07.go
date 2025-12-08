package day07

import (
	"fmt"
	"strings"
)

type Day07 struct{}

type point struct {
	row int
	col int
}

// parse reads the input and returns a map of points where splitters (ie. ^)
// exist as well as the column of the beam's origin (ie. S).
func parse(input string) (map[point]struct{}, int, int) {
	var origin int
	splitters := make(map[point]struct{})
	lines := strings.Split(strings.Trim(input, "\n "), "\n")
	for row, line := range lines {
		for col, char := range line {
			switch char {
			case 'S':
				origin = col
			case '^':
				splitters[point{row, col}] = struct{}{}
			}
		}
	}

	return splitters, origin, len(lines)
}

func (d Day07) Part1(input string) (string, error) {
	splitters, origin, rows := parse(input)

	beams := map[int]struct{}{origin: {}}

	var total int
	for r := 1; r < rows; r++ {
		splits := map[int]struct{}{}
		for c := range beams {
			if _, ok := splitters[point{r, c}]; ok {
				splits[c] = struct{}{}
			}
		}

		for s := range splits {
			delete(beams, s)
			beams[s-1] = struct{}{}
			beams[s+1] = struct{}{}
		}

		total += len(splits)
	}

	return fmt.Sprintf("%d", total), nil
}

func (d Day07) Part2(input string) (string, error) {
	splitters, origin, rows := parse(input)

	// store the number of timelines at this beam point instead of
	// just beam present/missing.
	beams := map[int]int{origin: 1}

	for r := 1; r < rows; r++ {
		splits := map[int]struct{}{}
		for c := range beams {
			if _, ok := splitters[point{r, c}]; ok {
				splits[c] = struct{}{}
			}
		}

		for s := range splits {
			timelines := beams[s]
			beams[s] = 0

			// if s > 0 {
			beams[s-1] += timelines
			// }
			// if s < cols-1 {
			beams[s+1] += timelines
			// }
			beams[s] = 0
		}
	}

	var total int
	for _, v := range beams {
		total += v
	}
	return fmt.Sprintf("%d", total), nil
}
