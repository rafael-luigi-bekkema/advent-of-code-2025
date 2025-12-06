package main

import (
	"bytes"
	"iter"
	"strings"
)

func sumslice(s []int) (sum int) {
	for _, n := range s {
		sum += n
	}

	return
}

func mulslice(s []int) int {
	mul := s[0]

	for _, n := range s[1:] {
		mul *= n
	}

	return mul
}

func day6a(input iter.Seq[string]) (total int) {
	var calcs [][]int

	for line := range input {
		fields := strings.Fields(line)
		if fields[0] == "*" || fields[0] == "+" {
			for i, field := range fields {
				switch field {
				case "*":
					total += mulslice(calcs[i])
				case "+":
					total += sumslice(calcs[i])
				}
			}
			break
		}
		if calcs == nil {
			calcs = make([][]int, len(fields))
		}
		for i, field := range fields {
			calcs[i] = append(calcs[i], strtoint(field))
		}
	}

	return
}

func day6b(input iter.Seq[string]) (total int) {
	var lines [][]byte

	for line := range input {
		lines = append(lines, append([]byte(line), ' '))
	}

	sumline := bytes.Fields(lines[len(lines)-1])
	lines = lines[:len(lines)-1]

	for _, op := range sumline {
		var max int
		for _, line := range lines {
			i := 0
			for ; line[i] == ' '; i++ {
			}
			if idx := bytes.Index(line[i:], []byte{' '}) + i; max == 0 || idx > max {
				max = idx
			}
		}

		nums := make([][]byte, max)
		for i := range lines {
			s := lines[i][:max]
			lines[i] = lines[i][max+1:]

			for j := range max {
				nums[j] = append(nums[j], s[j])
			}
		}

		var subtotal int
		for _, num := range nums {
			val := strtoint(string(bytes.TrimSpace(num)))

			switch op[0] {
			case '*':
				if subtotal == 0 {
					subtotal = val
				} else {
					subtotal *= val
				}
			case '+':
				subtotal += val
			}
		}

		total += subtotal
	}

	return
}
