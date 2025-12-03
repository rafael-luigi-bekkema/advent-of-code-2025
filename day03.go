package main

import (
	"iter"
)

func day3(input iter.Seq[string], size int) (total int) {
	for bank := range input {
		bbank := []byte(bank)

		maxi, maxes := -1, make([]byte, size)
		for n := range size {
			for i := maxi + 1; i < len(bbank)-(size-1-n); i++ {
				if maxes[n] == 0 || bbank[i] > maxes[n] {
					maxes[n] = bbank[i]
					maxi = i
				}
			}
		}

		total += strtoint(string(maxes[:]))
	}
	return
}
