package main

import "testing"

var example = `forward 5
down 5
forward 8
up 3
down 8
forward 2`

func TestPart1(t *testing.T) {
	tests := []struct {
		input  string
		expect int
	}{
		{
			input:  example,
			expect: 150,
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := get_position(tt.input); got != tt.expect {
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
			expect: 900,
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := get_position_with_aim(tt.input); got != tt.expect {
				t.Errorf("Expected %d got %d", tt.expect, got)
			}
		})
	}
}
