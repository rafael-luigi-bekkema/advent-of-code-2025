package main

import (
	"iter"
	"strings"
)

type Machine struct {
	lights  string
	buttons [][]int
	joltage []int
}

func day10a(input iter.Seq[string]) int {
	var machines []Machine
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
