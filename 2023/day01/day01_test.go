package main

import (
	"bytes"
	"testing"
)

var part1Test = `
1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
`

var part2Test = `
two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
`

func TestDay01(t *testing.T) {
	day01 := Day01{}
	t.Run("part 1", func(t *testing.T) {
		input := bytes.NewBufferString(part1Test)

		expected := 142
		actual := day01.Part1(input)

		if expected != actual {
			t.Fatalf("ERROR: expected %d actual %d\n", expected, actual)
		}
	})

	t.Run("part 2", func(t *testing.T) {
		input := bytes.NewBufferString(part2Test)

		expected := 281
		actual := day01.Part2(input)

		if expected != actual {
			t.Fatalf("ERROR: expected %d actual %d\n", expected, actual)
		}
	})
}
