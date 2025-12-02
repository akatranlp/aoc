package main

import (
	"aoc-lib/aoc"
	"aoc-lib/its"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Day02 struct{}

var _ aoc.Problem = (*Day02)(nil)

func (*Day02) Part1(r io.Reader) int {
	var res int
	for field := range its.ReaderToIter(r, its.SplitByByte(',')) {
		field = strings.TrimSpace(field)
		var l, r int
		fmt.Sscanf(field, "%d-%d", &l, &r)
		for i := l; i <= r; i++ {
			number := strconv.Itoa(i)
			numLen := len(number)
			if numLen%2 == 0 {
				part1 := number[:numLen/2]
				part2 := number[numLen/2:]
				if part1 == part2 {
					res += i
				}
			}
		}
	}
	return res
}

func (*Day02) Part2(r io.Reader) int {
	fmt.Println("Part2 not implemented")
	return -1
}
