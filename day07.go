package main

import (
	"fmt"
	"iter"
)

var zero = struct{}{}

func day7a(input iter.Seq[string]) (total int) {
	first := true
	beams := map[int]struct{}{}
	for line := range input {
		bline := []byte(line)
		for i, chr := range []byte(line) {
			switch chr {
			case 'S':
				beams[i] = zero
			case '^':
				_, ok := beams[i]
				if ok {
					total++
					delete(beams, i)
					beams[i-1] = zero
					beams[i+1] = zero
				}
			case '.':
			}
		}
		for i := range beams {
			if first {
				first = false
				continue
			}
			bline[i] = '|'
		}
		fmt.Println(string(bline))
	}

	return total
}
