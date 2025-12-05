package main

import (
	"aoc/lib/aoc"
	"aoc/lib/its"
	"aoc/lib/utils"
	"fmt"
	"io"
	"iter"
	"slices"
	"strconv"
	"strings"
)

type Day05 struct{}

var _ aoc.Problem = (*Day05)(nil)

type Range struct {
	up   int
	down int
}

func (r Range) InRange(v int) bool {
	return v >= r.up && v <= r.down
}

func (*Day05) Part1(r io.Reader) int {
	var ranges []Range
	var ids iter.Seq[int]
	for i, block := range its.Enumerate(its.ReaderToIter(r, its.SplitByBlocks)) {
		switch i {
		case 0:
			seq := strings.SplitSeq(block, "\n")
			seq = its.Filter(seq, its.FilterEmptyLines)
			seq2 := its.Map(seq, func(row string) Range {
				nums := strings.Split(row, "-")
				return Range{
					utils.Must(strconv.Atoi(nums[0])),
					utils.Must(strconv.Atoi(nums[1])),
				}
			})
			ranges = slices.Collect(seq2)
		case 1:
			seq := strings.SplitSeq(block, "\n")
			seq = its.Filter(seq, its.FilterEmptyLines)
			ids = its.Map(seq, utils.MapStrToInt)
		}
	}
	return its.Reduce(ids, 0, func(acc, id int) int {
		for _, r := range ranges {
			if r.InRange(id) {
				return acc + 1
			}
		}
		return acc
	})
}

func (*Day05) Part2(r io.Reader) int {
	fmt.Println("Part2 not implemented")
	return -1
}
