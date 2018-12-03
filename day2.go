package aoc_2018

import (
	"strings"
)

func Checksum(input string) int {
	inputs := strings.Split(input, "\n")
	m1 := 0
	m2 := 0
	for _, v := range inputs {
		freq := groupbyCharOccurrences(strings.TrimSpace(v))
		if _, ok := freq[2]; ok {
			m1++
		}
		if _, ok := freq[3]; ok {
			m2++
		}
	}
	return m1 * m2
}

func CommonLetters(input string) string {
	inputs := strings.Split(input, "\n")
	s1, s2 := pairWithDistance(inputs, 1)
	return commonChars(s1, s2)
}

func pairWithDistance(inputs []string, d int) (string, string) {
	for i := 0; i < len(inputs); i++ {
		for j := i + 1; j < len(inputs); j++ {
			s1 := strings.TrimSpace(inputs[i])
			s2 := strings.TrimSpace(inputs[j])

			if distance(s1, s2) == d {
				return s1, s2
			}
		}
	}
	return "", ""
}

// assume s1 and s2 are the same length
func commonChars(s1, s2 string) string {
	var s strings.Builder
	for k, _ := range s1 {
		if s1[k] == s2[k] {
			s.WriteByte(s1[k])
		}
	}
	return s.String()
}

func distance(s1, s2 string) int {
	return len(s2) - len(commonChars(s1, s2))
}

func groupbyCharOccurrences(s string) map[int]struct{} {
	freq := map[int32]int{}
	for _, v := range s {
		freq[v] = freq[v] + 1
	}

	res := make(map[int]struct{})
	for _, v := range freq {
		res[v] = struct{}{}
	}
	return res
}
