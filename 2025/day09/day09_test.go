package main

import (
	"bytes"
	"testing"
)

var part1Test = `
7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3
`

func TestDay09(t *testing.T) {
	day09 := Day09{}
	t.Run("part 1", func(t *testing.T) {
		input := bytes.NewBufferString(part1Test)

		expected := 50
		actual := day09.Part1(input)

		if expected != actual {
			t.Fatalf("ERROR: expected %d actual %d\n", expected, actual)
		}
	})

	t.Run("part 2", func(t *testing.T) {
		input := bytes.NewBufferString(part1Test)

		expected := 24
		actual := day09.Part2(input)

		if expected != actual {
			t.Fatalf("ERROR: expected %d actual %d\n", expected, actual)
		}
	})
}
