package aoc_2018

import "testing"

const (
	exampleInput = `[1518-11-01 00:00] Guard #10 begins shift
[1518-11-01 00:05] falls asleep
[1518-11-01 00:25] wakes up
[1518-11-01 00:30] falls asleep
[1518-11-01 00:55] wakes up
[1518-11-01 23:58] Guard #99 begins shift
[1518-11-02 00:40] falls asleep
[1518-11-02 00:50] wakes up
[1518-11-03 00:05] Guard #10 begins shift
[1518-11-03 00:24] falls asleep
[1518-11-03 00:29] wakes up
[1518-11-04 00:02] Guard #99 begins shift
[1518-11-04 00:36] falls asleep
[1518-11-04 00:46] wakes up
[1518-11-05 00:03] Guard #99 begins shift
[1518-11-05 00:45] falls asleep
[1518-11-05 00:55] wakes up`
)

func TestSolve1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"Example 1", exampleInput, 240},
		{"Input", readFile("input-4"), 104764},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Strategy1(tt.input); got != tt.want {
				t.Errorf("Strategy1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolve42(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"Example 1", exampleInput, 4455},
		{"Input", readFile("input-4"), 128617},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Strategy2(tt.input); got != tt.want {
				t.Errorf("Strategy2() = %v, want %v", got, tt.want)
			}
		})
	}
}
