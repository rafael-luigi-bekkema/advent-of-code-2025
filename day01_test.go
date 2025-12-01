package main

import (
	"bufio"
	"fmt"
	"iter"
	"os"
	"strings"
	"testing"
)

const day1TestInput = `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`

func TestDay1a_Example(t *testing.T) {
	expected := 3
	if value := day1a(strings.SplitSeq(day1TestInput, "\n")); value != expected {
		t.Errorf("expected %d, got %d", expected, value)
	}
}

func TestDay1a(t *testing.T) {
	expected := 1031

	if value := day1a(readlines("./inputs/day01.txt")); value != expected {
		t.Errorf("expected %d, got %d", expected, value)
	}
}

func TestDay1b_Example(t *testing.T) {
	expected := 6
	if value := day1b(strings.SplitSeq(day1TestInput, "\n")); value != expected {
		t.Errorf("expected %d, got %d", expected, value)
	}
}

func TestDay1b(t *testing.T) {
	expected := 5831

	if value := day1b(readlines("./inputs/day01.txt")); value != expected {
		t.Errorf("expected %d, got %d", expected, value)
	}
}

func readlines(path string) iter.Seq[string] {
	return func(yield func(string) bool) {
		file, err := os.Open(path)
		if err != nil {
			panic(fmt.Sprintf("Failed to open: %s", err))
		}
		defer func() { _ = file.Close() }()

		rdr := bufio.NewScanner(file)

		for rdr.Scan() {
			if !yield(rdr.Text()) {
				break
			}
		}
	}
}
