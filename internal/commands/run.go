package commands

import (
	"fmt"

	"github.com/hosackm/aoc/internal/days/day01"
	"github.com/hosackm/aoc/internal/days/inputs"
	"github.com/hosackm/aoc/internal/runner"
)

func HandleRun(day int) error {
	if day < 1 || day > 25 {
		return fmt.Errorf("day must be between 1 through 25")
	}

	r := runner.NewDayRunner()
	dayImplementations := []runner.Day{
		day01.Day01{},
	}

	for _, d := range dayImplementations {
		if err := r.Register(1, d); err != nil {
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

	type PartFunc func(string) (string, error)
	for name, f := range map[string]PartFunc{"part 1": d.Part1, "part 2": d.Part2} {
		output, err := f(string(b))
		if err != nil {
			return err
		}
		fmt.Printf("[%s]: %s\n", name, output)
	}

	return nil
}
