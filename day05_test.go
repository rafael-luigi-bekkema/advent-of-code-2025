package main

import (
	"strings"
	"testing"
)

var day5aTestInput = `3-5
10-14
16-20
12-18

1
5
8
11
17
32`

func TestDay5a_Example(t *testing.T) {
	expect := 3
	result := day5a(strings.SplitSeq(day5aTestInput, "\n"))
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

func TestDay5b_Example(t *testing.T) {
	tt := []struct {
		input  string
		expect int
	}{
		{day5aTestInput, 14},
		{"3-5\n3-5", 3},
		{"3-5\n3-6", 4},
		{"7-15\n3-5\n5-8", 13},
		{"1-4\n2-3", 4},
		{"1-4\n2-4", 4},
		{"2-3\n1-4", 4},
	}
	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			result := day5b(strings.SplitSeq(tc.input, "\n"))
			if result != tc.expect {
				t.Errorf("expected %d, got %d", tc.expect, result)
			}
		})
	}
}

func TestDay5b(t *testing.T) {
	input := readlines("./inputs/day05.txt")
	expect := 357907198933892
	result := day5b(input)
	if result != expect {
		t.Errorf("expected %d, got %d", expect, result)
	}
}
