package main

import (
	"strings"
	"testing"
)

const testInput = `[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}
[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}`

func TestDay10a_Example(t *testing.T) {
	expect := 7
	result := day10a(strings.SplitSeq(testInput, "\n"))
	if result != expect {
		t.Errorf("expected %d, got %d", expect, result)
	}
}

func TestDay10a(t *testing.T) {
	input := readlines("./inputs/day10.txt")
	expect := 411
	result := day10a(input)
	if result != expect {
		t.Errorf("expected %d, got %d", expect, result)
	}
}
