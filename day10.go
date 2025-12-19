package main

import (
	"fmt"
	"iter"
	"slices"
	"strings"
)

type Machine struct {
	lights  string
	buttons [][]int
	joltage []int
}

func day10parse(input iter.Seq[string]) (machines []Machine) {
	for line := range input {
		parts := strings.Split(line, " ")

		var m Machine

		m.lights = parts[0][1 : len(parts[0])-1]

		for n := range strings.SplitSeq(parts[len(parts)-1][1:len(parts[len(parts)-1])-1], ",") {
			m.joltage = append(m.joltage, strtoint(n))
		}

		for _, btn := range parts[1 : len(parts)-1] {
			var btnn []int
			for nbtn := range strings.SplitSeq(btn[1:len(btn)-1], ",") {
				btnn = append(btnn, strtoint(nbtn))
			}
			m.buttons = append(m.buttons, btnn)
		}

		machines = append(machines, m)
	}

	return
}

func day10a(input iter.Seq[string]) int {
	machines := day10parse(input)

	type Item struct {
		Count  int
		Btn    int
		Lights string
	}

	type CacheItem struct {
		Btn    int
		Lights string
	}

	solved := func(m *Machine, item *Item) (string, bool) {
		if item == nil {
			return strings.Repeat(".", len(m.lights)), false
		}

		lights := []byte(item.Lights)
		for _, i := range m.buttons[item.Btn] {
			if lights[i] == '#' {
				lights[i] = '.'
			} else {
				lights[i] = '#'
			}
		}
		slights := string(lights)

		return slights, slights == m.lights
	}

	solve := func(m *Machine) *Item {
		var queue []*Item
		queue = append(queue, nil)
		explored := map[CacheItem]struct{}{}

		for len(queue) > 0 {
			item := queue[0]
			queue = queue[1:]

			lights, ok := solved(m, item)
			if ok {
				return item
			}
			for i := range m.buttons {
				c := 1
				if item != nil {
					c = item.Count + 1
				}
				ss := Item{Count: c, Btn: i, Lights: lights}
				ci := CacheItem{Btn: i, Lights: lights}
				if _, ok := explored[ci]; !ok {
					explored[ci] = set
					queue = append(queue, &ss)
				}
			}
		}

		return nil
	}

	var total int

	for _, m := range machines {
		total += solve(&m).Count
	}

	return total
}

func day10b(input iter.Seq[string]) int {
	machines := day10parse(input)

	type Item struct {
		Count   int
		Btn     int
		Joltage []int
	}

	type CacheItem struct {
		Btn     int
		Joltage string
	}

	solved := func(m *Machine, btns []int) (over bool, same bool) {
		joltage := make([]int, len(m.joltage))
		for i, n := range btns {
			for _, btn := range m.buttons[i] {
				joltage[btn] += n
				if joltage[btn] > m.joltage[btn] {
					return true, false
				}

			}
		}

		return false, slices.Compare(joltage, m.joltage) == 0
	}

	var solveit func(m *Machine, btns []int, cmax int) int
	solveit = func(m *Machine, btns []int, cmax int) int {
		if cmax != -1 && sumslice(btns) >= cmax {
			return -1
		}

		fmt.Println(btns, cmax)

		over, ok := solved(m, btns)

		if over {
			return -1
		}

		if ok {
			return sumslice(btns)
		}

		total := cmax
		for i := range m.buttons {
			btns := slices.Clone(btns)
			btns[i]++
			val := solveit(m, btns, total)
			if val != -1 && (total == -1 || val < total) {
				total = val
			}
		}

		return total
	}

	var total int

	for _, m := range machines {
		total += solveit(&m, make([]int, len(m.buttons)), -1)
		fmt.Println("done!")
	}

	return total
}
