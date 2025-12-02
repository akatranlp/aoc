package its

import (
	"bufio"
	"bytes"
	"io"
	"iter"
)

func ReaderToIter(r io.Reader, splits ...bufio.SplitFunc) iter.Seq[string] {
	scanner := bufio.NewScanner(r)
	split := bufio.ScanLines
	if len(splits) > 0 {
		split = splits[0]
	}
	scanner.Split(split)
	return func(yield func(s string) bool) {
		for scanner.Scan() {
			msg := scanner.Text()
			if !yield(msg) {
				return
			}
		}
	}
}

func SplitByString(s string) bufio.SplitFunc {
	sLen := len(s)
	return func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}
		if i := bytes.Index(data, []byte(s)); i >= 0 {
			return i + sLen, data[:i], nil
		}
		if atEOF {
			return len(data), data, nil
		}
		return 0, nil, nil
	}
}

var SplitByBlocks = SplitByString("\n\n")

func SplitByByte(c byte) bufio.SplitFunc {
	return func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}
		if i := bytes.IndexByte(data, c); i >= 0 {
			return i + 1, data[:i], nil
		}
		if atEOF {
			return len(data), data, nil
		}
		return 0, nil, nil
	}
}
