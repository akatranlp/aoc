package utils

import (
	"golang.org/x/exp/constraints"
	"strconv"
)

type Number interface {
	constraints.Float | constraints.Integer
}

func Abs[T Number](l, r T) T {
	v := l - r
	if v < 0 {
		return -v
	}
	return v
}

func Mod[T constraints.Integer](v, mod T) T {
	if v > 0 && v < mod {
		return v
	}
	return ((v % mod) + mod) % mod
}

func MapStrToInt(s string) int { return Must(strconv.Atoi(s)) }
func MapByteToInt(c byte) int  { return int(c - '0') }
