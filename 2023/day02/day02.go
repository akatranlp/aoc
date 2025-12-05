package main

import (
	"aoc/lib/aoc"
	"aoc/lib/its"
	"aoc/lib/utils"
	"io"
	"strings"
)

type Day02 struct{}

var _ aoc.Problem = (*Day02)(nil)

type Round struct {
	r, g, b int
}

func (*Day02) Part1(r io.Reader) int {
	const maxRed int = 12
	const maxGreen int = 13
	const maxBlue int = 14
	var res int
	for i, row := range its.Enumerate(its.Filter(its.ReaderToIter(r), its.FilterEmptyLines)) {
		game := strings.Split(row, ": ")[1]
		var maxRound Round
		for roundStr := range strings.SplitSeq(game, "; ") {
			var round Round
			for color := range strings.SplitSeq(roundStr, ", ") {
				colorParts := strings.Split(color, " ")
				switch colorParts[1] {
				case "blue":
					round.b = utils.MapStrToInt(colorParts[0])
				case "green":
					round.g = utils.MapStrToInt(colorParts[0])
				case "red":
					round.r = utils.MapStrToInt(colorParts[0])
				}
			}
			maxRound.r = max(round.r, maxRound.r)
			maxRound.g = max(round.g, maxRound.g)
			maxRound.b = max(round.b, maxRound.b)
		}
		if maxRound.r <= maxRed && maxRound.g <= maxGreen && maxRound.b <= maxBlue {
			res += i + 1
		}
	}
	return res
}

func (*Day02) Part2(r io.Reader) int {
	var res int
	for row := range its.Filter(its.ReaderToIter(r), its.FilterEmptyLines) {
		game := strings.Split(row, ": ")[1]
		var maxRound Round
		for roundStr := range strings.SplitSeq(game, "; ") {
			var round Round
			for color := range strings.SplitSeq(roundStr, ", ") {
				colorParts := strings.Split(color, " ")
				switch colorParts[1] {
				case "blue":
					round.b = utils.MapStrToInt(colorParts[0])
				case "green":
					round.g = utils.MapStrToInt(colorParts[0])
				case "red":
					round.r = utils.MapStrToInt(colorParts[0])
				}
			}
			maxRound.r = max(round.r, maxRound.r)
			maxRound.g = max(round.g, maxRound.g)
			maxRound.b = max(round.b, maxRound.b)
		}
		res += maxRound.r * maxRound.g * maxRound.b
	}
	return res
}
