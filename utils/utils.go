package utils

import "strconv"

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
