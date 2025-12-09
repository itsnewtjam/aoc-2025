package main

import (
	"fmt"
	"maps"
	"os"
	"slices"
	"sort"

	"aoc2025/days"
	"aoc2025/days/day01"
	"aoc2025/days/day02"
	"aoc2025/days/day03"
	"aoc2025/days/day04"
	"aoc2025/days/day05"
	"aoc2025/days/day06"
	"aoc2025/days/day07"
	"aoc2025/days/day08"
	"aoc2025/days/day09"
)

func main() {
	var day_map = map[string](func() days.Day) {
		"01": func() days.Day { d := day01.Day01{}; return &d },
		"02": func() days.Day { d := day02.Day02{}; return &d },
		"03": func() days.Day { d := day03.Day03{}; return &d },
		"04": func() days.Day { d := day04.Day04{}; return &d },
		"05": func() days.Day { d := day05.Day05{}; return &d },
		"06": func() days.Day { d := day06.Day06{}; return &d },
		"07": func() days.Day { d := day07.Day07{}; return &d },
		"08": func() days.Day { d := day08.Day08{}; return &d },
		"09": func() days.Day { d := day09.Day09{}; return &d },
	}

	var dayKey = ""
	if len(os.Args) > 1 {
		dayKey = os.Args[1]
	}

	if len(dayKey) == 0 {
		var keys = slices.Collect(maps.Keys(day_map))
		sort.Strings(keys)
		dayKey = keys[len(keys)-1]
	}

	if len(dayKey) == 1 {
		dayKey = fmt.Sprintf("0%s", dayKey)
	}

	var day = day_map[dayKey]
	if day != nil {
		days.Solve(day())
	} else {
		panic("Day not found")
	}
}
