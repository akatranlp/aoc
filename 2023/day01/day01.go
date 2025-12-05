package main

import (
	"aoc/lib/aoc"
	"aoc/lib/its"
	"aoc/lib/utils"
	"io"
	"strings"
)

type Day01 struct{}

var _ aoc.Problem = (*Day01)(nil)

func (*Day01) Part1(r io.Reader) int {
	return its.Reduce(its.Filter(its.ReaderToIter(r), its.FilterEmptyLines), 0, func(acc int, row string) int {
		numbers := its.MapSlice([]byte(row), utils.MapByteToInt)
		numbers = its.FilterSlice(numbers, func(num int) bool {
			return num >= 0 && num <= 9
		})
		return acc + numbers[0]*10 + numbers[len(numbers)-1]
	})
}

var numString = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func (*Day01) Part2(r io.Reader) int {
	return its.Reduce(its.Filter(its.ReaderToIter(r), its.FilterEmptyLines), 0, func(acc int, row string) int {
		var first int
	outer1:
		for i := 0; i < len(row); i++ {
			c := row[i]
			num := utils.MapByteToInt(c)
			if num >= 0 && num <= 9 {
				first = num
				break outer1
			}

			subRow := row[i:]
			for k, v := range numString {
				if strings.HasPrefix(subRow, k) {
					first = v
					break outer1
				}
			}
		}
		var second int
	outer2:
		for i := len(row) - 1; i >= 0; i-- {
			num := utils.MapByteToInt(row[i])
			if num >= 0 && num <= 9 {
				second = num
				break outer2
			}

			subRow := row[:i+1]
			for k, v := range numString {
				if strings.HasSuffix(subRow, k) {
					second = v
					break outer2
				}
			}
		}

		return acc + first*10 + second
	})
}
