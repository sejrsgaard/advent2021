package main

import "testing"

var example = `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`

func Test(t *testing.T) {
	tests := []struct {
		input  string
		expect int
	}{
		{
			input:  example,
			expect: 5,
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
