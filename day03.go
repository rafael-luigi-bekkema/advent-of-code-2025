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

func day3b(input iter.Seq[string]) (total int) {
	for bank := range input {
		bbank := []byte(bank)

		maxi := -1
		var maxes [12]byte
		var maxset [12]bool
		for n := range 12 {
			for i := maxi + 1; i < len(bbank)-(11-n); i++ {
				jolt := bbank[i]
				if !maxset[n] || jolt > maxes[n] {
					maxset[n] = true
					maxes[n] = jolt
					maxi = i
				}
			}
		}

		result := string(maxes[:])
		total += strtoint(result)
	}
	return
}
