package main

import (
	"aoc-lib/aoc"
	"aoc-lib/its"
	"aoc-lib/utils"
	"fmt"
	"io"
)

type Day01 struct{}

var _ aoc.Problem = (*Day01)(nil)

func (*Day01) Part1(r io.Reader) int {
	dial := 50
	var res int
	for row := range its.Filter(its.ReaderToIter(r), its.FilterEmptyLines) {
		var direction rune
		var length int
		fmt.Sscanf(row, "%c%d", &direction, &length)

		switch direction {
		case 'R':
			dial = utils.IntMod(dial+length, 100)
		case 'L':
			dial = utils.IntMod(dial-length, 100)
		default:
			panic("unreachable")
		}

		if dial == 0 {
			res++
		}
	}
	return res
}

func (*Day01) Part2(r io.Reader) int {
	fmt.Println("Part2 not implemented")
	return -1
}
