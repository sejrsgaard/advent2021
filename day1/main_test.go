package main

import "testing"

var example = `199
200
208
210
200
207
240
269
260
263`

func Test(t *testing.T) {
	tests := []struct {
		input  string
		sw     int
		expect int
	}{
		{
			input:  example,
			sw:     1,
			expect: 7,
		},
		{
			input:  example,
			sw:     3,
			expect: 5,
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := rate(tt.input, tt.sw); got != tt.expect {
				t.Errorf("Expected %d for sliding window %d got %d", tt.expect, tt.sw, got)
			}
		})
	}
}
