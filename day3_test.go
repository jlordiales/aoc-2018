package aoc_2018

import (
	"testing"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"Example 1", `#1 @ 1,3: 4x4
					   #2 @ 3,1: 4x4
					   #3 @ 5,5: 2x2`,
			4},
		{"Input", readFile("input-3"), 103482},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumberOfOverlappingSquares(tt.input); got != tt.want {
				t.Errorf("Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolve2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"Example 1", `#1 @ 1,3: 4x4
					   #2 @ 3,1: 4x4
					   #3 @ 5,5: 2x2`,
			"#3"},
		{"Input", readFile("input-3"), "#686"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FirstNonOverlappingRectangle(tt.input); got != tt.want {
				t.Errorf("Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
