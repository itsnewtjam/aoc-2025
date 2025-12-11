package day10

import (
	"aoc2025/util"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Day10 struct {
	s1 int
	s2 int
}

func (d *Day10) Title() string {
	return "Day 10 - Factory"
}

func (d *Day10) Part1() {
	// var input = util.ReadInput("inputs/day10-example.txt")
	var input = util.ReadInput("inputs/day10.txt")

	machines := []machine{}
	for _, line := range input {
		parts := strings.Fields(line)
		machine := machine{}
		for _, part := range parts {
			switch part[0] {
			case '[':
				var mask uint32 = 0
				for i, light := range part[1:len(part) - 1] {
					if light == '#' {
						mask |= 1 << uint(i)
					}
				}
				machine.lights = mask
			case '(':
				// storing the buttons as a bit mask, where each bit represents
				// a light. 1 = it toggles that light
				var mask uint32 = 0
				lights := strings.Split(part[1:len(part) - 1], ",")
				for _, light := range lights {
					lightInt, _ := strconv.Atoi(light)
					mask |= 1 << uint(lightInt)
				}
				machine.buttons = append(machine.buttons, mask)
			}
		}
		machines = append(machines, machine)
	}

	// just run a lil bfs to get the minimum :)
	total := 0
	for _, machine := range machines {
		queue := []queueItem{{
			value: 0,
			presses: 0,
		}}
		visited := make(map[uint32]bool)
		visited[0] = true

		for len(queue) > 0 {
			curr := queue[0]
			queue = queue[1:]

			if curr.value == machine.lights {
				total += curr.presses
				break
			}

			for _, button := range machine.buttons {
				newValue := curr.value ^ button
				if !visited[newValue] {
					visited[newValue] = true
					queue = append(queue, queueItem{
						value: newValue,
						presses: curr.presses + 1,
					})
				}
			}
		}
	}

	d.s1 = total
}

func (d *Day10) Part2() {
	// var input = util.ReadInput("inputs/day10-example.txt")
	var input = util.ReadInput("inputs/day10.txt")

	// holy moly this was rough. i feel like it shouldn't have been such
	// a pain, i have some understanding of systems of equations with
	// matrices. but the whole clamping free parameters thing really
	// threw me for a loop :S
	// there's probably a way cleaner way to run through some of this,
	// it kind of turned into slapping on bandages towards the end...

	machines := []machine{}
	for _, line := range input {
		parts := strings.Fields(line)
		machine := machine{}

		joltagePart := parts[len(parts) - 1]
		joltageStrings := strings.Split(joltagePart[1:len(joltagePart) - 1], ",")
		joltages := make([]int, len(joltageStrings))
		for i, joltage := range joltageStrings {
			j, _ := strconv.Atoi(joltage)
			joltages[i] = j
		}
		machine.joltages = joltages

		for _, part := range parts {
			if part[0] != '(' {
				continue
			}
			// bit masks for buttons again
			var mask uint32 = 0
			joltages := strings.SplitSeq(part[1:len(part) - 1], ",")
			for joltage := range joltages {
				joltageInt, _ := strconv.Atoi(joltage)
				mask |= 1 << uint(joltageInt)
			}
			machine.buttons = append(machine.buttons, mask)
			// here we're making note of the maximum times a button can be pressed
			// aka the minimum target joltage that is affected by the button
			buttonMax := math.MaxInt
			for i, jolt := range machine.joltages {
				if mask & (1 << uint(i)) > 0 {
					if jolt < buttonMax {
						buttonMax = jolt
					}
				}
			}
			machine.buttonMax = append(machine.buttonMax, buttonMax)
		}

		// now we build a matrix A[i][j] where button [j] = 1 if it
		// increments counter [i]. ie for the first example:
		//   [.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
		//   [0 0 0 0 1 1]
		//   [0 1 0 0 0 1]
		//   [0 0 1 1 1 0]
		//   [1 1 0 1 0 0]
		joltageButtons := make([][]int, len(machine.joltages))
		for i := range machine.joltages {
			joltageButton := make([]int, len(machine.buttons))
			for j, button := range machine.buttons {
				if button & (1 << uint(i)) > 0 {
					joltageButton[j] = 1
				} else {
					joltageButton[j] = 0
				}
			}
			joltageButtons[i] = joltageButton
		}
		machine.joltageButtons = joltageButtons
		machines = append(machines, machine)
	}

	total := 0
	for _, machine := range machines {
		// now we just work on getting the joltage/button matrix
		// into row echelon
		for i := 0; i < len(machine.joltageButtons[0]); i++ {
			nonZero := -1
			j := i
			for nonZero == -1 && j < len(machine.joltageButtons[0]) {
				// swapping columns until the one at i has a non-zero value
				if i != j {
					for k := 0; k < len(machine.joltageButtons); k++ {
						machine.joltageButtons[k][i], machine.joltageButtons[k][j] = machine.joltageButtons[k][j], machine.joltageButtons[k][i]
						// making sure to swap the button max so they stay in sync
						machine.buttonMax[i], machine.buttonMax[j] = machine.buttonMax[j], machine.buttonMax[i]
					}
				}
				for k := i; k < len(machine.joltageButtons); k++ {
					if machine.joltageButtons[k][i] != 0 {
						nonZero = k
						break
					}
				}
				j++
			}

			if nonZero == -1 {
				break
			}

			// swap rows so the diagonal at [i][i] is non-zero
			if i != nonZero {
				machine.joltageButtons[i], machine.joltageButtons[nonZero] = machine.joltageButtons[nonZero], machine.joltageButtons[i]
				// making sure to swap joltages so they stay in sync
				machine.joltages[i], machine.joltages[nonZero] = machine.joltages[nonZero], machine.joltages[i]
			}

			// reduce the other rows
			for j := i + 1; j < len(machine.joltageButtons); j++ {
				if machine.joltageButtons[i][i] != 0 {
					x := machine.joltageButtons[i][i]
					y := machine.joltageButtons[j][i] * -1
					d := gcd(x, y)
					for k := range len(machine.joltageButtons[i]) {
						machine.joltageButtons[j][k] = (y * machine.joltageButtons[i][k] + x * machine.joltageButtons[j][k]) / d
					}
					machine.joltages[j] = (y * machine.joltages[i] + x * machine.joltages[j]) / d
				}
			}
		}

		// remove any all-zero rows
		noZeros := []int{}
		for i, row := range machine.joltageButtons {
			for j := range row {
				if row[j] != 0 {
					noZeros = append(noZeros, i)
					break
				}
			}
		}
		newJoltageButtons := [][]int{}
		newJoltages := []int{}
		for _, i := range noZeros {
			newJoltageButtons = append(newJoltageButtons, machine.joltageButtons[i])
			newJoltages = append(newJoltages, machine.joltages[i])
		}
		machine.joltageButtons = newJoltageButtons
		machine.joltages = newJoltages

		// back substitution to get just the diagonal, all other values in the
		// diagonal range will be zero. this way each row is a determined variable
		for i := len(machine.joltageButtons) - 1; i >= 0; i-- {
			for j := 0; j < i; j++ {
				if machine.joltageButtons[i][i] != 0 {
					x := machine.joltageButtons[i][i]
					y := machine.joltageButtons[j][i] * -1
					d := gcd(x, y)
					for k := range machine.joltageButtons[i] {
						machine.joltageButtons[j][k] = (y * machine.joltageButtons[i][k] + x * machine.joltageButtons[j][k]) / d
					}
					machine.joltages[j] = (y * machine.joltages[i] + x * machine.joltages[j]) / d
				}
			}
		}

		// clean up any negatives
		for i := 0; i < len(machine.joltageButtons); i++ {
			if machine.joltageButtons[i][i] < 0 {
        for j := range machine.joltageButtons[i] {
					machine.joltageButtons[i][j] *= -1
        }
        machine.joltages[i] *= -1
			}
		}

		k := len(machine.joltageButtons[0]) - len(machine.joltageButtons)

		// calculating the bounds for the free params, making sure to 
		// get them so that all determined variables stay non-negative
		freeBounds := make([]int, k)
		for paramIdx := range k {
			colIdx := len(machine.joltageButtons[0]) - k + paramIdx
			maxNeeded := 0
			
			for i := 0; i < len(machine.joltageButtons); i++ {
				coeff := machine.joltageButtons[i][colIdx]
				if coeff < 0 {
					coeff *= -1
				}
				diag := machine.joltageButtons[i][i]
				if diag < 0 {
					diag *= -1
				}
				target := machine.joltages[i]
				if target < 0 {
					target *= -1
				}
				
				if coeff != 0 && diag != 0 {
					estimate := target
					if coeff > 0 {
							estimate = target / coeff
					}
					maxNeeded = max(maxNeeded, estimate)
				}
			}
			
			originalBound := machine.buttonMax[len(machine.buttonMax) - k + paramIdx]
			freeBounds[paramIdx] = max(originalBound, min(maxNeeded + 10, 200))
		}

		// run through all the parameter combos and keep track of the
		// minimum number of presses
		mins := math.MaxInt
		for _, combo := range getCombos(k, freeBounds) {
			solution := 0
			for _, c := range combo {
				solution += c
			}

			for i := 0; i < len(machine.joltageButtons); i++ {
				cc := []int{}
				for j := range combo {
					cc = append(cc, combo[j] * machine.joltageButtons[i][len(machine.joltageButtons[0]) - k + j])
				}

				sum := 0
				for _, val := range cc {
					sum += val
				}

				s := machine.joltages[i] - sum
				diag := machine.joltageButtons[i][i]
				a := s / diag
				if s % diag != 0 {
					solution = math.MaxInt
					break
				}
				if (diag > 0 && a < 0) || (diag < 0 && a > 0) {
					solution = math.MaxInt
					break
				}
				if diag < 0 {
					a *= -1
				}

				solution += a
			}

			mins = min(mins, solution)
		}

		total += mins
	}


	d.s2 = total
}

func (d *Day10) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day10) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}

type machine struct {
	lights uint32
	buttons []uint32
	joltageButtons [][]int
	buttonMax []int
	joltages []int
}

type queueItem struct {
	value uint32
	presses int
}

func gcd(a, b int) int {
	if a < 0 {
		a *= -1
	}
	if b < 0 {
		b *= -1
	}
	for b != 0 {
		a, b = b, a % b
	}
	return a
}

func getCombos(n int, c []int) [][]int {
	if n == 0 {
		return [][]int{{}}
	}

	var ret [][]int
	for i := 0; i <= c[len(c) - n]; i++ {
		for _, l := range getCombos(n - 1, c) {
			ret = append(ret, append([]int{i}, l...))
		}
	}
	return ret
}
