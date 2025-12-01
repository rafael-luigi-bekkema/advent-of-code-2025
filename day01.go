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
			for range amount {
				dial++
				if dial == 100 {
					dial = 0
				}
				if dial == 0 {
					total++
				}
			}
		case dir == "L":
			for range amount {
				dial--
				if dial == -1 {
					dial = 99
				}
				if dial == 0 {
					total++
				}
			}
		case err != nil:
			panic(fmt.Sprintf("Error: %s", err))
		default:
			panic(fmt.Sprintf("Unexpected input: %s", line))
		}
	}

	return total
}
