package main

import (
	"aoc/lib/aoc"
	"aoc/lib/its"
	"aoc/lib/utils"
	"io"
	"math"
	"slices"
	"strings"
)

type Day04 struct{}

var _ aoc.Problem = (*Day04)(nil)

func (*Day04) Part1(r io.Reader) int {
	var res int
	for row := range its.Filter(
		its.ReaderToIter(r),
		its.FilterEmptyLines,
	) {
		numbers := slices.Collect(its.Map(
			strings.SplitSeq(strings.Split(row, ": ")[1], " | "),
			func(cards string) []int {
				return slices.Collect(its.Map(
					strings.FieldsSeq(cards),
					utils.MapStrToInt,
				),
				)
			},
		))
		winningNumbers := numbers[0]
		hand := numbers[1]
		var count int
		for _, num := range hand {
			if slices.Contains(winningNumbers, num) {
				count++
			}
		}
		res += int(math.Pow(2, float64(count-1)))
	}
	return res
}

func (*Day04) Part2(r io.Reader) int {
	type CardGame struct {
		count, times int
	}
	var cards []CardGame
	for row := range its.Filter(
		its.ReaderToIter(r),
		its.FilterEmptyLines,
	) {
		numbers := slices.Collect(its.Map(
			strings.SplitSeq(strings.Split(row, ": ")[1], " | "),
			func(cards string) []int {
				return slices.Collect(its.Map(
					strings.FieldsSeq(cards),
					utils.MapStrToInt,
				),
				)
			},
		))
		winningNumbers := numbers[0]
		hand := numbers[1]
		var count int
		for _, num := range hand {
			if slices.Contains(winningNumbers, num) {
				count++
			}
		}
		cards = append(cards, CardGame{count, 1})
	}

	for i, curr := range cards {
		for j := range curr.count {
			cards[i+j+1].times += curr.times
		}
	}
	var res int
	for _, curr := range cards {
		res += curr.times
	}
	return res
}
