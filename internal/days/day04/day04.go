package day04

import (
	"fmt"
	"strings"
)

type Day04 struct{}

func parse(s string) ([]string, error) {
	return strings.Split(s, "\n"), nil
}

type Grid struct {
	data []string
	rows int
	cols int
}

func NewGrid(s []string) *Grid {
	g := Grid{
		data: s,
		rows: len(s),
		cols: len(s[0]),
	}

	return &g
}

func (g Grid) At(x, y int) (byte, error) {
	if x < 0 || x >= g.cols || y < 0 || y >= g.rows {
		return '0', fmt.Errorf("x and y must be within 0 and numCols or 0 and numRows respectively")
	}

	return g.data[y][x], nil
}

func (g *Grid) Clear(x, y int) error {
	if x < 0 || x >= g.cols || y < 0 || y >= g.rows {
		return fmt.Errorf("x and y must be within 0 and numCols or 0 and numRows respectively")
	}

	var b strings.Builder
	for i, ch := range g.data[y] {
		if i == x {
			b.WriteByte('x')
		} else {
			b.WriteByte(byte(ch))
		}
	}
	g.data[y] = b.String()

	return nil
}

func (g Grid) GetRolls() [][2]int {
	rolls := [][2]int{}

	for j := range len(g.data) {
		for i := range len(g.data[0]) {
			if g.data[j][i] == '@' {
				rolls = append(rolls, [2]int{i, j})
			}
		}
	}

	return rolls
}

func (g Grid) Adjacent(x, y int) ([]byte, error) {
	if x < 0 || x >= g.cols || y < 0 || y >= g.rows {
		return nil, fmt.Errorf("x and y must be within 0 and numCols or 0 and numRows respectively")
	}

	type delta struct {
		dy int
		dx int
	}
	adjacents := []delta{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
	result := make([]byte, 0, 8)
	for _, d := range adjacents {
		x2, y2 := x+d.dx, y+d.dy
		ch, err := g.At(x2, y2)
		if err != nil {
			continue
		}

		result = append(result, ch)
	}
	return result, nil
}

func (d Day04) Part1(input string) (string, error) {
	i, err := parse(input)
	if err != nil {
		return "", err
	}

	g := NewGrid(i)
	var total int
	for _, s := range g.GetRolls() {
		adjs, err := g.Adjacent(s[0], s[1])
		if err != nil {
			return "", err
		}

		if strings.Count(string(adjs), "@") < 4 {
			total++
		}
	}

	return fmt.Sprintf("%d", total), nil
}

func (d Day04) Part2(input string) (string, error) {
	i, err := parse(input)
	if err != nil {
		return "", err
	}

	g := NewGrid(i)
	var total int
	for {
		var roundTotal int
		for _, s := range g.GetRolls() {
			adjs, err := g.Adjacent(s[0], s[1])
			if err != nil {
				return "", err
			}

			if strings.Count(string(adjs), "@") < 4 {
				g.Clear(s[0], s[1])
				roundTotal++
			}
		}

		if roundTotal == 0 {
			break
		}
		total += roundTotal
	}

	return fmt.Sprintf("%d", total), nil
}
