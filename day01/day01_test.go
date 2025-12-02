package main

import (
	"bytes"
	"testing"
)

var part1Test = `
L68
L30
R48
L5
R60
L55
L1
L99
R14
L82
`

func TestDay01(t *testing.T) {
	day01 := Day01{}
	t.Run("part 1", func(t *testing.T) {
		input := bytes.NewBufferString(part1Test)

		expected := 3
		actual := day01.Part1(input)

		if expected != actual {
			t.Fatalf("ERROR: expected %d actual %d\n", expected, actual)
		}
	})

	t.Run("part 2", func(t *testing.T) {
		input := bytes.NewBufferString(part1Test)

		expected := 6
		actual := day01.Part2(input)

		if expected != actual {
			t.Fatalf("ERROR: expected %d actual %d\n", expected, actual)
		}
	})
}
