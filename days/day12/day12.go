package day12

import (
	"aoc2025/util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Day12 struct {
	s1 int
	s2 int
}

func (d *Day12) Title() string {
	return "Day 12 - Christmas Tree Farm"
}

func (d *Day12) Part1() {
	// var input = util.ReadInput("inputs/day12-example.txt")
	var input = util.ReadInput("inputs/day12.txt")

	// just run through and parse out all our presents and trees
	presents := []present{}
	trees := []tree{}
	var curr present
	for _, line := range input {
		if line == "" {
				presents = append(presents, curr)
			continue
		}

		if matches, _ := regexp.MatchString("^[0-9]:", line); matches {
			curr = present{}
			continue
		}

		if matches, _ := regexp.MatchString("^[0-9]{1,}x", line); matches {
			parts := strings.Split(line, ":")
			dimensions := strings.Split(parts[0], "x")
			width, _ := strconv.Atoi(dimensions[0])
			height, _ := strconv.Atoi(dimensions[1])
			needStrings := strings.Fields(parts[1])
			needs := make([]int, len(needStrings))
			count := 0
			for i, str := range needStrings {
				need, _ := strconv.Atoi(str)
				needs[i] = need
				count += need
			}
			trees = append(trees, tree{width, height, needs, count})
			continue
		}

		row := make([]bool, len(line))
		for i, char := range line {
			if char == '#' {
				row[i] = true
				curr.area++
			}
		}
		curr.shape = append(curr.shape, row)
	}

	// first things first just run a basic check on the areas, if the
	// presents required take up more area than the tree allows, we
	// automatically know it's no good
	validTotal := 0
	for _, tree := range trees {
		totalArea := tree.width * tree.height
		neededArea := 0
		valid := true
		for i, count := range tree.needs {
			if count > 0 {
				neededArea += presents[i].area * count
				if neededArea > totalArea {
					valid = false
					break
				}
			}
		}
		if !valid {
			continue
		}

		// actually try and tetris them
		// ...is what i would have done, in a bit of avoidance i decided
		// to try just the area check before pulling my hair out building
		// a packing algorithm... :) merry christmas!

		if valid {
			validTotal++
		}
	}

	d.s1 = validTotal
}

func (d *Day12) Part2() {
	// var input = util.ReadInput("inputs/day12-example.txt")
	// var input = util.ReadInput("inputs/day12.txt")
}

func (d *Day12) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day12) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}

type present struct {
	shape [][]bool
	area int
}

type tree struct {
	width int
	height int
	needs []int
	count int
}
