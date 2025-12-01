package days

import (
	"fmt"
)

type Day interface {
	Title() string
	Part1()
	Part2()
	Solution1() string
	Solution2() string
}

func Solve(d Day) {
	d.Part1()
	d.Part2()

	fmt.Printf("=== %s ===\n", d.Title())
	fmt.Printf("  Solution 1: %s\n", d.Solution1())
	fmt.Printf("  Solution 2: %s\n", d.Solution2())
}
