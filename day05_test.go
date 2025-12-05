package main

import (
	"strings"
	"testing"
)

func TestDay5a_Example(t *testing.T) {
	input := `3-5
10-14
16-20
12-18

1
5
8
11
17
32`
	expect := 3
	result := day5a(strings.SplitSeq(input, "\n"))
	if result != expect {
		t.Errorf("expected %d, got %d", expect, result)
	}
}

func TestDay5a(t *testing.T) {
	input := readlines("./inputs/day05.txt")
	expect := 862
	result := day5a(input)
	if result != expect {
		t.Errorf("expected %d, got %d", expect, result)
	}
}
