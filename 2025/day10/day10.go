package main

import (
	"aoc/lib/aoc"
	"aoc/lib/its"
	"aoc/lib/utils"
	"bytes"
	"fmt"
	"io"
	"slices"
	"strings"
)

type Day10 struct{}

var _ aoc.Problem = (*Day10)(nil)

type Machine struct {
	light   uint16
	buttons []uint16
}

func (m *Machine) Print() {
	fmt.Printf("Light:\n- %010b (%d)\n", m.light, m.light)
	fmt.Printf("Buttons\n")
	for _, b := range m.buttons {
		fmt.Printf("- %010b (%d)\n", b, b)
	}
}

type Entry struct {
	alreadyCombined []int
	sum             uint16
}

func (*Day10) Part1(r io.Reader) int {
	return its.Reduce2(its.Enumerate(its.Map(its.Filter(its.ReaderToIter(r), its.FilterEmptyLines), func(row string) Machine {
		parts := strings.Fields(row)
		lightPart := parts[0]
		lightPart = lightPart[1 : len(lightPart)-1]
		lightLen := len(lightPart)
		var light uint16
		for i, c := range lightPart {
			if c == '.' {
				continue
			}
			light |= 1 << (lightLen - 1 - i)
		}
		buttonParts := parts[1 : len(parts)-1]
		var buttons []uint16

		for _, buttonPart := range buttonParts {
			var button uint16
			for numBytes := range bytes.SplitSeq([]byte(buttonPart[1:len(buttonPart)-1]), []byte(",")) {
				idx := utils.MapStrToInt(string(numBytes))
				button |= 1 << (lightLen - 1 - idx)
			}
			buttons = append(buttons, button)
		}

		return Machine{light, buttons}
	})), 0, func(acc int, i int, m Machine) int {
		// m.Print()
		combinations := slices.Collect(its.Map2(slices.All(m.buttons), func(i int, v uint16) Entry {
			return Entry{[]int{i}, v}
		}))

		for i := range len(m.buttons) {
			if checkCombinations(combinations, m.light) {
				return acc + i + 1
			}
			combinations = createCombinations(combinations, m.buttons, m.light)
		}

		return acc
	})
}

func createCombinations(combinations []Entry, buttons []uint16, light uint16) []Entry {
	var newCombinations []Entry
	for _, c := range combinations {
		for i, b := range buttons {
			if slices.Contains(c.alreadyCombined, i) {
				continue
			}
			num := c.sum ^ b
			// fmt.Printf("Checking %v %d\n", c.alreadyCombined, i)
			// fmt.Printf("   %010b\n", c.sum)
			// fmt.Printf("^  %010b\n", b)
			// fmt.Printf("=  %010b\n", num)
			// fmt.Printf("== %010b (%v)\n", light, light == num)
			newCombinations = append(newCombinations, Entry{append(c.alreadyCombined, i), num})
		}
	}
	return newCombinations
}

func checkCombinations(combinations []Entry, light uint16) bool {
	return slices.ContainsFunc(combinations, func(e Entry) bool { return e.sum == light })
}

func (*Day10) Part2(r io.Reader) int {
	fmt.Println("Part2 not implemented")
	return -1
}
