package main

import "testing"

var example = `3,4,3,1,2`

func Test(t *testing.T) {
	tests := []struct {
		input  string
		days   int
		expect int64
	}{
		{
			input:  example,
			days:   80,
			expect: 5934,
		},
		{
			input:  example,
			days:   256,
			expect: 26984457539,
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := solve(tt.input, tt.days); got != tt.expect {
				t.Errorf("Expected %d got %d", tt.expect, got)
			}
		})
	}
}
