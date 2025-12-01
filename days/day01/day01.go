package day01

import "fmt"

type Day01 struct {
	s1 string
	s2 string
}

func (d *Day01) Title() string {
	return "Day 01 - xxx"
}

func (d *Day01) Part1() {
}

func (d *Day01) Part2() {
}

func (d *Day01) Solution1() string {
	return fmt.Sprintf("%s", d.s1)
}

func (d *Day01) Solution2() string {
	return fmt.Sprintf("%s", d.s2)
}
