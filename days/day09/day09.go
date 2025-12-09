package day09

import (
	"aoc2025/util"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Day09 struct {
	s1 int
	s2 int
}

func (d *Day09) Title() string {
	return "Day 09 - Movie Theater"
}

func (d *Day09) Part1() {
	// var input = util.ReadInput("inputs/day09-example.txt")
	var input = util.ReadInput("inputs/day09.txt")

	redTiles := []tile{}
	for _, line := range input {
		coordStr := strings.Split(line, ",")
		coords := []int{}
		for _, str := range coordStr {
			coord, err := strconv.Atoi(str)
			if err == nil {
				coords = append(coords, coord)
			}
		}
		redTiles = append(redTiles, tile{x: coords[0], y: coords[1]})
	}

	rectangles := []rectangle{}
	for i, tile1 := range redTiles {
		for _, tile2 := range redTiles[i + 1:] {
			area := getArea(tile1, tile2)
			rectangles = append(rectangles, rectangle{
				corner1: tile1,
				corner2: tile2,
				area: area,
			})
		}
	}

	slices.SortFunc(rectangles, func (a, b rectangle) int {
		return b.area - a.area
	})

	d.s1 = rectangles[0].area
}

func (d *Day09) Part2() {
	// var input = util.ReadInput("inputs/day09-example.txt")
	var input = util.ReadInput("inputs/day09.txt")

	redTiles := []tile{}
	xSet := make(map[int]bool)
	ySet := make(map[int]bool)
	for _, line := range input {
		coordStr := strings.Split(line, ",")
		x, _ := strconv.Atoi(coordStr[0])
		y, _ := strconv.Atoi(coordStr[1])
		redTiles = append(redTiles, tile{x, y})
		xSet[x] = true
		ySet[y] = true
	}

	borderTiles := make(map[tile]bool)
	for _, t := range redTiles {
		borderTiles[t] = true
	}

	xCoords := make([]int, 0, len(xSet))
	for x := range xSet {
		xCoords = append(xCoords, x)
	}
	slices.Sort(xCoords)
	yCoords := make([]int, 0, len(ySet))
	for y := range ySet {
		yCoords = append(yCoords, y)
	}
	slices.Sort(yCoords)

	xMap := make(map[int]int, len(xCoords))
	yMap := make(map[int]int, len(yCoords))
	for i, x := range xCoords {
		xMap[x] = i
	}
	for i, y := range yCoords {
		yMap[y] = i
	}

	compressedTiles := []tile{}
	for _, t := range redTiles {
		cx := xMap[t.x]
		cy := yMap[t.y]
		compressedTiles = append(compressedTiles, tile{cx, cy})
	}

	width := len(xCoords)
	height := len(yCoords)
	grid := make([][]bool, height)
	for i := range grid {
		grid[i] = make([]bool, width)
	}

	prev := compressedTiles[0]
	for _, t := range compressedTiles[1:] {
		if prev.x == t.x {
			start, end := prev.y, t.y
			if start > end {
				start, end = end, start
			}
			for y := start; y <= end; y++ {
				grid[y][t.x] = true
			}
		} else {
			start, end := prev.x, t.x
			if start > end {
				start, end = end, start
			}
			for x := start; x <= end; x++ {
				grid[t.y][x] = true
			}
		}
		prev = t
	}
	t := compressedTiles[0]
	if prev.x == t.x {
		start, end := prev.y, t.y
		if start > end {
			start, end = end, start
		}
		for y := start; y <= end; y++ {
			grid[y][t.x] = true
		}
	} else {
		start, end := prev.x, t.x
		if start > end {
			start, end = end, start
		}
		for x := start; x <= end; x++ {
			grid[t.y][x] = true
		}
	}

	fmt.Printf("main polygon done\n")

	for y := range height {
		inside := false
		for x := range width {
			if grid[y][x] {
				if x == 0 || !grid[y][x - 1] {
					inside = !inside
				}
			} else if inside {
				grid[y][x] = true
			}
		}
	}

	fmt.Printf("main polygon filled\n")

	for _, row := range grid {
		for _, col := range row {
			if col {
				fmt.Printf("# ")
			} else {
				fmt.Printf(". ")
			}
		}
		fmt.Printf("\n")
	}

	rectangles := []rectangle{}
	for i, tile1 := range redTiles {
		c1x, c1y := xMap[tile1.x], yMap[tile1.y]
		for _, tile2 := range redTiles[i + 1:] {
			c2x, c2y := xMap[tile2.x], yMap[tile2.y]
			area := getArea(tile1, tile2)
			rectangles = append(rectangles, rectangle{
				corner1: tile{c1x, c1y},
				corner2: tile{c2x, c2y},
				area: area,
			})
		}
	}

	slices.SortFunc(rectangles, func (a, b rectangle) int {
		return b.area - a.area
	})

	fmt.Printf("got to the rectangle tests\n")

	for _, rect := range rectangles {
    minX, maxX := rect.corner1.x, rect.corner2.x
    if minX > maxX {
			minX, maxX = maxX, minX
    }
    minY, maxY := rect.corner1.y, rect.corner2.y
    if minY > maxY {
			minY, maxY = maxY, minY
    }
    
    allInside := true
    for y := minY; y <= maxY && allInside; y++ {
			for x := minX; x <= maxX; x++ {
				if !grid[y][x] {
					allInside = false
					break
				}
			}
    }
    
    if allInside {
			d.s2 = rect.area
			break
    }
	}
}

func (d *Day09) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day09) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}

type tile struct {
	x int
	y int
}

type rectangle struct {
	corner1 tile
	corner2 tile
	area int
}

func getArea(c1, c2 tile) int {
	w := c1.x - c2.x
	if w < 0 {
		w *= -1
	}
	h := c1.y - c2.y
	if h < 0 {
		h *= -1
	}
	return (w + 1) * (h + 1) 
}
