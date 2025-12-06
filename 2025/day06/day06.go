package main

import (
	"aoc/lib/aoc"
	"aoc/lib/its"
	"aoc/lib/utils"
	"fmt"
	"io"
	"slices"
	"strings"
)

type Day06 struct{}

var _ aoc.Problem = (*Day06)(nil)

func (*Day06) Part1(r io.Reader) int {
	type Operation struct {
		numbers   []int
		operation string
	}
	rows := slices.Collect(
		its.Map(
			its.Filter(its.ReaderToIter(r), its.FilterEmptyLines),
			func(row string) []string { return strings.Fields(row) }),
	)
	rowsLen := len(rows)
	operationsLen := len(rows[0])
	operations := make([]Operation, operationsLen)

	for x := range operationsLen {
		for y := range rowsLen {
			if y == rowsLen-1 {
				operations[x].operation = rows[y][x]
				continue
			}
			operations[x].numbers = append(operations[x].numbers, utils.MapStrToInt(rows[y][x]))
		}
	}

	var res int

	for _, op := range operations {
		var temp int
		switch op.operation {
		case "+":
			for _, num := range op.numbers {
				temp += num
			}
		case "*":
			temp = 1
			for _, num := range op.numbers {
				temp *= num
			}
		}
		res += temp
	}

	return res
}

func (*Day06) Part2(r io.Reader) int {
	fmt.Println("Part2 not implemented")
	return -1
}
