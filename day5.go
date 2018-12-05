package aoc_2018

import (
	"math"
	"strings"
	"unicode"
)

func reactPolymer(input string) int {
	return len(removeReactingUnits(input))
}

func shortestPolimer(input string) int {
	units := units(input)
	min := math.MaxInt32

	for _, v := range units {
		newInput := removeUnitFromInput(input, v)

		reacted := removeReactingUnits(newInput)
		if l := len(reacted); l < min {
			min = l
		}
	}
	return min
}

func removeUnitFromInput(input string, unit rune) string {
	replacer := strings.NewReplacer(
		string(unicode.ToUpper(unit)), "",
		string(unicode.ToLower(unit)), "")

	return replacer.Replace(input)
}

func removeReactingUnits(s string) string {
	var stack []rune

	for _, v := range s {
		if len(stack) == 0 {
			stack = append(stack, v)
		} else {
			topStack := stack[len(stack)-1]
			if cancelEachOther(topStack, v) {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, v)
			}
		}
	}
	return string(stack)
}

func cancelEachOther(b1, b2 rune) bool {
	return unicode.ToUpper(b1) == unicode.ToUpper(b2) && b1 != b2
}

func units(input string) []rune {
	units := make(map[rune]struct{})
	for _, v := range input {
		units[unicode.ToUpper(v)] = struct{}{}
	}

	var runes []rune
	for k, _ := range units {
		runes = append(runes, k)
	}
	return runes
}
