package commands

import (
	"fmt"
	"os"

	"github.com/hosackm/aoc/internal/days/day01"
	"github.com/hosackm/aoc/internal/runner"
)

func HandleRun(day int) error {
	if day < 1 || day > 25 {
		return fmt.Errorf("day must be between 1 through 25")
	}

	r := runner.NewDayRunner()
	r.Register(1, day01.Day01{})

	d, ok := r.Get(day)
	if !ok {
		return fmt.Errorf("day %d doesn't exist yet", day)
	}

	inputFilename := fmt.Sprintf("inputs/%02d_input.txt", day)
	b, err := os.ReadFile(inputFilename)
	if err != nil {
		return err
	}

	input := string(b)

	type PartFunc func(string) (string, error)
	for name, f := range map[string]PartFunc{"part 1": d.Part1, "part 2": d.Part2} {
		output, err := f(input)
		if err != nil {
			return err
		}
		fmt.Printf("[%s]: %s\n", name, output)
	}

	return nil
}
