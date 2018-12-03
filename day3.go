package aoc_2018

import (
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}
type rectangle struct {
	id     string
	start  point
	width  int
	height int
}

func NumberOfOverlappingSquares(input string) int {
	rectangles := parseInput(input)

	return overlappingPoints(rectangles)
}

func FirstNonOverlappingRectangle(input string) string {
	rectangles := parseInput(input)

	visitedPoints := make(map[point]int)
	for _, v := range rectangles {
		addRectangle(v, visitedPoints)
	}

	for _, v := range rectangles {
		if areaOnlyVisitedOnce(v, visitedPoints) {
			return v.id
		}
	}

	return ""
}

func fromInputLine(s string) rectangle {
	parts := strings.Split(s, " ")

	starting := strings.Split(parts[2], ",")
	startingX, _ := strconv.Atoi(starting[0])
	startingY, _ := strconv.Atoi(strings.Replace(starting[1], ":", "", 1))

	dimensions := strings.Split(parts[3], "x")
	width, _ := strconv.Atoi(dimensions[0])
	height, _ := strconv.Atoi(dimensions[1])

	return rectangle{
		id:     parts[0],
		start:  point{x: startingX, y: startingY},
		width:  width,
		height: height,
	}
}

func parseInput(input string) []rectangle {
	inputs := strings.Split(input, "\n")
	var rectangles []rectangle
	for _, v := range inputs {
		rectangles = append(rectangles, fromInputLine(strings.TrimSpace(v)))
	}
	return rectangles
}

func areaOnlyVisitedOnce(r rectangle, visited map[point]int) bool {
	for x := r.start.x; x < r.start.x+r.width; x++ {
		for y := r.start.y; y < r.start.y+r.height; y++ {
			p := point{x, y}
			if visited[p] > 1 {
				return false
			}
		}
	}
	return true
}

func overlappingPoints(r []rectangle) int {
	visitedPoints := make(map[point]int)
	for _, v := range r {
		addRectangle(v, visitedPoints)
	}

	visitedMoreThanOnce := 0
	for _, v := range visitedPoints {
		if v > 1 {
			visitedMoreThanOnce++
		}
	}

	return visitedMoreThanOnce
}

func addRectangle(r rectangle, points map[point]int) {
	for x := r.start.x; x < r.start.x+r.width; x++ {
		for y := r.start.y; y < r.start.y+r.height; y++ {
			p := point{x, y}
			points[p] = points[p] + 1
		}
	}
}
