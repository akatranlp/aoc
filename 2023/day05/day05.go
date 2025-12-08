package main

import (
	"aoc/lib/aoc"
	"aoc/lib/its"
	"aoc/lib/ranges"
	"aoc/lib/utils"
	"fmt"
	"io"
	"slices"
	"strings"
)

type Day05 struct{}

var _ aoc.Problem = (*Day05)(nil)

type Mapping struct {
	src ranges.Range
	dst ranges.Range
}

func (*Day05) Part1(r io.Reader) int {
	var nums []int
	var mappings [][]Mapping
	for i, block := range its.Enumerate(its.ReaderToIter(r, its.SplitByBlocks)) {
		if i == 0 {
			nums = slices.Collect(
				its.Map(
					strings.FieldsSeq(strings.Split(block, ":")[1]),
					utils.MapStrToInt,
				),
			)
			continue
		}
		mappings = append(mappings, slices.Collect(its.Map(its.Skip(
			its.Filter(
				strings.SplitSeq(block, "\n"),
				its.FilterEmptyLines,
			),
			1,
		), func(row string) Mapping {
			numbers := slices.Collect(
				its.Map(
					strings.FieldsSeq(row),
					utils.MapStrToInt,
				),
			)
			srcRange := ranges.NewRangeCount(numbers[1], numbers[2])
			dstRange := ranges.NewRangeCount(numbers[0], numbers[2])
			return Mapping{srcRange, dstRange}
		})))
	}

	for _, mapping := range mappings {
		for j := range len(nums) {
			num := nums[j]
			for _, m := range mapping {
				if m.src.InRange(num) {
					nums[j] = m.dst.Down + num - m.src.Down
				}
			}
		}
	}
	res := 1 << 62
	for _, num := range nums {
		res = min(res, num)
	}
	return res
}

func (*Day05) Part2(r io.Reader) int {
	fmt.Println("Part2 not implemented")
	return -1
}
