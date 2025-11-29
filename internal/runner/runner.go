package runner

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

func (r *DayRunner) Register(day int, impl Day) {
	r.days[day] = impl
}

func (r *DayRunner) Get(day int) (Day, bool) {
	d, ok := r.days[day]
	return d, ok
}
