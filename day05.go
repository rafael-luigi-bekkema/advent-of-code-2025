package main

import (
	"fmt"
	"iter"
	"strconv"
)

func day5a(input iter.Seq[string]) (total int) {
	var imode bool
	var ranges [][2]int
	for line := range input {
		if line == "" {
			imode = true
			continue
		}
		if imode {
			n, err := strconv.Atoi(line)
			if err != nil {
				panic("Not an int: " + line)
			}
			for _, r := range ranges {
				if n >= r[0] && n <= r[1] {
					total++
					break
				}
			}
			continue
		}
		var from, to int
		n, _ := fmt.Sscanf(line, "%d-%d", &from, &to)
		if n != 2 {
			panic("Unexpected range input: " + line)
		}
		ranges = append(ranges, [2]int{from, to})
	}

	return total
}
