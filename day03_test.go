package main

import (
	"strings"
	"testing"
)

func Test3a_Example(t *testing.T) {
	input := `987654321111111
811111111111119
234234234234278
818181911112111`
	expect := 357
	result := day3a(strings.SplitSeq(input, "\n"))
	if result != expect {
		t.Errorf("expected %d, got %d", expect, result)
	}
}

func Test3a(t *testing.T) {
	input := readlines("./inputs/day03.txt")
	expect := 17100
	result := day3a(input)
	if result != expect {
		t.Errorf("expected %d, got %d", expect, result)
	}
}

func Test3b_Example(t *testing.T) {
	input := `987654321111111
811111111111119
234234234234278
818181911112111`
	expect := 3121910778619
	result := day3b(strings.SplitSeq(input, "\n"))
	if result != expect {
		t.Errorf("expected %d, got %d", expect, result)
	}
}

func Test3b(t *testing.T) {
	input := readlines("./inputs/day03.txt")
	expect := 170418192256861
	result := day3b(input)
	if result != expect {
		t.Errorf("expected %d, got %d", expect, result)
	}
}
