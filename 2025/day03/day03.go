package main

import (
	"aoc/lib/aoc"
	"aoc/lib/its"
	"aoc/lib/utils"
	"io"
	"slices"
)

type Day03 struct{}

var _ aoc.Problem = (*Day03)(nil)

func (*Day03) Part1(r io.Reader) int {
	return its.Reduce(its.Filter(its.ReaderToIter(r), its.FilterEmptyLines), 0, func(acc int, row string) int {
		numbers := slices.Collect(its.Map(slices.Values([]byte(row)), utils.MapByteToInt))

		var maxIdx int
		for i := 1; i < len(numbers)-1; i++ {
			if numbers[i] > numbers[maxIdx] {
				maxIdx = i
			}
		}

		max2Idx := maxIdx + 1
		for i := max2Idx; i < len(numbers); i++ {
			if numbers[i] > numbers[max2Idx] {
				max2Idx = i
			}
		}

		maxNum := numbers[maxIdx]*10 + numbers[max2Idx]
		return acc + maxNum
	})
}

func (*Day03) Part2(r io.Reader) int {
	return its.Reduce(its.Filter(its.ReaderToIter(r), its.FilterEmptyLines), 0, func(acc int, row string) int {
		numbers := slices.Collect(its.Map(slices.Values([]byte(row)), utils.MapByteToInt))

		maxIdxs := make([]int, 12)
		for i := range maxIdxs {
			if i != 0 {
				maxIdxs[i] = maxIdxs[i-1] + 1
			}
			for j := maxIdxs[i]; j <= len(numbers)-(12-i); j++ {
				if numbers[j] > numbers[maxIdxs[i]] {
					maxIdxs[i] = j
				}
			}
		}
		maxNum := 0
		for i := range maxIdxs {
			maxNum = maxNum*10 + numbers[maxIdxs[i]]
		}
		return acc + maxNum
	})
}
