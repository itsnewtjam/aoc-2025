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

	// first populate a map of devices
	devices := make(map[string]*device, len(input))
	devices["out"] = &device{id: "out"}
	for _, line := range input {
		id := strings.Split(line, ":")[0]
		devices[id] = &device{id: id}
	}

	// now that they all exist, we run back through and add the connections
	for _, line := range input {
		parts := strings.Split(line, ":")
		conns := strings.Fields(parts[1])
		dev := devices[parts[0]]
		for _, conn := range conns {
			connDev := devices[conn]
			dev.connections = append(dev.connections, connDev)
		}
	}

	// now just bfs
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

	// here we just dfs for each piece of the path, and then
	// combine them for the final total
	svr := devices["svr"]
	out := devices["out"]
	dac := devices["dac"]
	fft := devices["fft"]

	pathsSRVtoDAC := dfs(svr, dac, make(map[string]int))
	pathsSRVtoFFT := dfs(svr, fft, make(map[string]int))
	pathsDACtoFFT := dfs(dac, fft, make(map[string]int))
	pathsFFTtoDAC := dfs(fft, dac, make(map[string]int))
	pathsDACtoOUT := dfs(dac, out, make(map[string]int))
	pathsFFTtoOUT := dfs(fft, out, make(map[string]int))

	total := 0
	total += pathsSRVtoDAC * pathsDACtoFFT * pathsFFTtoOUT
	total += pathsSRVtoFFT * pathsFFTtoDAC * pathsDACtoOUT

	d.s2 = total
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

func dfs(curr, target *device, memo map[string]int) int {
	if curr.id == target.id {
		return 1
	}

	if cached, exists := memo[curr.id]; exists {
		return cached
	}

	count := 0
	for _, conn := range curr.connections {
		count += dfs(conn, target, memo)
	}
	memo[curr.id] = count
	return count
}
