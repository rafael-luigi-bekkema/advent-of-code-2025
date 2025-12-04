package main

import (
	"testing"
)

func TestDay4a_Example(t *testing.T) {
	input := `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`
	expect := 13
	result := day4a(input)
	if result != expect {
		t.Errorf("expected %d, got %d", expect, result)
	}
}

func TestDay4a(t *testing.T) {
	input := readfile("./inputs/day04.txt")
	expect := 1451
	result := day4a(input)
	if result != expect {
		t.Errorf("expected %d, got %d", expect, result)
	}
}

func TestDay4b_Example(t *testing.T) {
	input := `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`
	expect := 43
	result := day4b(input)
	if result != expect {
		t.Errorf("expected %d, got %d", expect, result)
	}
}

func TestDay4b(t *testing.T) {
	input := readfile("./inputs/day04.txt")
	expect := 8701
	result := day4b(input)
	if result != expect {
		t.Errorf("expected %d, got %d", expect, result)
	}
}
