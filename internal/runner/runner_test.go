package runner

import (
	"testing"
	"testing/quick"
)

type Impl struct{}

func (i *Impl) Part1(s string) (string, error) { return "", nil }
func (i *Impl) Part2(s string) (string, error) { return "", nil }

func TestRunnerRegisterIsAvailable(t *testing.T) {
	runner := NewDayRunner()

	// starts empty
	for n := range 25 {
		_, ok := runner.Get(n + 1)
		if ok {
			t.Errorf("runner return a day for %d wanted none", n+1)
		}
	}

	err := quick.Check(func(day int) bool {
		err := runner.Register(day, &Impl{})
		if day > 25 || day < 1 {
			return err != nil
		}

		_, ok := runner.Get(day)
		return ok
	}, nil)

	if err != nil {
		t.Fatal(err)
	}
}
