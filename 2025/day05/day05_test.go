package main

import (
	"bytes"
	"testing"
)

var part1Test = `
3-5
10-14
16-20
12-18

1
5
8
11
17
32
`

func TestDay05(t *testing.T) {
	day05 := Day05{}
	t.Run("part 1", func(t *testing.T) {
		input := bytes.NewBufferString(part1Test)

		expected := 3
		actual := day05.Part1(input)

		if expected != actual {
			t.Fatalf("ERROR: expected %d actual %d\n", expected, actual)
		}
	})

	t.Run("part 2", func(t *testing.T) {
		input := bytes.NewBufferString(part1Test)

		expected := -1
		actual := day05.Part2(input)

		if expected != actual {
			t.Fatalf("ERROR: expected %d actual %d\n", expected, actual)
		}
	})
}
