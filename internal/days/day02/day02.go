package day02

import (
	"fmt"
	"strconv"
	"strings"

	// backtracking regexp with support for \1 required
	regexp "github.com/dlclark/regexp2"
)

type Day02 struct{}

type Range struct {
	Start int
	End   int
}

func getRanges(input string) []Range {
	r := make([]Range, 0)
	for pair := range strings.SplitSeq(strings.TrimRight(input, "\n"), ",") {
		nums := strings.Split(pair, "-")
		n1, _ := strconv.Atoi(nums[0])
		n2, _ := strconv.Atoi(nums[1])
		r = append(r, Range{n1, n2})
	}
	return r
}

func isRepeatingOriginal(n int) bool {
	s := fmt.Sprintf("%d", n)
	if len(s)%2 == 1 {
		return false
	}
	return s[:len(s)/2] == s[len(s)/2:]
}

func isRepeating(n int, re *regexp.Regexp) (bool, error) {
	s := fmt.Sprintf("%d", n)
	ok, err := re.MatchString(s)
	if err != nil {
		return false, err
	}

	return ok, nil
}

func (d Day02) Part1(input string) (string, error) {
	ranges := getRanges(input)

	re, err := regexp.Compile(`^(\d+)\1$`, 0)
	if err != nil {
		return "", err
	}

	sum := 0
	for _, r := range ranges {
		for i := r.Start; i <= r.End; i++ {
			// ok := isRepeatingOriginal(i)
			ok, err := isRepeating(i, re)
			if err != nil {
				return "", err
			}

			if ok {
				sum += i
			}
		}
	}

	return fmt.Sprintf("%d", sum), nil
}

func (d Day02) Part2(input string) (string, error) {
	ranges := getRanges(input)

	re, err := regexp.Compile(`^(\d+)\1+$`, 0)
	if err != nil {
		return "", err
	}

	sum := 0
	for _, r := range ranges {
		for i := r.Start; i <= r.End; i++ {
			ok, err := isRepeating(i, re)
			if err != nil {
				return "", err
			}

			if !ok {
				continue
			}

			sum += i
		}
	}

	return fmt.Sprintf("%d", sum), nil
}
