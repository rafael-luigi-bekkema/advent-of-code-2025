package main

import (
	"fmt"
	"iter"
)

func day1a(input iter.Seq[string]) int {
	var total int
	dial := 50

	for line := range input {
		dir := line[0:1]
		amount := strtoint(line[1:])

		switch dir {
		case "R":
			dial += amount
			for dial > 99 {
				dial -= 100
			}
		case "L":
			dial -= amount
			for dial < 0 {
				dial += 100
			}
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
		amount := strtoint(line[1:])

		switch dir {
		case "R":
			dial += amount

			total += dial / 100
			dial = dial % 100
		case "L":
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
		default:
			panic(fmt.Sprintf("Unexpected input: %s", line))
		}
	}

	return total
}
