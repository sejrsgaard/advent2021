package main

import "testing"

var example = `be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe
edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc
fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg
fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb
aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea
fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb
dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe
bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef
egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb
gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce`

func TestPart1(t *testing.T) {
	tests := []struct {
		input  string
		expect int
	}{
		{
			input:  example,
			expect: 26,
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
			expect: 61229,
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
