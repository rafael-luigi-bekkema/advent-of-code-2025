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

	return
}

func day5b(input iter.Seq[string]) (total int) {
	var ranges [][2]int
	for line := range input {
		if line == "" {
			break
		}
		var from, to int
		n, _ := fmt.Sscanf(line, "%d-%d", &from, &to)
		if n != 2 {
			panic("Unexpected range input: " + line)
		}

		ranges = append(ranges, [2]int{from, to})
	}

	// Merge overlapping ranges until there are none
	for {
		overlap := false

	outer:
		for i, r1 := range ranges {
			for j := i + 1; j < len(ranges); j++ {
				r2 := ranges[j]
				// r2 starts before, but ends after start of r1
				if r2[0] < r1[0] && r2[1] >= r1[0] {
					ranges[i][0] = r2[0]
					overlap = true
				}
				// r2 ends after, but starts before end of r1
				if r2[1] > r1[1] && r2[0] <= r1[1] {
					ranges[i][1] = r2[1]
					overlap = true
				}
				// r2 falls completely with r1
				if r2[0] >= r1[0] && r2[1] <= r1[1] {
					overlap = true
				}
				// r2 same as r1
				if r2 == r1 {
					overlap = true
				}
				if overlap {
					// remove r2 from list
					ranges = append(ranges[:j], ranges[j+1:]...)
					break outer
				}
			}
		}

		if !overlap {
			break
		}
	}

	for _, r := range ranges {
		total += r[1] - r[0] + 1
	}

	return
}
