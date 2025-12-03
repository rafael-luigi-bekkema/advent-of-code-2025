package main

import (
	"bufio"
	"bytes"
	"fmt"
	"iter"
	"os"
	"strconv"
)

func readfile(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(bytes.TrimSpace(data))
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

func strtoint(input string) int {
	val, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return val
}

func mkslice[T any](size int, init T) []T {
	result := make([]T, size)
	for i := range size {
		result[i] = init
	}
	return result
}
