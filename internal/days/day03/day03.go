package day03

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	JoltageLengthPartOne = 2
	JoltageLengthPartTwo = 12
)

type Day03 struct{}

func parseBatteries(s string) ([][]int, error) {
	batteries := [][]int{}
	for line := range strings.SplitSeq(s, "\n") {
		b := []int{}
		for _, battery := range line {
			n, err := strconv.Atoi(string(battery))
			if err != nil {
				return nil, err
			}
			b = append(b, n)
		}
		batteries = append(batteries, b)
	}
	return batteries, nil
}

func (d Day03) Part1(input string) (string, error) {
	batteries, err := parseBatteries(input)
	if err != nil {
		return "", err
	}

	var total int
	for _, b := range batteries {
		total += searchBatteries(b, JoltageLengthPartOne)
	}

	return fmt.Sprintf("%d", total), nil
}

func (d Day03) Part2(input string) (string, error) {
	batteries, err := parseBatteries(input)
	if err != nil {
		return "", err
	}

	var total int
	for _, b := range batteries {
		total += searchBatteries(b, JoltageLengthPartTwo)
	}

	return fmt.Sprintf("%d", total), nil
}

// searchBatteries iterates through the integers in batteries and
// creates a n-digit number by greedily choosing the maximum number
// to the right of the last number.
//
// For example:
//
//	batteries = [9,8,7,6,5,4,3,2,1,1,1,1] n = 2 -> 98
//	batteries = [2,3,4,2,3,4,2,3,4,2,3,4,2,7,8] n = 2 -> 78
func searchBatteries(batteries []int, n int) int {
	// keep track of n integer pairs, [0] is the maximum and [1]
	// is the index where the maximum was found
	maxs := make([][2]int, n)

	var total int
	for i := range n {
		var start int
		if i > 0 {
			start = maxs[i-1][1] + 1
		}

		// end early enough that we have numbers left to choose from
		end := len(batteries) - (n - i - 1)

		for j := start; j < end; j++ {
			if batteries[j] > maxs[i][0] {
				maxs[i][0] = batteries[j]
				maxs[i][1] = j
			}
		}

	}

	for i := range n {
		total *= 10
		total += maxs[i][0]
	}
	return total
}
