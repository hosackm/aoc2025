package bench

import (
	"os"
	"strings"
	"testing"

	"github.com/hosackm/aoc2025/internal/days/day01"
	"github.com/hosackm/aoc2025/internal/days/day02"
	"github.com/hosackm/aoc2025/internal/days/day03"
	"github.com/hosackm/aoc2025/internal/days/day04"
	"github.com/hosackm/aoc2025/internal/days/day05"
	"github.com/hosackm/aoc2025/internal/runner"
)

func BenchmarkPackageRunners(b *testing.B) {
	runners := []struct {
		name     string
		r        runner.Day
		filename string
	}{
		{"Day01", &day01.Day01{}, "../../inputs/01.txt"},
		{"Day02", &day02.Day02{}, "../../inputs/02.txt"},
		{"Day03", &day03.Day03{}, "../../inputs/03.txt"},
		{"Day04", &day04.Day04{}, "../../inputs/04.txt"},
		{"Day05", &day05.Day05{}, "../../inputs/05.txt"},
	}

	input := "some input data"

	for _, tt := range runners {
		data, _ := os.ReadFile(tt.filename)
		input = strings.Trim(string(data), "\n")

		b.Run(tt.name+"PartOne", func(b *testing.B) {
			for b.Loop() {
				_, _ = tt.r.Part1(input)
			}
		})

		b.Run(tt.name+"PartTwo", func(b *testing.B) {
			for b.Loop() {
				_, _ = tt.r.Part2(input)
			}
		})
	}
}
