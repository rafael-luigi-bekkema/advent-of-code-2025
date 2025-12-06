package main

import (
	"strings"
	"testing"
)

const day6TestInput = `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +`

func TestDay6a_Example(t *testing.T) {
	expect := 4277556
	result := day6a(strings.SplitSeq(day6TestInput, "\n"))
	if result != expect {
		t.Errorf("expect %d, got %d", expect, result)
	}
}

func TestDay6a(t *testing.T) {
	input := readlines("./inputs/day06.txt")
	expect := 6605396225322
	result := day6a(input)
	if result != expect {
		t.Errorf("expect %d, got %d", expect, result)
	}
}

func TestDay6b_Example(t *testing.T) {
	expect := 3263827
	result := day6b(strings.SplitSeq(day6TestInput, "\n"))
	if result != expect {
		t.Errorf("expect %d, got %d", expect, result)
	}
}

func TestDay6b(t *testing.T) {
	input := readlines("./inputs/day06.txt")
	expect := 11052310600986
	result := day6b(input)
	if result != expect {
		t.Errorf("expect %d, got %d", expect, result)
	}
}
