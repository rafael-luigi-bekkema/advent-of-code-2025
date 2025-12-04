package main

import (
	"bytes"
	"strings"
)

func day4a(input string) (total int) {
	rows := strings.Split(input, "\n")
	for y := range rows {
		for x := range rows[y] {
			count := 0
			if rows[y][x] != '@' {
				continue
			}
			for i := range 9 {
				cy := y - 1 + i/3
				cx := x - 1 + i%3
				if cy == y && cx == x ||
					cy < 0 || cy >= len(rows) || cx < 0 || cx >= len(rows[y]) {
					continue
				}

				if rows[cy][cx] == '@' {
					count++
				}
			}
			if count < 4 {
				total++
			}
		}
	}
	return
}

func day4b(input string) (total int) {
	rows := bytes.Split([]byte(input), []byte{'\n'})

	removeRolls := func() int {
		subtotal := 0
		var removes [][2]int
		for y := range rows {
			for x := range rows[y] {
				count := 0
				if rows[y][x] != '@' {
					continue
				}
				for i := range 9 {
					cy := y - 1 + i/3
					cx := x - 1 + i%3
					if cy == y && cx == x ||
						cy < 0 || cy >= len(rows) || cx < 0 || cx >= len(rows[y]) {
						continue
					}

					if rows[cy][cx] == '@' {
						count++
					}
				}
				if count < 4 {
					subtotal++
					removes = append(removes, [2]int{y, x})
				}
			}
		}

		for _, pos := range removes {
			rows[pos[0]][pos[1]] = '.'
		}

		return subtotal
	}

	for {
		subtotal := removeRolls()
		if subtotal == 0 {
			break
		}
		total += subtotal
	}

	return
}
