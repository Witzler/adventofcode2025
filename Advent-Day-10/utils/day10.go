package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

const MaxJoltagesLength = 10

type Machine struct {
	TargetDiagram  []rune
	Buttons        [][]int
	TargetJoltages [MaxJoltagesLength]int
	JoltagesLength int
}

/* =========================
   Parsing
========================= */

func ParseMachines(scanner *bufio.Scanner) []Machine {
	machines := []Machine{}

	for scanner.Scan() {
		text := scanner.Text()
		split1 := strings.Split(text, "] ")
		diagram := []rune(split1[0][1:])

		split2 := strings.Split(split1[1], " {")
		buttons := strings.Split(split2[0], " ")

		buttonsArr := [][]int{}
		for _, button := range buttons {
			split3 := strings.Split(button[1:len(button)-1], ",")
			intButton := []int{}
			for _, s := range split3 {
				v, err := strconv.Atoi(s)
				Check(err)
				intButton = append(intButton, v)
			}
			buttonsArr = append(buttonsArr, intButton)
		}

		joltagesArr := strings.Split(split2[1][:len(split2[1])-1], ",")
		joltagesLength := len(joltagesArr)
		if joltagesLength > MaxJoltagesLength {
			panic("Too many joltages")
		}

		joltages := [MaxJoltagesLength]int{}
		for i := 0; i < joltagesLength; i++ {
			v, err := strconv.Atoi(joltagesArr[i])
			Check(err)
			joltages[i] = v
		}

		machines = append(machines, Machine{
			TargetDiagram:  diagram,
			Buttons:        buttonsArr,
			TargetJoltages: joltages,
			JoltagesLength: joltagesLength,
		})
	}

	return machines
}

/* =========================
   Part 1
========================= */

func SolveMachinesPart1(machines []Machine) int {
	sum := 0

	for _, machine := range machines {
		length := len(machine.TargetDiagram)
		initial := slices.Repeat([]rune{'.'}, length)
		diagrams := [][]rune{initial}

		for level := 1; ; level++ {
			next := [][]rune{}
			for _, d := range diagrams {
				for _, button := range machine.Buttons {
					nd := make([]rune, length)
					copy(nd, d)
					for _, idx := range button {
						nd[idx] = switchIndicator(nd[idx])
					}
					if slices.Equal(nd, machine.TargetDiagram) {
						sum += level
						goto done
					}
					next = append(next, nd)
				}
			}
			diagrams = next
		}
	done:
	}

	return sum
}

func switchIndicator(r rune) rune {
	if r == '.' {
		return '#'
	}
	if r == '#' {
		return '.'
	}
	panic("invalid rune")
}

/* =========================
   Part 2
========================= */

func SolveMachinesPart2(machines []Machine) int {
	sum := 0

	for _, m := range machines {

		if _, err := exec.LookPath("z3"); err == nil {
			if level, ok := solveWithZ3(m); ok {
				sum += level
				continue
			}
		}

		if level, ok := fallbackSolveMachine(m); ok {
			sum += level
		}
	}

	return sum
}

/* =========================
   Z3 Solver
========================= */

func solveWithZ3(machine Machine) (int, bool) {
	n := len(machine.Buttons)
	if n == 0 {
		return 0, false
	}

	var sb strings.Builder
	sb.WriteString("(set-option :produce-models true)\n")

	for i := 0; i < n; i++ {
		sb.WriteString(fmt.Sprintf("(declare-const b%d Int)\n", i))
		sb.WriteString(fmt.Sprintf("(assert (>= b%d 0))\n", i))
	}

	for j := 0; j < machine.JoltagesLength; j++ {
		sb.WriteString("(assert (= (+")
		for i := 0; i < n; i++ {
			if slices.Contains(machine.Buttons[i], j) {
				sb.WriteString(fmt.Sprintf(" b%d", i))
			}
		}
		sb.WriteString(fmt.Sprintf(") %d))\n", machine.TargetJoltages[j]))
	}

	sb.WriteString("(minimize (+")
	for i := 0; i < n; i++ {
		sb.WriteString(fmt.Sprintf(" b%d", i))
	}
	sb.WriteString("))\n(check-sat)\n(get-model)\n")

	cmd := exec.Command("z3", "-in", "-smt2")
	cmd.Stdin = strings.NewReader(sb.String())

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	if err := cmd.Run(); err != nil {
		return 0, false
	}

	re := regexp.MustCompile(`Int\s+([0-9]+)\)`)
	matches := re.FindAllStringSubmatch(out.String(), -1)
	if matches == nil {
		return 0, false
	}

	total := 0
	for _, m := range matches {
		v, _ := strconv.Atoi(m[1])
		total += v
	}

	return total, true
}

/* =========================
   Fallback BFS + Pruning
========================= */

func fallbackSolveMachine(machine Machine) (int, bool) {
	initial := [MaxJoltagesLength]int{}
	candidates := map[[MaxJoltagesLength]int]bool{initial: true}

	for level := 1; level <= 1000; level++ {
		next := map[[MaxJoltagesLength]int]bool{}

		for c := range candidates {
			for _, button := range machine.Buttons {
				nc := c
				for _, idx := range button {
					nc[idx]++
				}

				if !checkJoltages(nc, machine.TargetJoltages, machine.JoltagesLength) {
					continue
				}

				if nc == machine.TargetJoltages {
					return level, true
				}

				next[nc] = true
			}
		}

		candidates = pruneCandidates(next, machine.JoltagesLength)
		if len(candidates) == 0 {
			return 0, false
		}
	}

	return 0, false
}

func checkJoltages(cur, target [MaxJoltagesLength]int, length int) bool {
	for i := 0; i < length; i++ {
		if cur[i] > target[i] {
			return false
		}
	}
	return true
}

func pruneCandidates(c map[[MaxJoltagesLength]int]bool, length int) map[[MaxJoltagesLength]int]bool {
	keys := make([][MaxJoltagesLength]int, 0, len(c))
	for k := range c {
		keys = append(keys, k)
	}

	keep := map[[MaxJoltagesLength]int]bool{}
	for i, a := range keys {
		dominated := false
		for j, b := range keys {
			if i == j {
				continue
			}
			le, eq := true, true
			for d := 0; d < length; d++ {
				if a[d] > b[d] {
					le = false
					break
				}
				if a[d] != b[d] {
					eq = false
				}
			}
			if le && !eq {
				dominated = true
				break
			}
		}
		if !dominated {
			keep[a] = true
		}
	}

	return keep
}
