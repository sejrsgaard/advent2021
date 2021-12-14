package main

import "testing"

var example = `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`

func TestPart1(t *testing.T) {
	tests := []struct {
		input  string
		expect int
	}{
		{
			input:  example,
			expect: 198,
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := solve_part1(tt.input); got != tt.expect {
				t.Errorf("Expected %d got %d", tt.expect, got)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		input  string
		expect int
	}{
		{
			input:  example,
			expect: 230,
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := solve_part2(tt.input); got != tt.expect {
				t.Errorf("Expected %d got %d", tt.expect, got)
			}
		})
	}
}
