package main

import (
	"strings"
	"testing"
)

const day9TestInput = `7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`

func TestDay9a_Example(t *testing.T) {
	expect := 50
	result := day9a(strings.SplitSeq(day9TestInput, "\n"))

	if result != expect {
		t.Errorf("expected %d, got %d", expect, result)
	}
}

func TestDay9a(t *testing.T) {
	input := readlines("./inputs/day09.txt")
	expect := 4781235324
	result := day9a(input)

	if result != expect {
		t.Errorf("expected %d, got %d", expect, result)
	}
}
