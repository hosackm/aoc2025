package runner

import (
	"fmt"
)

type Day interface {
	Part1(input string) (string, error)
	Part2(input string) (string, error)
}

type DayRunner struct {
	days map[int]Day
}

func NewDayRunner() *DayRunner {
	return &DayRunner{map[int]Day{}}
}

func (r *DayRunner) Register(day int, impl Day) error {
	if day < 1 || day > 25 {
		return fmt.Errorf("day must be between 1 and 25")
	}

	r.days[day] = impl
	return nil
}

func (r *DayRunner) Get(day int) (Day, bool) {
	d, ok := r.days[day]
	return d, ok
}
