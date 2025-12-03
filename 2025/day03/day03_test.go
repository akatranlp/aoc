package main

import (
	"bytes"
	"testing"
)

var part1Test = `
987654321111111
811111111111119
234234234234278
818181911112111
`

func TestDay03(t *testing.T) {
	day03 := Day03{}
	t.Run("part 1", func(t *testing.T) {
		input := bytes.NewBufferString(part1Test)

		expected := 357
		actual := day03.Part1(input)

		if expected != actual {
			t.Fatalf("ERROR: expected %d actual %d\n", expected, actual)
		}
	})

	t.Run("part 2", func(t *testing.T) {
		input := bytes.NewBufferString(part1Test)

		expected := 3121910778619
		actual := day03.Part2(input)

		if expected != actual {
			t.Fatalf("ERROR: expected %d actual %d\n", expected, actual)
		}
	})
}
