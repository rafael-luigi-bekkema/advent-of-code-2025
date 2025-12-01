package main

import (
	"fmt"
	"iter"
	"strconv"
)

func day1a(input iter.Seq[string]) int {
	var total int
	dial := 50

	for line := range input {
		dir := line[0:1]
		amount, err := strconv.Atoi(line[1:])

		switch {
		case dir == "R":
			dial += amount
			for dial > 99 {
				dial -= 100
			}
		case dir == "L":
			dial -= amount
			for dial < 0 {
				dial += 100
			}
		case err != nil:
			panic(fmt.Sprintf("Error: %s", err))
		default:
			panic(fmt.Sprintf("Unexpected input: %s", line))
		}

		if dial == 0 {
			total++
		}
	}

	return total
}

func day1b(input iter.Seq[string]) int {
	var total int
	dial := 50

	for line := range input {
		dir := line[0:1]
		amount, err := strconv.Atoi(line[1:])

		switch {
		case dir == "R":
			dial += amount

			total += dial / 100
			dial = dial % 100
		case dir == "L":
			startAtZero := dial == 0
			dial -= amount
			for dial < 0 {
				dial += 100
				if startAtZero {
					startAtZero = false
					continue
				}
				total += 1
			}
			if dial == 0 {
				total += 1
			}
		case err != nil:
			panic(fmt.Sprintf("Error: %s", err))
		default:
			panic(fmt.Sprintf("Unexpected input: %s", line))
		}
	}

	return total
}
