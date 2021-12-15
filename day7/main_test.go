package main

import "testing"

var example = `16,1,2,0,4,2,7,1,2,14`

func TestPart1(t *testing.T) {
	tests := []struct {
		input  string
		expect int
	}{
		{
			input:  example,
			expect: 37,
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := solve(tt.input); got != tt.expect {
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
			expect: 168,
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := solve2(tt.input); got != tt.expect {
				t.Errorf("Expected %d got %d", tt.expect, got)
			}
		})
	}
}
