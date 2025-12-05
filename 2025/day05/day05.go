package main

import (
	"aoc/lib/aoc"
	"aoc/lib/its"
	"aoc/lib/utils"
	"io"
	"iter"
	"slices"
	"strconv"
	"strings"
)

type Day05 struct{}

var _ aoc.Problem = (*Day05)(nil)

type Range struct {
	down  int
	up    int
	valid bool
}

func (r *Range) InRange(v int) bool {
	return v >= r.down && v <= r.up
}

func (r *Range) CombineRanges(v *Range) *Range {
	if v.down <= r.down && v.up >= r.down ||
		r.down <= v.down && r.up >= v.down ||
		v.up >= r.up && v.down <= r.up ||
		r.up >= v.up && r.down <= v.up {
		r.down = min(r.down, v.down)
		r.up = max(r.up, v.up)
		v.valid = false
	}
	return v
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
					true,
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
	var ranges []*Range
	for i, block := range its.Enumerate(its.ReaderToIter(r, its.SplitByBlocks)) {
		switch i {
		case 0:
			seq := strings.SplitSeq(block, "\n")
			seq = its.Filter(seq, its.FilterEmptyLines)
			for row := range seq {
				nums := strings.Split(row, "-")
				newRange := &Range{
					utils.Must(strconv.Atoi(nums[0])),
					utils.Must(strconv.Atoi(nums[1])),
					true,
				}
				for _, r := range ranges {
					newRange = r.CombineRanges(newRange)
					if !newRange.valid {
						break
					}
				}
				if newRange.valid {
					ranges = append(ranges, newRange)
				}
			}
		}
	}

	changed := true
	for changed {
		changed = false
		for _, r := range ranges {
			if !r.valid {
				continue
			}
			for _, or := range ranges {
				if !or.valid || r == or {
					continue
				}
				if !r.CombineRanges(or).valid {
					changed = true
				}
			}
		}
	}

	var res int
	for _, r := range ranges {
		if !r.valid {
			continue
		}
		res += r.up - r.down + 1
	}
	return res
}
