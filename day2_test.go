package aoc_2018

import "testing"

func TestChecksum(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"Example 1",
			`abcdef
			 bababc
			 abbcde
			 abcccd
			 aabcdd
			 abcdee
			 ababab`, 12},
		{"Input", readFile("input-2"), 5750},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Checksum(tt.input); got != tt.want {
				t.Errorf("Checksum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommonLetters(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"Example 1",
			`abcde
			 fghij
			 klmno
			 pqrst
			 fguij
			 axcye
			 wvxyz`, "fgij"},
		{"Input", readFile("input-2"), "tzyvunogzariwkpcbdewmjhxi"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CommonLetters(tt.input); got != tt.want {
				t.Errorf("CommonLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}
