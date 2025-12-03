package day02

import "testing"

func TestDay02Parse(t *testing.T) {
	input := `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`
	want := []Range{
		{11, 22},
		{95, 115},
		{998, 1012},
		{1188511880, 1188511890},
		{222220, 222224},
		{1698522, 1698528},
		{446443, 446449},
		{38593856, 38593862},
		{565653, 565659},
		{824824821, 824824827},
		{2121212118, 2121212124},
	}

	got := getRanges(input)
	for i, v := range got {
		if v != want[i] {
			t.Errorf("got %d want %d", got, want)
		}
	}
}

func TestPart1Example(t *testing.T) {
	input := "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"
	want := "1227775554"

	day := Day02{}
	got, err := day.Part1(input)
	if err != nil {
		t.Fatalf("Part1() returned an error - %v", err)
	}

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestPart2Example(t *testing.T) {
	input := "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"
	want := "4174379265"

	day := Day02{}
	got, err := day.Part2(input)
	if err != nil {
		t.Fatalf("Part2() returned an error - %v", err)
	}

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
