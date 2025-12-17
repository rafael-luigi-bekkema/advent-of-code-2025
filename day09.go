package main

import (
	"fmt"
	"iter"
	"math/rand"
	"os"
)

type coord2d [2]int

func day9Parse(input iter.Seq[string]) []coord2d {
	var coords []coord2d
	for line := range input {
		var c coord2d
		if _, err := fmt.Sscanf(line, "%d,%d", &c[0], &c[1]); err != nil {
			panic(err)
		}
		coords = append(coords, c)
	}

	return coords
}

func day9a(input iter.Seq[string]) int {
	coords := day9Parse(input)

	var top int

	for i, c1 := range coords {
		for j := i + 1; j < len(coords); j++ {
			c2 := coords[j]
			s := (abs(c1[0]-c2[0]) + 1) * (abs(c1[1]-c2[1]) + 1)

			if s > top {
				top = s
			}
		}
	}

	return top
}

func day9svg(input iter.Seq[string], file string) {
	coords := day9Parse(input)

	var ylines, xlines [][3]int
	minx, miny := -1, -1
	var maxx, maxy int
	for i, c := range coords {
		j := (i + 1) % len(coords)
		c2 := coords[j]

		if c[0] > maxx {
			maxx = c[0]
		}
		if c[1] > maxy {
			maxy = c[1]
		}

		if minx == -1 || c[0] < minx {
			minx = c[0]
		}
		if miny == -1 || c[1] < miny {
			miny = c[1]
		}
		if c[0] == c2[0] {
			ylines = append(ylines, [3]int{min(c[1], c2[1]), max(c[1], c2[1]), c[0]})
			continue
		}
		if c[1] == c2[1] {
			xlines = append(xlines, [3]int{min(c[0], c2[0]), max(c[0], c2[0]), c[1]})
			continue
		}

		panic("mo match")
	}

	f, err := os.Create(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	strokeWidth := "50"

	fmt.Fprintf(f, `<!doctype html>`)
	fmt.Fprintf(f, `<body style="padding: 5vh 0; width: 100%%">`)
	fmt.Fprintf(f, `<div style="margin: 0 auto; display: block; width: fit-content">`)
	fmt.Fprintf(f, `<svg style="height: 90vh" viewBox="%d %d %d %d" xmlns="http://www.w3.org/2000/svg">`, minx, miny, maxx+1, maxy+1)
	for _, line := range xlines {
		fmt.Fprintf(f, `<line x1="%d" y1="%d" x2="%d" y2="%d" stroke="black" stroke-width="%s" />`, line[0], line[2], line[1], line[2], strokeWidth)
	}
	for _, line := range ylines {
		fmt.Fprintf(f, `<line x1="%d" y1="%d" x2="%d" y2="%d" stroke="black" stroke-width="%s" />`, line[2], line[0], line[2], line[1], strokeWidth)
	}
}

func day9b(input iter.Seq[string]) int {
	coords := day9Parse(input)

	var ylines, xlines [][3]int
	for i, c := range coords {
		j := (i + 1) % len(coords)
		c2 := coords[j]

		if c[0] == c2[0] {
			ylines = append(ylines, [3]int{min(c[1], c2[1]), max(c[1], c2[1]), c[0]})
			continue
		}
		if c[1] == c2[1] {
			xlines = append(xlines, [3]int{min(c[0], c2[0]), max(c[0], c2[0]), c[1]})
			continue
		}
		panic("mo match")
	}

	imap := map[coord2d]bool{}

	inside := func(c coord2d) bool {
		if val, ok := imap[c]; ok {
			return val
		}

		var count int

		// From left
		for _, l := range ylines {
			if l[2] <= c[0] && l[0] <= c[1] && l[1] >= c[1] {
				count++
			}
		}
		for _, l := range xlines {
			if l[2] == c[1] && l[0] <= c[0] && l[1] <= c[0] {
				count++
			}
		}

		val := count%2 != 0
		imap[c] = val

		return val
	}

	var total int

	for i, c1 := range coords {
		fmt.Println(i+1, len(coords))
	outer:
		for j := i + 1; j < len(coords); j++ {
			c2 := coords[j]

			s := (abs(c1[0]-c2[0]) + 1) * (abs(c1[1]-c2[1]) + 1)
			if total > 0 && s < total {
				continue
			}

			minx, maxx := min(c1[0], c2[0]), max(c1[0], c2[0])
			miny, maxy := min(c1[1], c2[1]), max(c1[1], c2[1])

			for _, x := range rand.Perm(maxx - minx + 1) {
				x := x + minx
				if !inside(coord2d{x, miny}) || !inside(coord2d{x, maxy}) {
					continue outer
				}
			}
			for y := range rand.Perm(maxy - miny + 1) {
				y := y + miny
				if !inside(coord2d{minx, y}) || !inside(coord2d{maxx, y}) {
					continue outer
				}
			}

			if total == 0 || s > total {
				fmt.Println(s, "size")
				total = s
			}
		}
	}

	return total
}
