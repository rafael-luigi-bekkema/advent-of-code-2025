package main

import (
	"fmt"
	"iter"
)

type coord2d [2]int

func day9a(input iter.Seq[string]) int {
	var coords []coord2d
	for line := range input {
		var c coord2d
		if _, err := fmt.Sscanf(line, "%d,%d", &c[0], &c[1]); err != nil {
			panic(err)
		}
		coords = append(coords, c)
	}

	var top int

	for i, c1 := range coords {
		for j := i + 1; j < len(coords); j++ {
			c2 := coords[j]
			s := (abs(c1[0]-c2[0]) + 1) * (abs(c1[1]-c2[1]) + 1)

			if s > top {
				fmt.Println(c1, c2)
				top = s
			}
		}
	}

	return top
}
