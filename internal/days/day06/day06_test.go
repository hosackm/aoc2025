package day06

import (
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	input := `123 328  51 64
45 64  387 23
6 98  215 314
*   +   *   +`
	nums, ops := parse(input)
	if len(nums) != 3 || len(nums[0]) != 4 {
		t.Errorf("expected nums of length %d,%d got %d,%d, %v", 3, 4, len(nums), len(nums[0]), nums)
	}
	if len(ops) != 4 {
		t.Errorf("expected ops of length %d got %d, %v", 4, len(ops), ops)
	}
}

func TestPart1(t *testing.T) {
	input := `123 328  51 64
45 64  387 23
6 98  215 314
*   +   *   +`
	got, err := Day06{}.Part1(input)
	if err != nil {
		t.Errorf("part 1 returned error - %s", err.Error())
	}

	want := "4277556"
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestPart2(t *testing.T) {
	inputs := []string{"123 328  51 64 ", " 45 64  387 23 ", "  6 98  215 314", "*   +   *   +  "}
	input := strings.Join(inputs, "\n")

	want := "3263827"
	got, err := Day06{}.Part2(input)
	if err != nil {
		t.Errorf("part 2 returned error - %s", err.Error())
	}

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
