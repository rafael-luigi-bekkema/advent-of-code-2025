package main

import (
	"strings"
	"testing"
)

const day8TestInput = `162,817,812
57,618,57
906,360,560
592,479,940
352,342,300
466,668,158
542,29,236
431,825,988
739,650,466
52,470,668
216,146,977
819,987,18
117,168,530
805,96,715
346,949,466
970,615,88
941,993,340
862,61,35
984,92,344
425,690,689`

func TestDay8a_Example(t *testing.T) {
	expect := 40
	result := day8a(strings.SplitSeq(day8TestInput, "\n"), 10)

	if result != expect {
		t.Errorf("expected %d, got %d", expect, result)
	}
}

func TestDay8a(t *testing.T) {
	input := readlines("./inputs/day08.txt")
	expect := 96672
	result := day8a(input, 1000)

	if result != expect {
		t.Errorf("expected %d, got %d", expect, result)
	}
}

func TestDay8b_Example(t *testing.T) {
	expect := 25272
	result := day8b(strings.SplitSeq(day8TestInput, "\n"))

	if result != expect {
		t.Errorf("expected %d, got %d", expect, result)
	}
}

func TestDay8b(t *testing.T) {
	input := readlines("./inputs/day08.txt")
	expect := 22517595
	result := day8b(input)

	if result != expect {
		t.Errorf("expected %d, got %d", expect, result)
	}
}
