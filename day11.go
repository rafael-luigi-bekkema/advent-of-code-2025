package main

import (
	"iter"
	"slices"
	"strings"
)

func day11a(input iter.Seq[string]) int {
	outputs := map[string][]string{}

	for line := range input {
		from, to, ok := strings.Cut(line, ": ")
		if !ok {
			panic("unexpected line: " + line)
		}
		outputs[from] = strings.Split(to, " ")
	}

	var stack [][]string
	stack = append(stack, []string{"you"})
	explored := map[string]empty{}

	var total int

	for len(stack) > 0 {
		item := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		key := strings.Join(item, " ")
		if _, ok := explored[key]; !ok {
			explored[key] = set

			if item[len(item)-1] == "out" {
				total++
				continue
			}

			for _, device := range outputs[item[len(item)-1]] {
				newitem := append(slices.Clone(item), device)
				stack = append(stack, newitem)
			}
		}
	}

	return total
}
