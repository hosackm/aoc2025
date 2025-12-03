package commands

import (
	"fmt"
	"strings"

	"github.com/hosackm/aoc2025/internal/days/day01"
	"github.com/hosackm/aoc2025/internal/days/day02"
	"github.com/hosackm/aoc2025/internal/days/day03"
	"github.com/hosackm/aoc2025/internal/days/inputs"
	"github.com/hosackm/aoc2025/internal/runner"
)

func HandleRun(day int) error {
	if day < 1 || day > 25 {
		return fmt.Errorf("day must be between 1 through 25")
	}

	r := runner.NewDayRunner()
	dayImplementations := map[int]runner.Day{
		1: day01.Day01{},
		2: day02.Day02{},
		3: day03.Day03{},
	}

	for num, day := range dayImplementations {
		if err := r.Register(num, day); err != nil {
			return err
		}
	}

	d, ok := r.Get(day)
	if !ok {
		return fmt.Errorf("day %d doesn't exist yet", day)
	}

	b, err := inputs.FS.ReadFile(fmt.Sprintf("%02d.txt", day))
	if err != nil {
		return err
	}

	input := strings.Trim(string(b), "\n")

	type PartFunc func(string) (string, error)
	for i, f := range []PartFunc{d.Part1, d.Part2} {
		output, err := f(input)
		if err != nil {
			return err
		}
		fmt.Printf("[part %d]: %s\n", i+1, output)
	}

	return nil
}
