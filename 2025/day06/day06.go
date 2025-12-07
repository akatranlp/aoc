package main

import (
	"aoc/lib/aoc"
	"aoc/lib/its"
	"aoc/lib/utils"
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
	input := slices.Collect(
		its.Filter(
			its.ReaderToIter(r),
			its.FilterEmptyLines,
		),
	)
	rows := len(input)
	cols := len(input[0])

	var operations [][]string
	var currentIdx int
	for x := range cols {
		var someChars bool
		for y := range rows - 1 {
			if input[y][x] != ' ' {
				someChars = true
			}
		}
		if !someChars || x == cols-1 {
			var matrix []string
			if x == cols-1 {
				x = cols
			}

			for y := range rows {
				matrix = append(matrix, input[y][currentIdx:x])
			}
			currentIdx = x + 1
			operations = append(operations, matrix)
		}
	}

	var res int
	for _, op := range operations {
		operation := strings.TrimSpace(op[rows-1])
		var sum int
		if operation == "*" {
			sum = 1
		}
		numLen := len(op[0])
		for x := numLen - 1; x >= 0; x-- {
			var numStr []byte
			for y := range rows - 1 {
				if op[y][x] != ' ' {
					numStr = append(numStr, op[y][x])
				}
			}
			num := utils.MapStrToInt(string(numStr))
			switch operation {
			case "+":
				sum += num
			case "*":
				sum *= num
			}
		}
		res += sum
	}

	return res
}
