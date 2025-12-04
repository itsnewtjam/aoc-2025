package day04

import (
	"aoc2025/util"
	"fmt"
)

type Day04 struct {
	s1 int
	s2 int
}

func (d *Day04) Title() string {
	return "Day 04 - Printing Department"
}

func (d *Day04) Part1() {
	// var input = util.ReadInput("inputs/day04-example.txt")
	var input = util.ReadInput("inputs/day04.txt")

	accessible := 0
	for row, line := range input {
		for col, pos := range line {
			if pos == '.' {
				continue;
			}

			adj := 0
			for _, dir := range directions {
				if col + dir.x >= 0 &&
					col + dir.x < len(line) &&
					row + dir.y >= 0 &&
					row + dir.y < len(input) {

					checkLine := input[row + dir.y]
					checkPos := []rune(checkLine)[col + dir.x]
					if checkPos == '@' {
						adj++
						// just check if we hit four adjacent and eject early
						if adj == 4 {
							break;
						}
					}
				}
			}


			if adj < 4 {
				accessible++
			}
		}
	}

	d.s1 = accessible
}

func (d *Day04) Part2() {
	// var input = util.ReadInput("inputs/day04-example.txt")
	var input = util.ReadInput("inputs/day04.txt")

	total := 0
	removed := 0
	for {
		for row, line := range input {
			for col, pos := range line {
				if pos == '.' || pos == 'x' {
					continue;
				}

				adj := 0
				for _, dir := range directions {
					if col + dir.x >= 0 &&
						col + dir.x < len(line) &&
						row + dir.y >= 0 &&
						row + dir.y < len(input) {

						checkLine := input[row + dir.y]
						checkPos := []rune(checkLine)[col + dir.x]
						if checkPos == '@' {
							adj++
							if adj == 4 {
								break;
							}
						}
					}
				}

				if adj < 4 {
					// replace the roll and update the input
					newLine := []rune(line)
					newLine[col] = 'x'
					input[row] = string(newLine)
					// update the current loop variable to reflect the changes
					line = string(newLine)

					removed++
					total++
				}
			}
		}

		if removed == 0 {
			break
		} else {
			removed = 0
		}
	}

	d.s2 = total
}

func (d *Day04) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day04) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}

type direction struct {
	x int
	y int
}

var directions = []direction{
	{x: 0, y: -1},		// up
	{x: 0, y: 1}, 	// down
	{x: 1, y: 0},		// right
	{x: -1, y: 0},	// left
	{x: 1, y: -1}, 	// up-right
	{x: 1, y: 1}, 	// down-right
	{x: -1, y: 1},	// down-left
	{x: -1, y: -1},	// up-left
}
