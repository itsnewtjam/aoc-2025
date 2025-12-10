package day09

import (
	"aoc2025/util"
	"container/list"
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

	// build out all the possible rectangles from the red tiles
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

	// then sort them by area to get the largest
	slices.SortFunc(rectangles, func (a, b rectangle) int {
		return b.area - a.area
	})

	d.s1 = rectangles[0].area
}

func (d *Day09) Part2() {
	// var input = util.ReadInput("inputs/day09-example.txt")
	var input = util.ReadInput("inputs/day09.txt")

	// parse through the reds, keeping track of the x and y of each
	// so we can compress the grid
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

	// get sorted slices of the coordinates...
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

	// ...so we can map them for the compressed grid
	xMap := make(map[int]int, len(xCoords))
	yMap := make(map[int]int, len(yCoords))
	for i, x := range xCoords {
		xMap[x] = i
	}
	for i, y := range yCoords {
		yMap[y] = i
	}

	// just convert all our tiles to compressed coordinates
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

	// here we're filling between the red tiles to get the border of the polygon
	prev := compressedTiles[0]
	for _, t := range compressedTiles[1:] {
		if prev.x == t.x {
			// if the x's are equal we're drawing a vertical line
			start, end := prev.y, t.y
			if start > end {
				start, end = end, start
			}
			for y := start; y <= end; y++ {
				grid[y][t.x] = true
			}
		} else {
			// horizontal line
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
	// make sure to connect the final tile to the first one to close the shape
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

	// find a suitable starting position to fill the polygon
	var startX, startY int
	for y := range height {
		for x := range width {
			if grid[y][x] && y - 1 >= 0 && !grid[y - 1][x] {
				startX, startY = x, y - 1
			}
		}
	}

	// flood fill the polygon
	queue := list.New()
	queue.PushBack(tile{startX, startY})

	for queue.Len() > 0 {
		front := queue.Remove(queue.Front()).(tile)
		x, y := front.x, front.y

		if x < 0 || x >= width || y < 0 || y >= height || grid[y][x] {
			continue
		}

		grid[y][x] = true

		queue.PushBack(tile{x, y - 1})
		queue.PushBack(tile{x, y + 1})
		queue.PushBack(tile{x - 1, y})
		queue.PushBack(tile{x + 1, y})
	}

	// build our rectangles and sort them by area
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

	// run through the rectangles looking for the first that is
	// entirely contained in the polygon
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
