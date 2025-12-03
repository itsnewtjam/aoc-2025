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
)

func main() {
	var day_map = map[string](func() days.Day) {
		"01": func() days.Day { d := day01.Day01{}; return &d },
		"02": func() days.Day { d := day02.Day02{}; return &d },
		"03": func() days.Day { d := day03.Day03{}; return &d },
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
