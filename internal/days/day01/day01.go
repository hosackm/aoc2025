package day01

import (
	"fmt"
	"strings"
)

type Day01 struct{}

type command struct {
	direction byte
	amount    int
}

type Dial struct {
	count int
	zeros int
}

func NewDial(start int) *Dial {
	return &Dial{start, 0}
}

// Get returns the current value of the dial
func (d *Dial) Get() int {
	if d == nil {
		return 0
	}
	return d.count
}

// GetZeros returns the number of zero crossings that have ocurred
func (d *Dial) GetZeros() int {
	if d == nil {
		return 0
	}
	return d.zeros
}

// Incr increments the dial by x while also tracking the number of zero crossings in d.zeros
func (d *Dial) Incr(x int) {
	// track where we were before to see if we cross 0 again
	before := d.count

	// count all the complete spins that bring us back to the same point
	fullSpins := x / 100
	d.zeros += fullSpins

	// perform final spin
	d.count += x % 100
	d.count %= 100

	// check if the leftover spin also crosses 0
	if before != 0 && d.count < before || d.count == 0 {
		d.zeros++
	}
}

// Decr increments the dial by x while also tracking the number of zero crossings in d.zeros
func (d *Dial) Decr(x int) {
	// track where we were before to see if we cross 0 again
	before := d.count

	// count all the complete spins that bring us back to the same point
	fullSpins := x / 100
	d.zeros += fullSpins

	// perform final spin
	d.count -= x % 100
	if d.count < 0 {
		d.count = 100 + d.count // no negatives
	}
	d.count %= 100

	// check if the leftover spin also crosses 0
	if before != 0 && d.count > before || d.count == 0 {
		d.zeros++
	}
}

// getCommands parses the puzzle input into a slice of commands to act upon
func getCommands(input string) []command {
	var (
		ch  byte
		amt int
	)

	lines := strings.Split(input, "\n")
	commands := make([]command, 0, len(lines))

	for _, s := range lines {
		_, _ = fmt.Sscanf(s, "%c%d", &ch, &amt)
		commands = append(commands, command{ch, amt})
	}

	return commands
}

// Part1 solves part 1 of the puzzle
func (d Day01) Part1(input string) (string, error) {
	count := 0
	dial := NewDial(50)
	for _, c := range getCommands(input) {
		switch c.direction {
		case 'L':
			dial.Decr(c.amount)
		case 'R':
			dial.Incr(c.amount)
		}

		if dial.Get()%100 == 0 {
			count++
		}
	}

	return fmt.Sprintf("%d", count), nil
}

// Part2 solves part 2 of the puzzle
func (d Day01) Part2(input string) (string, error) {
	dial := NewDial(50)
	for _, c := range getCommands(input) {
		switch c.direction {
		case 'L':
			dial.Decr(c.amount)
		case 'R':
			dial.Incr(c.amount)
		}
	}

	return fmt.Sprintf("%d", dial.GetZeros()), nil
}
