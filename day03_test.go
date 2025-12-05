package main

import (
	"strings"
	"testing"
)

func TestDay3a_Example(t *testing.T) {
	input := `987654321111111
811111111111119
234234234234278
818181911112111`
	expect := 357
	result := day3(strings.SplitSeq(input, "\n"), 2)
	if result != expect {
		t.Errorf("expected %d, got %d", expect, result)
	}
}

func TestDay3a(t *testing.T) {
	input := readlines("./inputs/day03.txt")
	expect := 17100
	result := day3(input, 2)
	if result != expect {
		t.Errorf("expected %d, got %d", expect, result)
	}
}

func TestDay3b_Example(t *testing.T) {
	input := `987654321111111
811111111111119
234234234234278
818181911112111`
	expect := 3121910778619
	result := day3(strings.SplitSeq(input, "\n"), 12)
	if result != expect {
		t.Errorf("expected %d, got %d", expect, result)
	}
}

func TestDay3b(t *testing.T) {
	input := readlines("./inputs/day03.txt")
	expect := 170418192256861
	result := day3(input, 12)
	if result != expect {
		t.Errorf("expected %d, got %d", expect, result)
	}
}
