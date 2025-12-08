package main

import (
	"aoc/lib/aoc"
	"aoc/lib/its"
	aocSlices "aoc/lib/slices"
	"aoc/lib/utils"
	"io"
	"slices"
	"strings"
)

type Day08 struct{}

var _ aoc.Problem = (*Day08)(nil)

type Vec3 struct {
	x, y, z int
}

func (v Vec3) sqDistance(o Vec3) int {
	dx := v.x - o.x
	dy := v.y - o.y
	dz := v.z - o.z
	return dx*dx + dy*dy + dz*dz

}

type Entry struct {
	LIdx, RIdx, distance int
}

func InCircuit(circuits [][]int, idx int) int {
	for i, c := range circuits {
		if slices.Contains(c, idx) {
			return i
		}
	}
	return -1
}

func (d *Day08) Part1(r io.Reader) int {
	return d.Part1Impl(r, false)
}

func (*Day08) Part1Impl(r io.Reader, test bool) int {
	coords := slices.Collect(its.Map(its.Filter(
		its.ReaderToIter(r),
		its.FilterEmptyLines,
	), func(row string) Vec3 {
		coords := slices.Collect(its.Map(strings.SplitSeq(row, ","), utils.MapStrToInt))
		return Vec3{coords[0], coords[1], coords[2]}
	}))

	distances := slices.Collect(
		its.Map(its.AllCombinationsWithIdx(coords, false), func(v its.CombinationWithIdx[Vec3]) Entry {
			return Entry{v.LIdx, v.RIdx, v.L.sqDistance(v.R)}
		}))
	slices.SortFunc(distances, func(a, b Entry) int { return a.distance - b.distance })

	remaining := aocSlices.NewSetFromIter(its.Range(len(coords)))
	var circuits [][]int

	count := 1000
	if test {
		count = 10
	}

	for _, d := range distances {
		if len(remaining) == 0 || count == 0 {
			break
		}
		LCIdx, RCIdx := -1, -1
		if !remaining.Has(d.LIdx) {
			LCIdx = InCircuit(circuits, d.LIdx)
		}
		if !remaining.Has(d.RIdx) {
			RCIdx = InCircuit(circuits, d.RIdx)
		}
		if LCIdx != -1 && RCIdx != -1 {
			if LCIdx == RCIdx {
				count--
				continue
			}
			rc := circuits[RCIdx]
			circuits[LCIdx] = append(circuits[LCIdx], rc...)
			circuits = its.RemoveIndex(circuits, RCIdx)
			count--
		} else if LCIdx == -1 && RCIdx == -1 {
			delete(remaining, d.LIdx)
			delete(remaining, d.RIdx)
			count--
			circuits = append(circuits, []int{d.LIdx, d.RIdx})
		} else if LCIdx == -1 {
			delete(remaining, d.LIdx)
			count--
			circuits[RCIdx] = append(circuits[RCIdx], d.LIdx)
		} else if RCIdx == -1 {
			delete(remaining, d.RIdx)
			count--
			circuits[LCIdx] = append(circuits[LCIdx], d.RIdx)
		}
	}

	res := 1
	circuitLens := its.MapSlice(circuits, func(c []int) int { return len(c) })
	slices.SortFunc(circuitLens, func(a, b int) int { return b - a })
	for i, c := range circuitLens {
		if i > 2 {
			break
		}
		res *= c
	}

	return res
}

func (*Day08) Part2(r io.Reader) int {
	coords := slices.Collect(its.Map(its.Filter(
		its.ReaderToIter(r),
		its.FilterEmptyLines,
	), func(row string) Vec3 {
		coords := slices.Collect(its.Map(strings.SplitSeq(row, ","), utils.MapStrToInt))
		return Vec3{coords[0], coords[1], coords[2]}
	}))

	distances := slices.Collect(
		its.Map(its.AllCombinationsWithIdx(coords, false), func(v its.CombinationWithIdx[Vec3]) Entry {
			return Entry{v.LIdx, v.RIdx, v.L.sqDistance(v.R)}
		}))
	slices.SortFunc(distances, func(a, b Entry) int { return a.distance - b.distance })

	remaining := aocSlices.NewSetFromIter(its.Range(len(coords)))
	var circuits [][]int

	var lastLIdx, lastRIdx int
	for _, d := range distances {
		if len(remaining) == 0 {
			break
		}
		LCIdx, RCIdx := -1, -1
		if !remaining.Has(d.LIdx) {
			LCIdx = InCircuit(circuits, d.LIdx)
		}
		if !remaining.Has(d.RIdx) {
			RCIdx = InCircuit(circuits, d.RIdx)
		}
		if LCIdx != -1 && RCIdx != -1 {
			if LCIdx == RCIdx {
				continue
			}
			lastLIdx, lastRIdx = d.LIdx, d.RIdx
			rc := circuits[RCIdx]
			circuits[LCIdx] = append(circuits[LCIdx], rc...)
			circuits = its.RemoveIndex(circuits, RCIdx)
		} else if LCIdx == -1 && RCIdx == -1 {
			delete(remaining, d.LIdx)
			delete(remaining, d.RIdx)
			lastLIdx, lastRIdx = d.LIdx, d.RIdx
			circuits = append(circuits, []int{d.LIdx, d.RIdx})
		} else if LCIdx == -1 {
			delete(remaining, d.LIdx)
			lastLIdx, lastRIdx = d.LIdx, d.RIdx
			circuits[RCIdx] = append(circuits[RCIdx], d.LIdx)
		} else if RCIdx == -1 {
			delete(remaining, d.RIdx)
			lastLIdx, lastRIdx = d.LIdx, d.RIdx
			circuits[LCIdx] = append(circuits[LCIdx], d.RIdx)
		}
	}

	return coords[lastLIdx].x * coords[lastRIdx].x
}
