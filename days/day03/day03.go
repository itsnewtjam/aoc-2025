package day03

import (
	"aoc2025/util"
	"fmt"
	"strconv"
)

type Day03 struct {
	s1 int
	s2 int
}

func (d *Day03) Title() string {
	return "Day 03 - Lobby"
}

func (d *Day03) Part1() {
	// var input = util.ReadInput("inputs/day03-example.txt")
	var input = util.ReadInput("inputs/day03.txt")

	total := 0

	for _, line := range input {
		var first, second byte
		for i := 0; i < len(line); i++ {
			if line[i] > first {
				// check if its the last battery, if it is we just set it as
				// the second because there wont be a second if we set it as first
				if i + 1 == len(line) {
					second = line[i]
				} else {
					first = line[i]
					second = 0
				}
			} else if line[i] > second {
				second = line[i]
			}
		}

		joltage, err := strconv.Atoi(string([]byte{first, second}))
		if err != nil {
			panic(err)
		}

		total += joltage
	}

	d.s1 = total
}

func (d *Day03) Part2() {
	// var input = util.ReadInput("inputs/day03-example.txt")
	var input = util.ReadInput("inputs/day03.txt")

	total := 0
	for _, line := range input {
		batteries := make([]byte, 12)

		for i := 0; i < len(line); i++ {
			reset := false
			// count how many more batteries we can possibly take from this bank
			remaining := min(12, len(line) - (i + 1))
			for idx, val := range batteries {
				// skip batteries that are too early in our list
				if idx + 1 < 12 - remaining {
					continue
				}
				if reset {
					// we set an earlier battery, so we reset this position
					batteries[idx] = '0'
				} else if line[i] > val {
					// the current battery is greater than the one we have at this
					// position, so set it to this position and flag to reset all
					// following batteries
					reset = true
					batteries[idx] = line[i]
				}
			}
		}

		joltage, err := strconv.Atoi(string(batteries))
		if err != nil {
			panic(err)
		}

		total += joltage
	}

	d.s2 = total
}

func (d *Day03) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day03) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}
