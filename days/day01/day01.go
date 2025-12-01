package day01

import (
	"aoc2025/util"
	"fmt"
	"strconv"
)

type Day01 struct {
	s1 int
	s2 int
}

func (d *Day01) Title() string {
	return "Day 01 - Secret Entrance"
}

func (d *Day01) Part1() {
	// var input = util.ReadInput("inputs/day01-example.txt")
	var input = util.ReadInput("inputs/day01.txt")

	var pos = 50
	var zeroCount = 0
	for _, line := range input {
		dir := line[0]
		count, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}
		if dir == 'L' {
			count *= -1
		}

		pos = (pos + count) % 100
		if pos < 0 {
			pos += 100
		}
		if (pos == 0) {
			zeroCount++
		}
	}

	d.s1 = zeroCount
}

func (d *Day01) Part2() {
	// var input = util.ReadInput("inputs/day01-example.txt")
	var input = util.ReadInput("inputs/day01.txt")

	var pos = 50
	var zeroCount = 0
	for _, line := range input {
		dir := line[0]
		count, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}

		var flips = count / 100
		zeroCount += flips
		count %= 100
		if count < 0 {
			count *= -1
		}

		if dir == 'L' {
			count *= -1
		}

		var travel = pos + count
		var flipped = (travel < 0 || travel > 100) && pos > 0
		pos = travel % 100
		if pos < 0 {
			pos += 100
		}
		if (pos == 0 || flipped) {
			zeroCount++
		}
	}

	d.s2 = zeroCount
}

func (d *Day01) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day01) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}
