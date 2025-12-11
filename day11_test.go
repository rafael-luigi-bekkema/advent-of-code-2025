package main

import (
	"strings"
	"testing"
)

const day11TestAInput = `aaa: you hhh
you: bbb ccc
bbb: ddd eee
ccc: ddd eee fff
ddd: ggg
eee: out
fff: out
ggg: out
hhh: ccc fff iii
iii: out`

func TestDay11a_Example(t *testing.T) {
	expect := 5
	result := day11a(strings.SplitSeq(day11TestAInput, "\n"))
	if result != expect {
		t.Errorf("expected %d, got %d", expect, result)
	}
}

func TestDay11a(t *testing.T) {
	expect := 636
	result := day11a(readlines("./inputs/day11.txt"))
	if result != expect {
		t.Errorf("expected %d, got %d", expect, result)
	}
}

// const day11TestBInput = `svr: aaa bbb
// aaa: fft
// fft: ccc
// bbb: tty
// tty: ccc
// ccc: ddd eee
// ddd: hub
// hub: fff
// eee: dac
// dac: fff
// fff: ggg hhh
// ggg: out
// hhh: out`
//
// func TestDay11b_Example(t *testing.T) {
// 	expect := 2
// 	result := day11b(strings.SplitSeq(day11TestBInput, "\n"))
// 	if result != expect {
// 		t.Errorf("expected %d, got %d", expect, result)
// 	}
// }
//
// func TestDay11b(t *testing.T) {
// 	expect := 0
// 	result := day11b(readlines("./inputs/day11.txt"))
// 	if result != expect {
// 		t.Errorf("expected %d, got %d", expect, result)
// 	}
// }
