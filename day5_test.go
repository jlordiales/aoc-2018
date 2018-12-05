package aoc_2018

import (
	"testing"
)

func TestReactPolymer(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"Example 1", "dabAcCaCBAcCcaDA", 10},
		{"Input", readFile("input-5"), 11668},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reactPolymer(tt.input); got != tt.want {
				t.Errorf("reactPolimer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolve51(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"Example 1", "dabAcCaCBAcCcaDA", 4},
		{"Input", readFile("input-5"), 4652},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestPolimer(tt.input); got != tt.want {
				t.Errorf("shortestPolimer() = %v, want %v", got, tt.want)
			}
		})
	}
}
