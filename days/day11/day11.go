package day11

import (
	"aoc2025/util"
	"fmt"
	"strings"
)

type Day11 struct {
	s1 int
	s2 int
}

func (d *Day11) Title() string {
	return "Day 11 - Reactor"
}

func (d *Day11) Part1() {
	// var input = util.ReadInput("inputs/day11-example.txt")
	var input = util.ReadInput("inputs/day11.txt")

	devices := make(map[string]*device, len(input))
	devices["out"] = &device{id: "out"}
	for _, line := range input {
		id := strings.Split(line, ":")[0]
		devices[id] = &device{id: id}
	}

	for _, line := range input {
		parts := strings.Split(line, ":")
		conns := strings.Fields(parts[1])
		dev := devices[parts[0]]
		for _, conn := range conns {
			connDev := devices[conn]
			dev.connections = append(dev.connections, connDev)
		}
	}

	paths := 0
	queue := [][]string{{"you"}}

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		curr := devices[path[len(path) - 1]]

		if curr.id == "out" {
			paths++
		}

		for _, conn := range curr.connections {
			newPath := append([]string(nil), path...)
			newPath = append(newPath, conn.id)
			queue = append(queue, newPath)
		}
	}

	d.s1 = paths
}

func (d *Day11) Part2() {
	// var input = util.ReadInput("inputs/day11-example2.txt")
	var input = util.ReadInput("inputs/day11.txt")

	devices := make(map[string]*device, len(input))
	devices["out"] = &device{id: "out"}
	for _, line := range input {
		id := strings.Split(line, ":")[0]
		devices[id] = &device{id: id}
	}

	for _, line := range input {
		parts := strings.Split(line, ":")
		conns := strings.Fields(parts[1])
		dev := devices[parts[0]]
		for _, conn := range conns {
			connDev := devices[conn]
			dev.connections = append(dev.connections, connDev)
		}
	}
}

func (d *Day11) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day11) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}

type device struct {
	id string
	connections []*device
}

func dfsCount (curr, target *device, visited map[string]bool, count *int) {
	if curr.id == target.id {
		*count++
		return
	}

	visited[curr.id] = true
	for _, conn := range curr.connections {
		if !visited[conn.id] {
			dfsCount(conn, target, visited, count)
		}
	}
	visited[curr.id] = false
}
