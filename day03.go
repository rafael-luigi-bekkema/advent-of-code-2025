package main

import (
	"iter"
)

func day3a(input iter.Seq[string]) (total int) {
	for bank := range input {
		bbank := []byte(bank)
		var maxj byte
		var maxjset bool
		var maxi int
		for i, jolt := range bbank[:len(bbank)-1] {
			if !maxjset || jolt > maxj {
				maxjset = true
				maxj = jolt
				maxi = i
			}
		}
		var maxk byte
		var maxkset bool
		for _, jolt := range bbank[maxi+1:] {
			if !maxkset || jolt > maxk {
				maxkset = true
				maxk = jolt
			}
		}

		result := string([]byte{maxj, maxk})
		total += strtoint(result)
	}
	return
}
