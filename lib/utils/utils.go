package utils

import "strconv"

func IntAbs(l, r int) int {
	v := l - r
	if v < 0 {
		return -v
	}
	return v
}

func IntMod(v, mod int) int {
	if v > 0 && v < mod {
		return v
	}
	return ((v % mod) + mod) % mod
}

func MapToInt(s string) int { return Must(strconv.Atoi(s)) }
