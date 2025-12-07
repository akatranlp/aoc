package main

import (
	"bytes"
	"testing"
)

var part1Test = `
.......S.......
...............
.......^.......
...............
......^.^......
...............
.....^.^.^.....
...............
....^.^...^....
...............
...^.^...^.^...
...............
..^...^.....^..
...............
.^.^.^.^.^...^.
...............
`

func TestDay07(t *testing.T) {
	day07 := Day07{}
	t.Run("part 1", func(t *testing.T) {
		input := bytes.NewBufferString(part1Test)

		expected := 21
		actual := day07.Part1(input)

		if expected != actual {
			t.Fatalf("ERROR: expected %d actual %d\n", expected, actual)
		}
	})

	t.Run("part 2", func(t *testing.T) {
		input := bytes.NewBufferString(part1Test)

		expected := -1
		actual := day07.Part2(input)

		if expected != actual {
			t.Fatalf("ERROR: expected %d actual %d\n", expected, actual)
		}
	})
}
