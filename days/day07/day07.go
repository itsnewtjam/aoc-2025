package day07

import (
	"aoc2025/util"
	"fmt"
	"strings"
)

type Day07 struct {
	s1 int
	s2 int
}

func (d *Day07) Title() string {
	return "Day 07 - Laboratories"
}

func (d *Day07) Part1() {
	// var input = util.ReadInput("inputs/day07-example.txt")
	var input = util.ReadInput("inputs/day07.txt")

	width := len(input[0])
	tachyonColumns := make([]int, width)
	tachyonColumns[strings.Index(input[0], "S")] = 1
	splits := 0
	for _, line := range input[1:] {
		for i, col := range line {
			if tachyonColumns[i] == 0 {
				continue
			}
			if col == '^' {
				splits++
				tachyonColumns[i] = 0
				if i - 1 >= 0 {
					tachyonColumns[i - 1] = 1
				}
				if i + 1 < width {
					tachyonColumns[i + 1] = 1
				}
			}
		}
	}

	d.s1 = splits
}

func (d *Day07) Part2() {
	// var input = util.ReadInput("inputs/day07-example.txt")
	var input = util.ReadInput("inputs/day07.txt")

	// this is just part 1 but instead of just tracking *if* a column has 
	// a beam, we track the number of beams in a column at that point
	width := len(input[0])
	tachyonColumns := make([]int, width)
	tachyonColumns[strings.Index(input[0], "S")] = 1
	for _, line := range input[1:] {
		for i, col := range line {
			if tachyonColumns[i] == 0 {
				continue
			}
			if col == '^' {
				if i - 1 >= 0 {
					tachyonColumns[i - 1] += tachyonColumns[i]
				}
				if i + 1 < width {
					tachyonColumns[i + 1] += tachyonColumns[i]
				}
				tachyonColumns[i] = 0
			}
		}
	}

	total := 0
	for _, col := range tachyonColumns {
		total += col
	}

	d.s2 = total
}

func (d *Day07) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day07) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}
