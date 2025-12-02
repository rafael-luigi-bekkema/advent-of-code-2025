package main

import "testing"

func TestDay2a_Example(t *testing.T) {
	input := "11-22,95-115,998-1012,1188511880-1188511890,222220-222224," +
		"1698522-1698528,446443-446449,38593856-38593862,565653-565659," +
		"824824821-824824827,2121212118-2121212124"
	expect := 1227775554
	result := day2a(input)
	if result != expect {
		t.Errorf("Expected %d, got %d", expect, result)
	}
}

func TestDay2a(t *testing.T) {
	input := readfile("inputs/day02.txt")
	expect := 38158151648
	result := day2a(input)
	if result != expect {
		t.Errorf("Expected %d, got %d", expect, result)
	}
}

func TestDay2b_Example(t *testing.T) {
	input := "11-22,95-115,998-1012,1188511880-1188511890,222220-222224," +
		"1698522-1698528,446443-446449,38593856-38593862,565653-565659," +
		"824824821-824824827,2121212118-2121212124"
	expect := 4174379265
	result := day2b(input)
	if result != expect {
		t.Errorf("Expected %d, got %d", expect, result)
	}
}

func TestDay2b(t *testing.T) {
	input := readfile("inputs/day02.txt")
	expect := 45283684555
	result := day2b(input)
	if result != expect {
		t.Errorf("Expected %d, got %d", expect, result)
	}
}
