package main

import (
	"aoc/lib/aoc"
	"aoc/lib/map2d"
	"fmt"
	"io"
)

type Day04 struct{}

var _ aoc.Problem = (*Day04)(nil)

var neighbors = []int{-1, 0, 1}

func (*Day04) Part1(r io.Reader) int {
	field := map2d.NewCellMap(r, map2d.CellMapFn)
	var res int

	for cell := range field.Iter() {
		if cell.Value != '@' {
			continue
		}
		var paperRollCount int
		for _, dRow := range neighbors {
			for _, dCol := range neighbors {
				if dRow == 0 && dCol == 0 {
					continue
				}
				dir := map2d.NewVector2(dRow, dCol)
				newVec := cell.ExtractCoords().Add(dir)
				// fmt.Printf("point: %+v, dir: %+v, newDir: %+v", cell, dir, newVec)
				if !field.InBounce(newVec) {
					// fmt.Println("Not in bounce")
					continue
				}
				// fmt.Printf(", look: %+v ->", field.Get(newVec))
				if field.Get(newVec).Value == '@' {
					// fmt.Println("HIT")
					paperRollCount++
				} else {
					// fmt.Println("MISS")
				}
			}
		}
		if paperRollCount < 4 {
			// fmt.Printf("point: %+v has less papers: %d\n", cell, paperRollCount)
			res++
		}
	}

	return res
}

func (*Day04) Part2(r io.Reader) int {
	fmt.Println("Part2 not implemented")
	return -1
}
