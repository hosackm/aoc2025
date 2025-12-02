package day02

import (
	"strconv"
	"strings"
)

type Day02 struct{}

type Range struct {
	Start int
	End   int
}

func getRanges(input string) []Range {
	r := make([]Range, 0)
	for pair := range strings.SplitSeq(input, ",") {
		nums := strings.Split(pair, "-")
		n1, _ := strconv.Atoi(nums[0])
		n2, _ := strconv.Atoi(nums[1])
		r = append(r, Range{n1, n2})
	}
	return r
}

func (d Day02) Part1(input string) (string, error) { return "", nil }

func (d Day02) Part2(input string) (string, error) { return "", nil }
