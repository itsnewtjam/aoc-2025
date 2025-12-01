# Advent of Code 2025

My solutions for [Advent of Code 2025](https://adventofcode.com/2025) using Go.

## Project Structure

```
aoc-2025/
├── days/
│   ├── dayXX/         # Daily solutions
│   │   └── dayXX.go   # Solution runner
│   └── day.go         # Solution runner
├── inputs/            # Puzzle inputs
├── util/              # Util functions
└── main.go            # Entry
```

## Setup

1. Create a new solution file:
```bash
mkdir days/day02
cp days/day01/day01.go days/day02/day02.go
```

2. Add the day to the map in `main.go`:
```diff
  var day_map = map[string](func() days.Day) {
    "01": func() days.Day { d := day01.Day01{}; return &d },
+   "02": func() days.Day { d := day02.Day02{}; return &d },
  }
```

3. Add your puzzle input to `inputs/day02.txt`

## Commands

- Run today's solution:
```bash
go run main.go
```

- Run a specific day:
```bash
go run main.go 1
```

## Progress

⭐ Total stars: 2/24

- [x] Day 1
- [ ] Day 2
- [ ] Day 3
- [ ] Day 4
- [ ] Day 5
- [ ] Day 6
- [ ] Day 7
- [ ] Day 8
- [ ] Day 9
- [ ] Day 10
- [ ] Day 11
- [ ] Day 12
