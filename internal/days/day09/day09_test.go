package day09

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	input := `7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`
	want := []point{
		{7, 1},
		{11, 1},
		{11, 7},
		{9, 7},
		{9, 5},
		{2, 5},
		{2, 3},
		{7, 3},
	}
	got := parse(input)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestExample1(t *testing.T) {
	input := `7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`
	want := "50"
	got, err := Day09{}.Part1(input)
	if err != nil {
		t.Errorf("part 1 returned err - %s", err.Error())
	}

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestExample2(t *testing.T) {
	input := `7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`
	want := "50"
	got, err := Day09{}.Part2(input)
	if err != nil {
		t.Errorf("part 2 returned err - %s", err.Error())
	}

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
