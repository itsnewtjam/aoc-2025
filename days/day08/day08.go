package day08

import (
	"aoc2025/util"
	"fmt"
	"maps"
	"math"
	"slices"
	"strconv"
	"strings"
)

type Day08 struct {
	s1 int
	s2 int
}

func (d *Day08) Title() string {
	return "Day 08 - Playground"
}

func (d *Day08) Part1() {
	// var input = util.ReadInput("inputs/day08-example.txt")
	// var iter = 10
	var input = util.ReadInput("inputs/day08.txt")
	var iter = 1000

	// just parsing all of the boxes and creating the initial circuits
	boxes := []junctBox{}
	circuits := make(map[string]map[string]bool)
	for _, line := range input {
		circuits[line] = make(map[string]bool)
		circuits[line][line] = true

		coordsStr := strings.Split(line, ",")
		coords := []int{}
		for _, str := range coordsStr {
			coord, err := strconv.Atoi(str)
			if err == nil {
				coords = append(coords, coord)
			}
		}
		boxes = append(boxes, junctBox{id: line, x: int(coords[0]), y: coords[1], z: coords[2]})
	}

	// create a list of all the possible pairs
	pairs := []pair{}
	for i, box1 := range boxes {
		for _, box2 := range boxes[i + 1:] {
			dist := distance(box1, box2)
			pairs = append(pairs, pair{
				box1: box1,
				box2: box2,
				distance: dist,
			})
		}
	}

	// sort pairs by distance
	slices.SortFunc(pairs, func (a, b pair) (int) {
		if a.distance < b.distance {
			return -1
		} else if a.distance > b.distance {
			return 1
		} else {
			return 0
		}
	})

	for i := range iter {
		pair := pairs[i]
		// find which circuit each box is in
		b1Circuit := findCircuit(circuits, pair.box1.id)
		b2Circuit := findCircuit(circuits, pair.box2.id)

		if b1Circuit == b2Circuit {
			// they're in the same circuit, no changes to be made
			continue
		}
		// combine them by moving all the boxes from box2's circuit to box1's circuit
		for k := range circuits[b2Circuit] {
			circuits[b1Circuit][k] = true
		}
		// then delete box2's circuit since it's been combined
		delete(circuits, b2Circuit)
	}

	// sort circuits by size
	sortedCircuits := slices.Collect(maps.Values(circuits))
	slices.SortFunc(sortedCircuits, func (a, b map[string]bool) int {
		return len(b) - len(a)
	})

	d.s1 = len(sortedCircuits[0]) * len(sortedCircuits[1]) * len(sortedCircuits[2])
}

func (d *Day08) Part2() {
	// var input = util.ReadInput("inputs/day08-example.txt")
	var input = util.ReadInput("inputs/day08.txt")

	boxes := []junctBox{}
	circuits := make(map[string]map[string]bool)
	for _, line := range input {
		circuits[line] = make(map[string]bool)
		circuits[line][line] = true

		coordsStr := strings.Split(line, ",")
		coords := []int{}
		for _, str := range coordsStr {
			coord, err := strconv.Atoi(str)
			if err == nil {
				coords = append(coords, coord)
			}
		}
		boxes = append(boxes, junctBox{id: line, x: int(coords[0]), y: coords[1], z: coords[2]})
	}

	pairs := []pair{}
	for i, box1 := range boxes {
		for _, box2 := range boxes[i + 1:] {
			dist := distance(box1, box2)
			pairs = append(pairs, pair{
				box1: box1,
				box2: box2,
				distance: dist,
			})
		}
	}

	slices.SortFunc(pairs, func (a, b pair) (int) {
		if a.distance < b.distance {
			return -1
		} else if a.distance > b.distance {
			return 1
		} else {
			return 0
		}
	})

	var last pair
	for _, pair := range pairs {
		b1Circuit := findCircuit(circuits, pair.box1.id)
		b2Circuit := findCircuit(circuits, pair.box2.id)

		if b1Circuit == b2Circuit {
			continue
		}
		for k := range circuits[b2Circuit] {
			circuits[b1Circuit][k] = true
		}
		delete(circuits, b2Circuit)
		// the only difference from part 1:
		// instead of a set number of connections, we loop until we
		// only have one circuit, then pull the last pair we connected
		if len(circuits) == 1 {
			last = pair
			break
		}
	}

	d.s2 = last.box1.x * last.box2.x
}

func (d *Day08) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day08) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}

type pair struct {
	box1 junctBox
	box2 junctBox
	distance float64
}

type junctBox struct {
	id string
	x int
	y int
	z int
}

func distance(p1, p2 junctBox) (float64) {
	xSeg := math.Pow(float64(p1.x - p2.x), 2)
	ySeg := math.Pow(float64(p1.y - p2.y), 2)
	zSeg := math.Pow(float64(p1.z - p2.z), 2)
	return math.Sqrt(xSeg + ySeg + zSeg)
}

func findCircuit(m map[string]map[string]bool, s string) string {
	for k, v := range m {
		if v[s] {
			return k
		}
	}
	return ""
}
