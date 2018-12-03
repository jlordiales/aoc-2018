package aoc_2018

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFrequency(t *testing.T) {
	assert := assert.New(t)

	result := FrequencyResult(readFile("input-1"))

	assert.Equal(433, result)
}

func TestFirstDuplicate(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int
	}{
		{"Example 1", "+1\n-1", 0},
		{"Example 2", "+3\n+3\n+4\n-2\n-4", 10},
		{"Example 3", readFile("input-1"), 256},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if want, have := tt.expect, FirstDuplicate(tt.input); want != have {
				t.Fatalf("Wanted %d, got %d", want, have)
			}
		})
	}
}

func readFile(name string) string {
	b, _ := ioutil.ReadFile("testdata/" + name)
	return string(b)
}
