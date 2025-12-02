package day02

import (
	"aoc2025/util"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Day02 struct {
	s1 int
	s2 int
}

func (d *Day02) Title() string {
	return "Day 02 - Gift Shop"
}

func (d *Day02) Part1() {
	// var input = util.ReadInput("inputs/day02-example.txt")
	var input = util.ReadInput("inputs/day02.txt")

	var invalids []int
	ranges := strings.Split(input[0], ",")
	for _, curr := range ranges {
		low, high := splitRange(curr)
		for i := low; i <= high; i++ {
			// id has odd digits so it can't satisfy repeat
			if intLen(i) % 2 != 0 {
				continue
			}

			iStr := strconv.Itoa(i)
			front := iStr[:len(iStr) / 2]
			back := iStr[len(iStr) / 2:]
			if front == back {
				invalids = append(invalids, i)
			}
		}
	}

	total := 0
	for _, val := range invalids {
		total += val
	}

	d.s1 = total
}

func (d *Day02) Part2() {
	// var input = util.ReadInput("inputs/day02-example.txt")
	var input = util.ReadInput("inputs/day02.txt")

	var invalids []int
	ranges := strings.Split(input[0], ",")
	for _, curr := range ranges {
		low, high := splitRange(curr)
		for i := low; i <= high; i++ {
			iStr := strconv.Itoa(i)

			// split the value into substrings of length 1 up to half the value length
			halfLen := len(iStr) / 2
			for j := halfLen; j > 0; j-- {
				subs := getSubstrings(iStr, j)
				// remove concurrent duplicate substrings
				// eg {12, 12, 13} becomes {12, 13}
				cleanUniq(&subs)

				// if the deduped slice only has one element,
				// the value was only made up of that sequence
				if len(subs) == 1 {
					invalids = append(invalids, i)
					break
				}
			}
		}
	}

	total := 0
	for _, val := range invalids {
		total += val
	}

	d.s2 = total
}

func (d *Day02) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day02) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}

func splitRange(rng string) (int, int) {
	ends := strings.Split(rng, "-")
	low, err := strconv.Atoi(ends[0])
	if err != nil {
		panic(err)
	}
	high, err := strconv.Atoi(ends[1])
	if err != nil {
		panic(err)
	}

	return low, high
}

func intLen(i int) int {
	if i == 0 {
		return 1
	}
	count := 0
	for i != 0 {
		i /= 10
		count++
	}
	return count
}

func getSubstrings(s string, size int) []string {
	var substrings []string
	for i := 0; i < len(s); i += size {
		if i + size <= len(s) {
			substrings = append(substrings, s[i:i + size])
		} else {
			substrings = append(substrings, s[i:])
		}
	}

	return substrings
}

func cleanUniq(s *[]string) {
	slices.Compact(*s)
	var newS []string
	for _, val := range *s {
		if val != "" {
			newS = append(newS, val)
		}
	}
	*s = newS
}
