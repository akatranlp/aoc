package main

import (
	"bytes"
	"testing"
)

var part1Test = `
..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.
`

func TestDay04(t *testing.T) {
	day04 := Day04{}
	t.Run("part 1", func(t *testing.T) {
		input := bytes.NewBufferString(part1Test)

		expected := 13
		actual := day04.Part1(input)

		if expected != actual {
			t.Fatalf("ERROR: expected %d actual %d\n", expected, actual)
		}
	})

	t.Run("part 2", func(t *testing.T) {
		input := bytes.NewBufferString(part1Test)

		expected := 43
		actual := day04.Part2(input)

		if expected != actual {
			t.Fatalf("ERROR: expected %d actual %d\n", expected, actual)
		}
	})
}
