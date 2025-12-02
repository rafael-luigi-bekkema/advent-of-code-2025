package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day2a(input string) (total int) {
	for idr := range strings.SplitSeq(input, ",") {
		var from, to int
		_, _ = fmt.Sscanf(idr, "%d-%d", &from, &to)
		if to < from {
			panic("oops")
		}
		for i := from; i <= to; i++ {
			sn := strconv.Itoa(i)
			if sn[:len(sn)/2] == sn[len(sn)/2:] {
				total += i
			}
		}
	}

	return
}

func day2b(input string) (total int) {
	for idr := range strings.SplitSeq(input, ",") {
		var from, to int
		_, _ = fmt.Sscanf(idr, "%d-%d", &from, &to)
		if to < from {
			panic("oops")
		}
		fmt.Printf("start %q ", idr)
		for i := from; i <= to; i++ {
			sn := strconv.Itoa(i)
			for j := 0; j < len(sn)/2; j++ {
				if len(sn)%(j+1) != 0 {
					continue
				}
				all := true
				for k := 0; k < len(sn)/(j+1)-1; k++ {
					if !strings.HasPrefix(sn[(k+1)*(j+1):], sn[:j+1]) {
						all = false
						break
					}
				}
				if all {
					fmt.Print(sn, " ")
					total += i
					break
				}
			}
		}
		fmt.Println("done")
	}

	return
}
