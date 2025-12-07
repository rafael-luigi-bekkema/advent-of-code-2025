package main

import (
	"iter"
)

type empty struct{}

var set = empty{}

func day7a(input iter.Seq[string]) (total int) {
	beams := map[int]empty{}
	for line := range input {
		for i, chr := range []byte(line) {
			switch chr {
			case 'S':
				beams[i] = set
			case '^':
				if _, ok := beams[i]; ok {
					total++
					delete(beams, i)
					beams[i-1] = set
					beams[i+1] = set
				}
			}
		}
	}

	return total
}

func day7b(input iter.Seq[string]) (total int) {
	var lines [][]byte
	for line := range input {
		lines = append(lines, []byte(line))
	}

	beams := map[int]int{}
	for i, line := range lines {
		for i, chr := range line {
			switch chr {
			case 'S':
				beams[i] = 1
			case '^':
				if val, ok := beams[i]; ok {
					delete(beams, i)
					beams[i-1] += val
					beams[i+1] += val
				}
			}
		}
		if i == len(lines)-1 {
			for _, val := range beams {
				total += val
			}
		}
	}

	return
}
