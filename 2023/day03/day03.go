package main

import (
	"aoc/lib/aoc"
	"aoc/lib/its"
	"aoc/lib/map2d"
	aocSlices "aoc/lib/slices"
	"aoc/lib/utils"
	"io"
	"slices"
)

type Day03 struct{}

var _ aoc.Problem = (*Day03)(nil)

var neighbors = slices.Collect(its.Matrix([]int{-1, 0, 1}))

type Number struct {
	row, start, end, number int
}

func (*Day03) Part1(r io.Reader) int {
	set := aocSlices.NewSet[map2d.Vector2]()
	var numbers []Number
	for y, row := range its.Enumerate(its.Filter(its.ReaderToIter(r), its.FilterEmptyLines)) {
		rowLen := len(row)
		startIndex := rowLen
		for x, c := range []byte(row) {
			if utils.IsDigit(c) {
				startIndex = min(startIndex, x)
			}
			if x == rowLen-1 || !utils.IsDigit(c) {
				endIndex := x
				if x == rowLen-1 && utils.IsDigit(c) {
					endIndex += 1
				}
				if startIndex < endIndex {
					numbers = append(numbers, Number{y, startIndex, endIndex, utils.MapStrToInt(row[startIndex:endIndex])})
				}

				startIndex = rowLen
			}
			if !utils.IsDigit(c) && c != '.' {
				set.Set(map2d.NewVector2(x, y))
			}
		}
	}

	var res int
	for _, number := range numbers {
	number:
		for x := range its.RangeFromTo(number.start, number.end) {
			for _, n := range neighbors {
				if set.Has(map2d.NewVector2(x+n.L, number.row+n.R)) {
					res += number.number
					break number
				}
			}
		}
	}

	return res
}

func (*Day03) Part2(r io.Reader) int {
	set := aocSlices.NewSet[map2d.Vector2]()
	var numbers []Number
	for y, row := range its.Enumerate(its.Filter(its.ReaderToIter(r), its.FilterEmptyLines)) {
		rowLen := len(row)
		startIndex := rowLen
		for x, c := range []byte(row) {
			if utils.IsDigit(c) {
				startIndex = min(startIndex, x)
			}
			if x == rowLen-1 || !utils.IsDigit(c) {
				endIndex := x
				if x == rowLen-1 && utils.IsDigit(c) {
					endIndex += 1
				}
				if startIndex < endIndex {
					numbers = append(numbers, Number{y, startIndex, endIndex, utils.MapStrToInt(row[startIndex:endIndex])})
				}

				startIndex = rowLen
			}
			if c == '*' {
				set.Set(map2d.NewVector2(x, y))
			}
		}
	}

	var res int

	for symbolVec := range set {
		var neighborNumbers []Number

		for _, n := range neighbors {
			coord := symbolVec.Add(map2d.NewVector2(n.L, n.R))

			for _, number := range numbers {
				if coord.Y != number.row {
					continue
				}

				if coord.X >= number.start && coord.X < number.end {
					if !slices.Contains(neighborNumbers, number) {
						neighborNumbers = append(neighborNumbers, number)
					}
				}
			}
		}
		if len(neighborNumbers) == 2 {
			res += neighborNumbers[0].number * neighborNumbers[1].number
		}
	}

	return res
}
