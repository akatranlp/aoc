package aoc

import (
	"aoc/lib/utils"
	"fmt"
	"io"
	"os"
	"time"
)

type Problem interface {
	Part1(r io.Reader) int
	Part2(r io.Reader) int
}

func Run(fileName string, p Problem, problems ...int) {
	file := utils.Must(os.Open(fileName))
	defer file.Close()

	if len(problems) == 0 {
		run1(file, p)
		file.Seek(0, io.SeekStart)
		run2(file, p)
	} else {
		switch problems[0] {
		case 1:
			run1(file, p)
		case 2:
			run2(file, p)
		default:
			panic("false problem number")
		}
	}
}

func run1(r io.Reader, p Problem) {
	start := time.Now()
	answer := p.Part1(r)
	duration := time.Since(start)
	fmt.Printf("Part1: %d - elapsed: %s\n", answer, duration)
}

func run2(r io.Reader, p Problem) {
	start := time.Now()
	answer := p.Part2(r)
	duration := time.Since(start)
	fmt.Printf("Part2: %d - elapsed: %s\n", answer, duration)
}
