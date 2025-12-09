package main

import (
	"aoc/lib/aoc"
	"aoc/lib/its"
	"aoc/lib/map2d"
	"aoc/lib/utils"
	"fmt"
	"io"
	"slices"
	"strings"
)

type Day09 struct{}

var _ aoc.Problem = (*Day09)(nil)

type Entry struct {
	LIdx, RIdx, area int
}

func (*Day09) Part1(r io.Reader) int {
	points := slices.Collect(its.Map(
		its.Filter(
			its.ReaderToIter(r),
			its.FilterEmptyLines),
		func(row string) map2d.Vector2 {
			nums := strings.Split(row, ",")
			return map2d.NewVector2(utils.MapStrToInt(nums[0]), utils.MapStrToInt(nums[1]))
		},
	))

	areas := slices.Collect(its.Map(its.AllCombinationsWithIdx(points, false), func(v its.CombinationWithIdx[map2d.Vector2]) Entry {
		return Entry{v.LIdx, v.RIdx, (utils.Abs(v.L.X, v.R.X) + 1) * (utils.Abs(v.L.Y, v.R.Y) + 1)}
	}))
	slices.SortFunc(areas, func(a, b Entry) int { return b.area - a.area })
	return areas[0].area
}

func (*Day09) Part2(r io.Reader) int {
	fmt.Println("Part2 not implemented")
	return -1
}
