package aoc_2018

import (
	"strconv"
	"strings"
)

func FrequencyResult(input string) int {
	result := 0
	ints := toIntSlice(input)
	for _, v := range ints {
		result += v
	}
	return result
}

func FirstDuplicate(input string) int {
	seen := make(map[int]struct{})
	result := 0
	ints := toIntSlice(input)

	i := 0
	seen[0] = struct{}{}
	for {
		result += ints[i]
		if _, ok := seen[result]; ok {
			return result
		}
		seen[result] = struct{}{}
		i = (i + 1) % len(ints)
	}
}

func toIntSlice(input string) []int {
	var res []int
	lines := strings.Split(input, "\n")
	for _, v := range lines {
		k, _ := strconv.Atoi(v)
		res = append(res, k)
	}
	return res
}
