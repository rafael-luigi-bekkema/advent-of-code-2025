package main

import (
	"fmt"
	"iter"
)

func day1a(input iter.Seq[string]) int {
	total := 0
	dial := 50

	for line := range input {
		dir := line[0:1]
		amount := strtoint(line[1:])

		switch dir {
		case "R":
			dial = (dial + amount) % 100
		case "L":
			dial = (dial - amount) % 100
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
			dial -= amount

			total -= (dial-100)/100 + booltoint(dial+amount == 0)
			dial = (dial%100 + 100) % 100 // Compensate for negative number
		default:
			panic(fmt.Sprintf("Unexpected input: %s", line))
		}
	}

	return total
}
