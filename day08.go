package main

import (
	"cmp"
	"fmt"
	"iter"
	"maps"
	"math"
	"slices"
)

type coord [3]int

func distance3d(p, q coord) float64 {
	return math.Sqrt(
		math.Pow(float64(p[0]-q[0]), 2) + math.Pow(float64(p[1]-q[1]), 2) + math.Pow(float64(p[2]-q[2]), 2),
	)
}

func mergeCircuits(circuits []map[coord]empty) []map[coord]empty {
repeat:
	for i, c1 := range circuits {
		for j := i + 1; j < len(circuits); j++ {
			c2 := circuits[j]
			for cc2 := range c2 {
				if _, ok := c1[cc2]; !ok {
					continue
				}

				maps.Copy(circuits[i], c2)

				circuits = slices.Delete(circuits, j, j+1)
				goto repeat
			}
		}
	}

	return circuits
}

type cdist struct {
	c1, c2 coord
	dist   float64
}

func day8prep(input iter.Seq[string]) ([]coord, []cdist) {
	var coords []coord

	for line := range input {
		var c coord
		_, err := fmt.Sscanf(line, "%d,%d,%d", &c[0], &c[1], &c[2])
		if err != nil {
			panic(err)
		}

		coords = append(coords, c)
	}

	var dists []cdist
	for i, c1 := range coords {
		for j := i + 1; j < len(coords); j++ {
			c2 := coords[j]

			dists = append(dists, cdist{c1, c2, distance3d(c1, c2)})
		}
	}

	slices.SortFunc(dists, func(a, b cdist) int {
		return cmp.Compare(a.dist, b.dist)
	})

	return coords, dists
}

func day8a(input iter.Seq[string], n int) int {
	_, dists := day8prep(input)

	var circuits []map[coord]empty

	for _, dist := range dists[:n] {
		circuits = append(circuits, map[coord]empty{
			dist.c1: set,
			dist.c2: set,
		})
	}

	circuits = mergeCircuits(circuits)

	var lens []int
	for _, c := range circuits {
		lens = append(lens, len(c))
	}
	slices.Sort(lens)
	return mulslice(lens[len(lens)-3:])
}

func day8b(input iter.Seq[string]) int {
	coords, dists := day8prep(input)

	var circuits []map[coord]empty

	for _, dist := range dists {
		circuits = append(circuits, map[coord]empty{
			dist.c1: set,
			dist.c2: set,
		})

		circuits = mergeCircuits(circuits)
		if len(circuits[0]) == len(coords) {
			return dist.c1[0] * dist.c2[0]
		}
	}

	panic("Failed to find solution")
}
