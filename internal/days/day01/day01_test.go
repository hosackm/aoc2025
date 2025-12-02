package day01

import (
	"testing"

	"github.com/hosackm/aoc2025/internal/days/inputs"
)

func TestDay1Part1(t *testing.T) {
	day := Day01{}

	testInput := `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`
	want := "3"

	got, err := day.Part1(testInput)
	if err != nil {
		t.Errorf("part 1 returned error on valid input")
	}

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestDay1Part2(t *testing.T) {
	day := Day01{}

	testInput := `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`
	want := "6"

	got, err := day.Part2(testInput)
	if err != nil {
		t.Errorf("part 2 returned error on valid input")
	}

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestPart2Tip(t *testing.T) {
	// Be careful: if the dial were pointing at 50, a single rotation like R1000
	// would cause the dial to point at 0 ten times before returning back to 50!
	dial := NewDial(50)

	// goes around 10 times for 1000, then one more time for 50 + 51
	dial.Incr(1051)

	want := 11
	got := dial.GetZeros()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func BenchmarkPart1(b *testing.B) {
	data, err := inputs.FS.ReadFile("01.txt")
	if err != nil {
		b.Fatalf("couldn't read 01.txt for benchmark")
	}

	input := string(data)
	day1 := &Day01{}
	for b.Loop() {
		_, err := day1.Part1(input)
		if err != nil {
			b.Fatalf("Part1 returned an error")
		}
	}
}

func BenchmarkPart2(b *testing.B) {
	data, err := inputs.FS.ReadFile("01.txt")
	if err != nil {
		b.Fatalf("couldn't read 01.txt for benchmark")
	}

	input := string(data)
	day1 := &Day01{}
	for b.Loop() {
		_, err := day1.Part2(input)
		if err != nil {
			b.Fatalf("Part1 returned an error")
		}
	}
}
