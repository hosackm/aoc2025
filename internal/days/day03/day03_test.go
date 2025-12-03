package day03

import (
	"testing"
)

func TestParseInput(t *testing.T) {
	input := `987654321111111
811111111111119
234234234234278
818181911112111`
	batteries, err := parseBatteries(input)
	if err != nil {
		t.Fatalf("returned error - %s", err.Error())
	}

	if len(batteries) != 4 {
		t.Errorf("got %d battery rows, wanted 4", len(batteries))
	}

	for _, b := range batteries {
		if len(b) != 15 {
			t.Errorf("got %d batteries in a row, wanted 15", len(b))
		}
	}
}

func TestDayExamplePart1(t *testing.T) {
	input := `987654321111111
811111111111119
234234234234278
818181911112111`
	want := "357"

	d := Day03{}
	got, err := d.Part1(input)
	if err != nil {
		t.Fatalf("day3 part1 example 1 returned an error: %s", err.Error())
	}

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestDayExamplePart2(t *testing.T) {
	input := `987654321111111
811111111111119
234234234234278
818181911112111`
	want := "3121910778619"

	d := Day03{}
	got, err := d.Part2(input)
	if err != nil {
		t.Fatalf("day3 part1 example 1 returned an error: %s", err.Error())
	}

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
