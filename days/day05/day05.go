package day05

import (
	"aoc2025/util"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Day05 struct {
	s1 int
	s2 int
}

func (d *Day05) Title() string {
	return "Day 05 - Cafeteria"
}

func (d *Day05) Part1() {
	// var input = util.ReadInput("inputs/day05-example.txt")
	var input = util.ReadInput("inputs/day05.txt")

	separator := 0
	for idx, line := range input {
		if line == "" {
			separator = idx
			break;
		}
	}

	freshTotal := 0
	for i := separator + 1; i < len(input); i++ {
		ingId, err := strconv.Atoi(input[i])
		if err != nil {
			panic(err)
		}

		for j := 0; j < separator; j++ {
			line := input[j]
			ends := strings.Split(line, "-")
			start, err := strconv.Atoi(ends[0])
			if err != nil {
				panic(err)
			}
			end, err := strconv.Atoi(ends[1])
			if err != nil {
				panic(err)
			}

			if ingId >= start && ingId <= end {
				freshTotal++
				break;
			}
		}
	}

	d.s1 = freshTotal
}

func (d *Day05) Part2() {
	// var input = util.ReadInput("inputs/day05-example.txt")
	var input = util.ReadInput("inputs/day05.txt")

	// first just parsing all the ranges into structs
	ranges := []idRange{}
	for _, line := range input {
		if line == "" {
			break;
		}

		ends := strings.Split(line, "-")
		start, err := strconv.Atoi(ends[0])
		if err != nil {
			panic(err)
		}
		end, err := strconv.Atoi(ends[1])
		if err != nil {
			panic(err)
		}

		ranges = append(ranges, idRange{start: start, end: end})
	}

	// now we sort the ranges by their start values
	slices.SortFunc(ranges, func (a, b idRange) int {
		return a.start - b.start
	})

	// now we do the work...
	combinedRanges := []idRange{}
	var curr idRange
	for idx, rng := range ranges {
		// pull the first range as initial current range
		if idx == 0 {
			curr = rng
			continue
		}

		// check if this range's start is less than the current range's end
		if curr.end >= rng.start {
			if curr.end >= rng.end {
				// this range is contained in the current one, so we just keep moving
				continue
			} else {
				// this range starts in the current one, but it ends further
				// so we can extend the current range and keep moving!
				curr.end = rng.end
				continue
			}
		} else {
			// this range's start is greater than the current range's end,
			// meaning we've hit a gap and can finalize the current range.
			// we then set this range as the new current and carry on
			combinedRanges = append(combinedRanges, curr)
			curr = rng
		}
	}
	// make sure to add the final current range
	combinedRanges = append(combinedRanges, curr)

	total := 0
	for _, rng := range combinedRanges {
		total += rng.end - rng.start + 1
	}

	d.s2 = total
}

func (d *Day05) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day05) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}

type idRange struct {
	start int
	end int
}
