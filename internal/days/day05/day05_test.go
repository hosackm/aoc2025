package day05

import "testing"

func TestParse(t *testing.T) {
	input := `3-5
10-14
16-20
12-18

1
5
8
11
17
32`
	r, a := parse(input)
	if len(r) != 4 {
		t.Errorf("got ranges of length %d, wanted %d", len(r), 4)
	}

	if len(a) != 6 {
		t.Errorf("got available of length %d, wanted %d", len(a), 6)
	}
}

func TestPart1(t *testing.T) {
	input := `3-5
10-14
16-20
12-18

1
5
8
11
17
32`
	want := "3"

	got, err := Day05{}.Part1(input)
	if err != nil {
		t.Errorf("part 1 returned an error")
	}

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := `3-5
10-14
16-20
12-18

1
5
8
11
17
32`
	want := "14"

	got, err := Day05{}.Part2(input)
	if err != nil {
		t.Errorf("part 2 returned an error")
	}

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
