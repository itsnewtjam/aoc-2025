package day06

import (
	"aoc2025/util"
	"fmt"
	"strconv"
	"strings"
)

type Day06 struct {
	s1 int
	s2 int
}

func (d *Day06) Title() string {
	return "Day 06 - Trash Compactor"
}

func (d *Day06) Part1() {
	// var input = util.ReadInput("inputs/day06-example.txt")
	var input = util.ReadInput("inputs/day06.txt")

	cols := len(strings.Fields(input[0]))
	problems := make([]problem, cols)
	for idx, line := range input {
		// this is a sweet function... splits the data for us!
		fields := strings.Fields(line)
		for i, val := range fields {
			if idx + 1 == len(input) {
				// we're at the last line so set the operations
				problems[i].operation = val
			} else {
				// add the number to the appropraite problem
				intVal, err := strconv.Atoi(val)
				if err != nil {
					panic(err)
				}
				problems[i].numbers = append(problems[i].numbers, intVal)
			}
		}
	}

	total := 0
	for _, problem := range problems {
		subtotal := problem.numbers[0]
		for _, num := range problem.numbers[1:] {
			switch problem.operation {
			case "+":
			subtotal += num
			case "*":
			subtotal *= num
			}
		}
		total += subtotal
	}

	d.s1 = total
}

func (d *Day06) Part2() {
	// var input = util.ReadInput("inputs/day06-example.txt")
	var input = util.ReadInput("inputs/day06.txt")

	// here we just loop through the width of the input (aka by column)
	// and build the problems as we go
	problems := []problem{}
	inProblem := false
	currProblem := problem{}
	for i := 0; i < len(input[0]); i++ {
		opRow := input[len(input) - 1][i]
		if opRow != ' ' {
			// we encountered an operator on the last row
			if inProblem == true {
				// we were building a problem, so remove the last number (will be a
				// zero from the space between columns) and save the problem
				currProblem.numbers = currProblem.numbers[:len(currProblem.numbers) - 1]
				problems = append(problems, currProblem)
			}
			// start a new problem
			inProblem = true
			currProblem = problem{operation: string(opRow)}
		}
		// build the number from the column and add it to the problem
		num := getColumn(input[:len(input) - 1], i)
		currProblem.numbers = append(currProblem.numbers, num)
	}
	// make sure to not forget to save the last problem!
	problems = append(problems, currProblem)

	total := 0
	for _, problem := range problems {
		subtotal := problem.numbers[0]
		for _, num := range problem.numbers[1:] {
			switch problem.operation {
			case "+":
			subtotal += num
			case "*":
			subtotal *= num
			}
		}
		total += subtotal
	}

	d.s2 = total
}

func (d *Day06) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day06) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}

type problem struct {
	numbers []int
	operation string
}

func getColumn(rows []string, col int) (number int) {
	var column []byte
	for _, row := range rows {
		column = append(column, row[col])
	}
	columnStr := strings.TrimSpace(string(column))
	if len(columnStr) == 0 {
		return 0
	}
	number, err := strconv.Atoi(columnStr)
	if err != nil {
		panic(err)
	}
	return number
}
