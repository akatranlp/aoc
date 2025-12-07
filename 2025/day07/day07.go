package main

import (
	"aoc/lib/aoc"
	"aoc/lib/map2d"
	aocSlices "aoc/lib/slices"
	"io"
	"maps"
	"slices"
)

type Day07 struct{}

var _ aoc.Problem = (*Day07)(nil)

func (*Day07) Part1(r io.Reader) int {
	inputMap := map2d.NewCellMap(r, map2d.CellMapFn)

	beams := aocSlices.NewSet[int]()
	for col := range inputMap.Cols {
		vec := map2d.NewVector2(col, 0)
		cell := inputMap.Get(vec)
		if cell.Value == 'S' {
			beams.Set(col)
			break
		}
	}
	var count int
	for row := range inputMap.Rows {
		for _, b := range slices.Collect(maps.Keys(beams)) {
			vec := map2d.NewVector2(b, row)
			cell := inputMap.Get(vec)
			if cell.Value == '^' {
				count++
				delete(beams, b)
				beams.Set(b - 1)
				beams.Set(b + 1)
			}
		}
	}

	return count
}

func (*Day07) Part2(r io.Reader) int {
	inputMap := map2d.NewCellMap(r, map2d.CellMapFn)

	beams := make(map[int]int)
	for col := range inputMap.Cols {
		vec := map2d.NewVector2(col, 0)
		cell := inputMap.Get(vec)
		if cell.Value == 'S' {
			beams[col] = 1
			break
		}
	}
	sum := 1
	for row := range inputMap.Rows {
		for b, count := range maps.Collect(maps.All(beams)) {
			if count == 0 {
				continue
			}
			vec := map2d.NewVector2(b, row)
			cell := inputMap.Get(vec)
			if cell.Value == '^' {
				sum += count
				beams[b] -= count
				beams[b-1] += count
				beams[b+1] += count
			}
		}
	}

	return sum
}
