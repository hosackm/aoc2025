package day04

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	input := `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`

	want := []string{
		"..@@.@@@@.",
		"@@@.@.@.@@",
		"@@@@@.@.@@",
		"@.@@@@..@.",
		"@@.@@@@.@@",
		".@@@@@@@.@",
		".@.@.@.@@@",
		"@.@@@.@@@@",
		".@@@@@@@@.",
		"@.@.@@@.@.",
	}

	got, err := parse(input)
	if err != nil {
		t.Fatalf("parse returned err - %s", err.Error())
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestClear(t *testing.T) {
	input := []string{
		"..@@.@@@@.",
		"@@@.@.@.@@",
		"@@@@@.@.@@",
		"@.@@@@..@.",
		"@@.@@@@.@@",
		".@@@@@@@.@",
		".@.@.@.@@@",
		"@.@@@.@@@@",
		".@@@@@@@@.",
		"@.@.@@@.@.",
	}

	g := NewGrid(input)

	want := byte('x')
	for _, pt := range [][2]int{{0, 0}, {2, 0}, {5, 5}} {
		x, y := pt[0], pt[1]
		err := g.Clear(x, y)
		if err != nil {
			t.Fatalf("clear returned error - %s", err.Error())
		}

		got, err := g.At(x, y)
		if got != want {
			t.Errorf("got %c want %c", got, want)
		}
	}
}

func TestPartOne(t *testing.T) {
	input := `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`
	want := "13"

	d := Day04{}
	got, err := d.Part1(input)
	if err != nil {
		t.Fatalf("parse returned err - %s", err.Error())
	}

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestPartTwo(t *testing.T) {
	input := `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`
	want := "43"

	d := Day04{}
	got, err := d.Part2(input)
	if err != nil {
		t.Fatalf("parse returned err - %s", err.Error())
	}

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
