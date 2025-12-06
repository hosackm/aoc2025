package day06

import (
	"fmt"
	"strconv"
	"strings"
)

type Day06 struct{}

func parse(input string) ([][]int, []byte) {
	input = strings.Trim(input, "\n ")

	nums := make([][]int, 0)
	ops := make([]byte, 0)

	lines := strings.Split(input, "\n")
	for _, line := range lines[:len(lines)-1] {
		row := make([]int, 0)
		for snum := range strings.SplitSeq(line, " ") {
			if snum == "" {
				continue
			}

			num, _ := strconv.Atoi(snum)
			row = append(row, num)
		}
		nums = append(nums, row)
	}

	for op := range strings.SplitSeq(lines[len(lines)-1], " ") {
		if op == "" {
			continue
		}
		ops = append(ops, op[0])
	}
	return nums, ops
}

func (d Day06) Part1(input string) (string, error) {
	nums, ops := parse(input)

	var (
		rows int = len(nums)
		cols int = len(nums[0])
	)

	var grandTotal int
	for i := range cols {

		var op byte = ops[i]
		var total int = nums[0][i]
		for j := 1; j < rows; j++ {
			switch op {
			case '*':
				total *= nums[j][i]
			case '+':
				total += nums[j][i]
			}
		}
		grandTotal += total
	}

	return fmt.Sprintf("%d", grandTotal), nil
}

func mul(nums ...int) int {
	total := nums[0]
	for _, n := range nums[1:] {
		total *= n
	}
	return total
}

func sum(nums ...int) int {
	total := nums[0]
	for _, n := range nums[1:] {
		total += n
	}
	return total
}

func (d Day06) Part2(input string) (string, error) {
	lines := strings.Split(input, "\n")
	rows := len(lines)
	cols := len(lines[0])

	var total int

	nums := []int{}
	var op byte
	var sb strings.Builder
	for c := cols - 1; c >= 0; c-- {
		var all bool = true
		for r := range rows {
			if lines[r][c] != ' ' {
				all = false
			}
		}
		if all {
			var answer int
			switch op {
			case '*':
				answer = mul(nums...)
			case '+':
				answer = sum(nums...)
			}
			total += answer
			nums = []int{}
			continue
		}

		for r := range rows {
			if r == rows-1 {
				op = lines[r][c]
				break
			}
			if lines[r][c] != ' ' {
				sb.WriteByte(lines[r][c])
			}
		}

		snum := sb.String()
		sb.Reset()

		n, _ := strconv.Atoi(snum)

		nums = append(nums, n)
	}

	if len(nums) > 0 {
		var answer int
		switch op {
		case '*':
			answer = mul(nums...)
		case '+':
			total += sum(nums...)
		}
		total += answer
	}

	return fmt.Sprintf("%d", total), nil
}
