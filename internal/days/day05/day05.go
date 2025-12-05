package day05

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Day05 struct{}

type Range struct {
	start int
	end   int
}

func parse(input string) ([]Range, []int) {
	lines := strings.Split(input, "\n")

	var (
		blank int
		start int
		end   int
	)

	// parse ranges first
	ranges := make([]Range, 0)
	for i, ln := range lines {
		if ln == "" {
			blank = i
			break
		}
		fmt.Sscanf(ln, "%d-%d", &start, &end)
		ranges = append(ranges, Range{start, end})
	}

	// parse available
	numAvailable := len(lines) - blank - 1
	available := make([]int, 0, numAvailable)
	for i := blank + 1; i < len(lines); i++ {
		n, _ := strconv.Atoi(lines[i])
		available = append(available, n)
	}
	return ranges, available
}

func (d Day05) Part1(input string) (string, error) {
	rs, available := parse(input)

	var total int
	for _, a := range available {
		for _, r := range rs {
			if a >= r.start && a <= r.end {
				total++
				break
			}
		}
	}

	return fmt.Sprintf("%d", total), nil
}

func overlap(r1, r2 Range) bool {
	return r1.start <= r2.end && r1.end >= r2.start
}

func min(nums ...int) int {
	if len(nums) == 0 {
		return -1
	}

	m := nums[0]
	for _, x := range nums[1:] {
		if x < m {
			m = x
		}
	}
	return m
}

func max(nums ...int) int {
	if len(nums) == 0 {
		return -1
	}

	m := nums[0]
	for _, x := range nums[1:] {
		if x > m {
			m = x
		}
	}
	return m
}

func (d Day05) Part2(input string) (string, error) {
	rs, _ := parse(input)

	// sort them
	slices.SortFunc(rs, func(a, b Range) int {
		if a.start < b.start {
			return -1
		}
		if a.start > b.start {
			return 1
		}

		if a.end < b.end {
			return -1
		}
		if a.end > b.end {
			return 1
		}

		return 0
	})

	// merge all the intervals
	intervals := make([]Range, 0)
	for _, r := range rs {
		newIntervals := make([]Range, 0)
		var inserted bool
		for _, interval := range intervals {
			if overlap(r, interval) {
				r = Range{min(r.start, interval.start), max(r.end, interval.end)}
				continue
			}
			if r.end < interval.start {
				newIntervals = append(newIntervals, r)
				inserted = true
				break
			}

			newIntervals = append(newIntervals, interval)
		}

		if !inserted {
			newIntervals = append(newIntervals, r)
		}
		intervals = newIntervals
	}

	// add up all the merged intervals
	var total int
	for _, interval := range intervals {
		total += interval.end - interval.start + 1
	}

	return fmt.Sprintf("%d", total), nil
}
