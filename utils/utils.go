package utils

import (
	"strconv"
	"strings"
)

func Cast(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func Max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func Sort(s string) string {
	r := Runes(s)
	for x := range r {
		y := x + 1
		for y = range r {
			if r[x] < r[y] {
				r[x], r[y] = r[y], r[x]
			}
		}
	}
	return string(r)
}

func Runes(s string) (runes []rune) {
	for _, r := range s {
		runes = append(runes, r)
	}
	return runes
}

func ContainsRunes(s string, runes []rune) bool {
	for _, r := range runes {
		if !strings.ContainsRune(s, r) {
			return false
		}
	}
	return true
}
