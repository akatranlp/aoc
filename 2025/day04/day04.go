package main

import (
	"aoc/lib/aoc"
	"aoc/lib/map2d"
	"io"
	"maps"
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
	// count neighbours of all fields
	// then iterate again if neighbour has only 3, elimenate and decrement count of neighbours of neighbours

	neighborCount := make(map[map2d.Vector2]int)

	field := map2d.NewCellMap(r, map2d.CellMapFn)
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
		neighborCount[cell.ExtractCoords()] = paperRollCount
	}

	var res int
	for {
		newMap := maps.Collect(maps.All(neighborCount))

		var count int
		for k, v := range neighborCount {
			if v >= 4 {
				continue
			}
			count++
			delete(newMap, k)
			for _, dRow := range neighbors {
				for _, dCol := range neighbors {
					if dRow == 0 && dCol == 0 {
						continue
					}
					dir := map2d.NewVector2(dRow, dCol)
					newVec := k.Add(dir)
					if _, ok := newMap[newVec]; ok {
						newMap[newVec]--
					}
				}
			}
		}
		if count == 0 {
			break
		}
		res += count

		neighborCount = newMap
	}

	return res
}
